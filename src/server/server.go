// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/15 23:44:43

package server

import (
	"errors"
	"fmt"
	"sync"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/server/http"
	"github.com/avino-plan/postar/src/server/jsonrpc"
)

var (
	// servers stores all servers that can be used.
	servers = map[string]Server{
		"http":    http.NewServer(),
		"jsonrpc": jsonrpc.NewServer(),
		"grpc":    nil,
	}
)

// Server is the interface of servers.
type Server interface {

	// Init should init this server with port and closedPort,
	// and return a pointer of sync.WaitGroup.
	Init(port string, closedPort string) *sync.WaitGroup

	// Stop should stop this server and return an error if failed.
	Stop() error
}

// InitServer initializes a server for use.
// Notice that the returning value is *sync.WaitGroup, so you can use it to
// block your main goroutine before closing the server.
func InitServer() *sync.WaitGroup {
	server, ok := servers[core.ServerType()]
	if !ok {
		core.Logger().Errorf("The initializing server type %s doesn't exist! Try these ['http', 'jsonrpc', 'grpc']?", core.ServerType())
		return &sync.WaitGroup{}
	}
	return server.Init(core.ServerPort(), core.ServerClosedPort())
}

// StopServer stops the running server.
func StopServer() error {
	server, ok := servers[core.ServerType()]
	if !ok {
		msg := fmt.Sprintf("The stopping server type %s doesn't exist! Try these ['http', 'jsonrpc', 'grpc']?", core.ServerType())
		core.Logger().Errorf(msg)
		return errors.New(msg)
	}
	return server.Stop()
}
