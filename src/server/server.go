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
)

var (
	// initServers stores all servers that can be initialized.
	initServers = map[string]func(port string, closedPort string) *sync.WaitGroup{
		"http": http.InitServer,
	}

	// stopServers stores all servers that can be shutdown.
	stopServers = map[string]func() error{
		"http": http.StopServer,
	}
)

// RunServer runs a server for service and shutdown.
// Notice that the returning value is *sync.WaitGroup, so you can use it to
// block your main goroutine before closing the server.
func RunServer() *sync.WaitGroup {
	initServer, ok := initServers[core.ServerType()]
	if !ok {
		core.Logger().Errorf("The initializing server type %s doesn't exist! Try 'http'?", core.ServerType())
		return nil
	}
	return initServer(core.ServerPort(), core.ServerClosedPort())
}

// ShutdownServer shutdowns the running server.
func ShutdownServer() error {
	stopServer, ok := stopServers[core.ServerType()]
	if !ok {
		msg := fmt.Sprintf("The stopping server type %s doesn't exist! Try 'http'?", core.ServerType())
		core.Logger().Errorf(msg)
		return errors.New(msg)
	}
	return stopServer()
}
