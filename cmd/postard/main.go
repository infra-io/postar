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
	"github.com/avinoplan/postar/internal/biz"
	"github.com/avinoplan/postar/internal/server"
	"github.com/panjf2000/ants/v2"
)

func main() {
	logger := logit.NewLogger()
	pool, err := ants.NewPool(64)
	if err != nil {
		panic(err)
	}
	defer pool.Release()

	smtpBiz := biz.NewSmtpBiz(pool, "", 0, "", "")
	svr := server.NewGrpcServer(logger, smtpBiz)

	listener, err := net.Listen("tcp", ":5897")
	if err != nil {
		panic(err)
	}

	err = svr.Run(listener)
	if err != nil {
		panic(err)
	}
}
