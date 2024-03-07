// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"net"

	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/logit"
	postaradminv1 "github.com/infra-io/postar/api/genproto/postaradmin/v1"
	"github.com/infra-io/postar/configs"
	"github.com/infra-io/postar/internal/postar-admin/service"
	"github.com/infra-io/postar/pkg/grpc/contextutil"
	"github.com/infra-io/postar/pkg/grpc/logging"
	grpcx "github.com/infra-io/servicex/net/grpc"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	postaradminv1.UnimplementedSpaceServiceServer
	postaradminv1.UnimplementedAccountServiceServer
	postaradminv1.UnimplementedTemplateServiceServer

	conf   *configs.PostarAdminConfig
	server *grpc.Server

	spaceService    service.SpaceService
	accountService  service.AccountService
	templateService service.TemplateService
}

func NewGrpcServer(conf *configs.PostarAdminConfig, spaceService service.SpaceService, accountService service.AccountService, templateService service.TemplateService) (Server, error) {
	timeout := conf.Server.RequestTimeout.Standard()
	interceptor := grpcx.Interceptor(timeout, logging.ResolveRequest)
	checkSpace := checkSpaceInterceptor(spaceService)
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor, checkSpace))

	gs := &GrpcServer{
		conf:            conf,
		server:          grpcServer,
		spaceService:    spaceService,
		accountService:  accountService,
		templateService: templateService,
	}

	postaradminv1.RegisterSpaceServiceServer(gs.server, gs)
	postaradminv1.RegisterAccountServiceServer(gs.server, gs)
	postaradminv1.RegisterTemplateServiceServer(gs.server, gs)

	return gs, nil
}

func (gs *GrpcServer) Serve() error {
	endpoint := gs.conf.Server.GrpcEndpoint

	listener, err := net.Listen("tcp", endpoint)
	if err != nil {
		logit.Error("listen tcp failed", "err", err, "endpoint", endpoint)
		return err
	}

	defer listener.Close()
	go monitorCloseSignals(gs)

	return gs.server.Serve(listener)
}

func (gs *GrpcServer) Close() error {
	stopCh := make(chan struct{}, 1)

	go func() {
		gs.server.GracefulStop()
		stopCh <- struct{}{}
	}()

	timeout := gs.conf.Server.MaxCloseWaitTime.Standard()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-stopCh:
		return nil
	}
}

func checkSpaceInterceptor(spaceService service.SpaceService) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		spaceID := contextutil.GetSpaceID(ctx)
		spaceToken := contextutil.GetSpaceToken(ctx)

		if spaceID <= 0 {
			return handler(ctx, req)
		}

		space, err := spaceService.GetSpace(ctx, spaceID, true)
		if err != nil {
			return nil, err
		}

		if space.Token != spaceToken {
			err = errors.New("wrong token")
			return nil, errors.Forbidden(err, errors.WithMsg("业务空间的令牌错误"))
		}

		return handler(ctx, req)
	}
}
