// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/config"
	"github.com/infra-io/postar/internal/postar/service"
)

var (
	servers = map[string]newServerFunc{
		"grpc":    NewGrpcServer,
		"gateway": NewGatewayServer,
	}
)

type newServerFunc func(conf *config.PostarConfig, emailBiz service.EmailService) (Server, error)

type Server interface {
	Serve() error
	Close() error
}

func New(conf *config.PostarConfig, emailBiz service.EmailService) (Server, error) {
	newServer, ok := servers[conf.Server.Type]
	if !ok {
		return nil, fmt.Errorf("server: type %s not found", conf.Server.Type)
	}

	return newServer(conf, emailBiz)
}

func monitorCloseSignals(svr Server) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	sign := <-signalCh

	if err := svr.Close(); err != nil {
		logit.Error("close server failed", "err", err, "signal", sign)
	}
}
