// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"net"

	"github.com/FishGoddess/logit"
	postarv1 "github.com/infra-io/postar/api/genproto/postar/v1"
	"github.com/infra-io/postar/configs"
	"github.com/infra-io/postar/internal/postar/service"
	"github.com/infra-io/postar/pkg/grpc/logging"
	grpcx "github.com/infra-io/servicex/net/grpc"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	postarv1.UnimplementedEmailServiceServer

	conf   *configs.PostarConfig
	server *grpc.Server

	emailService service.EmailService
}

func NewGrpcServer(conf *configs.PostarConfig, emailService service.EmailService) (Server, error) {
	timeout := conf.Server.RequestTimeout.Standard()
	interceptor := grpcx.Interceptor(timeout, logging.ResolveRequest)
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor))

	gs := &GrpcServer{
		conf:         conf,
		server:       grpcServer,
		emailService: emailService,
	}

	postarv1.RegisterEmailServiceServer(gs.server, gs)

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
