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
	"strings"
	"sync"

	"github.com/avino-plan/postar/src/core"
)

// ServerImpl is an implement of Server interface which provides jsonrpc functions.
type serverImpl struct {

	// listenerForService listens the main service.
	listenerForService net.Listener

	// wg is for waiting these servers.
	wg *sync.WaitGroup
}

// NewServer returns an empty server implement.
func NewServer() *serverImpl {
	return &serverImpl{}
}

// initServerForService initializes the server for service.
func (si *serverImpl) initServerForService(port string, beforeServing func(), cleanUp func()) {

	// Create a new rpc server, and register some services.
	server := rpc.NewServer()
	server.Register(&PostarService{})

	// Create a listener, and bind it to port.
	var err error
	si.listenerForService, err = net.Listen("tcp", ":"+port)
	if err != nil {
		core.Logger().Errorf("Failed to listen to the port %s! Please try another one. Error: %s.", port, err.Error())
		cleanUp()
		return
	}
	beforeServing()

	// Start serving.
	go func() {
		connWg := &sync.WaitGroup{}
		for {
			conn, err := si.listenerForService.Accept()
			if err != nil {
				// The err will be "use of closed network connection" if listener has been closed.
				// Actually, this is a stupid way but em...
				if strings.Contains(err.Error(), "use of closed network connection") {
					break
				}

				core.Logger().Errorf("Failed to accept the connection! Error: %s.", err.Error())
				continue
			}

			// Record every connection.
			connWg.Add(1)
			go func(conn net.Conn) {
				server.ServeCodec(stdJsonRPC.NewServerCodec(conn))
				connWg.Done()
			}(conn)
		}

		// Wait for all connections have done.
		connWg.Wait()
		cleanUp()
	}()
}

// initServerForShutdown initializes the server for shutdown.
func (si *serverImpl) initServerForShutdown(port string, cleanUp func()) {

	// Create a new rpc server, and register some services.
	server := rpc.NewServer()
	server.Register(NewCloseService(si.listenerForService))

	// Create a listener, and bind it to port.
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		core.Logger().Errorf("Failed to listen to the port %s! Please try another one. Error: %s.", port, err.Error())
		cleanUp()
		return
	}

	// Start serving.
	go func() {
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				core.Logger().Errorf("Failed to accept the connection! Error: %s.", err.Error())
				continue
			}
			server.ServeCodec(stdJsonRPC.NewServerCodec(conn))
			break
		}
		cleanUp()
	}()
}

// InitServer initializes servers with given two ports.
func (si *serverImpl) Init(port string, closedPort string) *sync.WaitGroup {

	// Create a wait group to wait these servers.
	si.wg = &sync.WaitGroup{}

	// Notice that wg.Add must be executed before wg.Done, so they can't code in go func.
	si.wg.Add(1)
	si.initServerForService(
		port,
		func() {
			si.wg.Add(1)
			si.initServerForShutdown(closedPort, func() {
				core.Logger().Debug("Add 1 to wg in initServerForShutdown...")
				si.wg.Done()
			})
		},
		func() {
			core.Logger().Debug("Add 1 to wg in initServerForService...")
			si.wg.Done()
		},
	)

	core.Logger().Infof("The main service is using port %s.", port)
	core.Logger().Infof("The closed service is using port %s.", closedPort)
	return si.wg
}

// StopServer stops running servers.
func (si *serverImpl) Stop(closedPort string) error {

	// Connect to the remote server.
	conn, err := stdJsonRPC.Dial("tcp", "127.0.0.1:"+closedPort)
	if err != nil {
		core.Logger().Errorf("Failed to request server. Error: %s.", err.Error())
		return err
	}
	defer conn.Close()

	// Send a request to server.
	req := &CloseRequest{}
	resp := &Result{}
	err = conn.Call("CloseService.Close", req, resp)
	if err != nil {
		core.Logger().Errorf("Failed to call the remote service. Error: %s.", err.Error())
		return err
	}

	core.Logger().Info("Successfully closed the running servers.")
	return nil
}
