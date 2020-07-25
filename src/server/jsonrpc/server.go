// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/25 21:41:13

package jsonrpc

import (
	"net"
	"net/rpc"
	stdJsonRPC "net/rpc/jsonrpc"
	"sync"

	"github.com/avino-plan/postar/src/core"
)

// ServerImpl is an implement of Server interface which provides jsonrpc functions.
type serverImpl struct {
}

// NewServer returns an empty server implement.
func NewServer() *serverImpl {
	return &serverImpl{}
}

// initServerForService initializes the server for service.
func (si *serverImpl) initServerForService(port string, afterListening func()) {

}

// initServerForShutdown initializes the server for shutdown.
func (si *serverImpl) initServerForShutdown(port string, afterListening func()) {}

// InitServer initializes servers with given two ports.
func (si *serverImpl) Init(port string, closedPort string) *sync.WaitGroup {

	// 注册服务，并使用 TCP 协议作为传输载体
	serverForService := rpc.NewServer()
	serverForService.Register(&PostarService{})

	// 创建监听器
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		core.Logger().Errorf("Failed to listen to the port %s! Please try another one. Error: %s.", port, err.Error())
		return &sync.WaitGroup{}
	}
	defer listener.Close()

	// 使用 TCP 服务器进行服务
	for {
		conn, err := listener.Accept()
		if err != nil {
			core.Logger().Errorf("Failed to accept the connection! Error: %s.", err.Error())
			continue
		}

		go func(conn net.Conn) {
			stdJsonRPC.ServeConn(conn)
		}(conn)
	}
}

// StopServer stops running servers.
func (si *serverImpl) Stop() error {
	return nil
}
