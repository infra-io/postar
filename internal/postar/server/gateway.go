// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"net"
	"net/http"

	"github.com/FishGoddess/logit"
	postarv1 "github.com/infra-io/postar/api/genproto/postar/v1"
	"github.com/infra-io/postar/configs"
	"github.com/infra-io/postar/internal/postar/service"
	grpcx "github.com/infra-io/postar/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayServer struct {
	conf       *configs.PostarConfig
	grpcServer Server
	httpServer *http.Server
}

func newGrpcServer(conf *configs.PostarConfig, emailService service.EmailService) (Server, http.Handler, error) {
	ctx := context.Background()
	mux := grpcx.NewGatewayMux()
	endpoint := conf.Server.GrpcEndpoint
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := postarv1.RegisterEmailServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	grpcServer, err := NewGrpcServer(conf, emailService)
	return grpcServer, mux, err
}

func NewGatewayServer(conf *configs.PostarConfig, emailService service.EmailService) (Server, error) {
	grpcServer, handler, err := newGrpcServer(conf, emailService)
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
