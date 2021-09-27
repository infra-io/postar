// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 01:33:43

package main

import (
	"net"

	"github.com/FishGoddess/logit"
	"github.com/avino-plan/postar/internal/pkg/concurrency"
	"github.com/avino-plan/postar/internal/postard/server"
	"github.com/avino-plan/postar/internal/postard/service"
)

func main() {
	logger := logit.NewLogger()
	contextService := service.NewContextService(logger)

	pool := concurrency.NewPool()
	smtpService := service.NewSmtpService(pool, "", 0, "", "")

	svr := server.NewPostardGrpcServer(contextService, smtpService)

	listener, err := net.Listen("tcp", ":5897")
	if err != nil {
		panic(err)
	}

	err = svr.Run(listener)
	if err != nil {
		panic(err)
	}
}
