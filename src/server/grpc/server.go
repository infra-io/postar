// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/26 20:53:56

package grpc

import (
	"sync"

	"github.com/avino-plan/postar/src/server/based"
)

// ServerImpl is an implement of Server interface which provides grpc functions.
type serverImpl struct {

	// Based on this server.
	*based.BasedServer
}

// NewServer returns an empty server implement.
func NewServer() *serverImpl {
	return &serverImpl{
		BasedServer: &based.BasedServer{},
	}
}

// initServerForService initializes the server for service.
func (si *serverImpl) initServerForService(port string, beforeServing func(), cleanUp func()) {

}

// initServerForShutdown initializes the server for shutdown.
func (si *serverImpl) initServerForShutdown(port string, cleanUp func()) {

}

// InitServer initializes servers with given two ports.
func (si *serverImpl) Init(port string, closedPort string) *sync.WaitGroup {
	return si.BasedServer.Init(si.initServerForService, si.initServerForShutdown, port, closedPort)
}

// StopServer stops the running servers.
func (si *serverImpl) Stop(closedPort string) error {
	// TODO
	return nil
}
