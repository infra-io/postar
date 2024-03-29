// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"net"
	"net/http"

	"github.com/FishGoddess/logit"
	postaradminv1 "github.com/infra-io/postar/api/genproto/postaradmin/v1"
	"github.com/infra-io/postar/configs"
	"github.com/infra-io/postar/internal/postar-admin/service"
	grpcx "github.com/infra-io/postar/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayServer struct {
	conf       *configs.PostarAdminConfig
	grpcServer Server
	httpServer *http.Server
}

func newGrpcServer(conf *configs.PostarAdminConfig, spaceService service.SpaceService, accountService service.AccountService, templateService service.TemplateService) (Server, http.Handler, error) {
	ctx := context.Background()
	mux := grpcx.NewGatewayMux()
	endpoint := conf.Server.GrpcEndpoint

	var opts []grpc.DialOption
	if conf.Server.TLS() {
		creds, err := credentials.NewClientTLSFromFile(conf.Server.CertFile, "")
		if err != nil {
			return nil, nil, err
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	err := postaradminv1.RegisterSpaceServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	err = postaradminv1.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	err = postaradminv1.RegisterTemplateServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	grpcServer, err := NewGrpcServer(conf, spaceService, accountService, templateService)
	return grpcServer, mux, err
}

func NewGatewayServer(conf *configs.PostarAdminConfig, spaceService service.SpaceService, accountService service.AccountService, templateService service.TemplateService) (Server, error) {
	grpcServer, handler, err := newGrpcServer(conf, spaceService, accountService, templateService)
	if err != nil {
		return nil, err
	}

	httpServer := &http.Server{
		Addr:    conf.Server.HttpEndpoint,
		Handler: handler,
	}

	gs := &GatewayServer{
		conf:       conf,
		grpcServer: grpcServer,
		httpServer: httpServer,
	}

	return gs, nil
}

func (gs *GatewayServer) Serve() error {
	httpEndpoint := gs.conf.Server.HttpEndpoint

	listener, err := net.Listen("tcp", httpEndpoint)
	if err != nil {
		logit.Error("listen tcp failed", "err", err, "endpoint", httpEndpoint)
		return err
	}

	defer listener.Close()

	go monitorCloseSignals(gs)
	go func() {
		if err := gs.grpcServer.Serve(); err != nil {
			grpcEndpoint := gs.conf.Server.GrpcEndpoint
			logit.Error("serve grpc failed", "err", err, "endpoint", grpcEndpoint)
		}
	}()

	err = gs.httpServer.Serve(listener)
	if err != nil && err != http.ErrServerClosed {
		logit.Error("serve http failed", "err", err, "endpoint", httpEndpoint)
		return err
	}

	return nil
}

func (gs *GatewayServer) Close() error {
	timeout := gs.conf.Server.MaxCloseWaitTime.Standard()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return gs.httpServer.Shutdown(ctx)
}
