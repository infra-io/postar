// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/12 23:27:54

package server

import (
	"fmt"

	"github.com/avino-plan/postar/module"
	"github.com/avino-plan/postar/module/sender"
)

var (
	servers = map[string]func() Server{
		"http": newHttpServer,
	}
)

type SendRequest struct {
	Email   *sender.Email       `json:"email"`
	Options *sender.SendOptions `json:"options"`
}

func newSendRequest() *SendRequest {
	sendOptions := sender.DefaultSendOptions()
	return &SendRequest{
		Email:   nil,
		Options: &sendOptions,
	}
}

type Server interface {
	module.Configurer
	ConfigureSender(sender sender.Sender)
	Serve() error
	Close() error
}

func Initialize(config *module.Config) (Server, error) {

	newServer, ok := servers[config.Global.ServerType]
	if !ok {
		return nil, fmt.Errorf("server type %s not found", config.Global.ServerType)
	}

	server := newServer()
	return server, server.Configure(config)
}
