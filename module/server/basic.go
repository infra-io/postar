// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/12 23:27:54

package server

import (
	"github.com/avino-plan/postar/module"
	"github.com/avino-plan/postar/module/sender"
)

var (
	servers = map[string]func() Server{
		"http": newHttpServer,
	}
)

type Server interface {
	module.Configurer
	SetSender(sender sender.Sender)
	Serve() error
	Close() error
}

func Initialize(config *module.Config) (Server, error) {
	server := servers[config.Global.ServerType]()
	return server, server.Configure(config)
}
