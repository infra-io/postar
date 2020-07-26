// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/26 20:53:56

package grpc

import (
	"context"
	"math"
	"net"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/server/based"
	"github.com/avino-plan/postar/src/server/grpc/services"
	"google.golang.org/grpc"
)

// ServerImpl is an implement of Server interface which provides grpc functions.
type serverImpl struct {

	// Based on this server.
	*based.BasedServer

	// server is the main service.
	server *grpc.Server
}

// NewServer returns an empty server implement.
func NewServer() *serverImpl {
	result := &serverImpl{}
	result.BasedServer = &based.BasedServer{
		InitServerForService:  result.initServerForService,
		InitServerForShutdown: result.initServerForShutdown,
	}
	return result
}

// initServerForService initializes the server for service.
func (si *serverImpl) initServerForService(port string, beforeServing func(), cleanUp func()) {

	// Create a gRPC server, and set MsgSize to unlimited.
	si.server = grpc.NewServer(grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32))

	// Register services.
	services.RegisterPostarServiceServer(si.server, &services.PostarServiceImpl{})

	// Start listening.
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		core.Logger().Errorf("Failed to listen to the port %s! Please try another one. Error: %s.", port, err.Error())
		cleanUp()
		return
	}
	beforeServing()

	// Start serving.
	go func() {
		core.Logger().Debug("Before main serving...")
		err := si.server.Serve(listener)
		if err != nil {
			core.Logger().Errorf("Failed to serve! Error: %s.", err.Error())
		}
		cleanUp()
	}()
}

// initServerForShutdown initializes the server for shutdown.
func (si *serverImpl) initServerForShutdown(port string, cleanUp func()) {

	// Create a gRPC server, and set MsgSize to unlimited.
	server := grpc.NewServer(grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32))

	// Register services.
	services.RegisterCloseServiceServer(server, services.NewCloseService(si.server, server))

	// Start listening.
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		core.Logger().Errorf("Failed to listen to the port %s! Please try another one. Error: %s.", port, err.Error())
		cleanUp()
		return
	}

	// Start serving.
	go func() {
		core.Logger().Debug("Before shutdown serving...")
		err := server.Serve(listener)
		if err != nil {
			core.Logger().Errorf("Failed to serve! Error: %s.", err.Error())
		}
		cleanUp()
	}()
}

// StopServer stops the running servers.
func (si *serverImpl) Stop(closedPort string) error {

	// Connect to the remote server.
	conn, err := grpc.Dial("127.0.0.1:"+closedPort, grpc.WithInsecure())
	if err != nil {
		core.Logger().Errorf("Failed to connect to server. Error: %s.", err.Error())
		return err
	}
	defer conn.Close()

	// Create a new client.
	client := services.NewCloseServiceClient(conn)

	// Send a request to server.
	resp, err := client.Close(context.Background(), &services.EmptyRequest{})
	if err != nil {
		core.Logger().Errorf("Failed to call the remote service. Error: %s.", err.Error())
		return err
	}

	core.Logger().Debug(string(resp.Data))
	core.Logger().Info("Successfully closed the running servers.")
	return nil
}
