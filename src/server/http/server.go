// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/16 00:07:38

package http

import (
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/server/based"
	"github.com/kataras/iris/v12"
)

// ServerImpl is an implement of Server interface which provides http functions.
type serverImpl struct {

	// Based on this server.
	*based.BasedServer

	// serverForService is for main service.
	serverForService *iris.Application

	// serverForShutdown is for closed service.
	serverForShutdown *iris.Application
}

// NewServer returns an empty server implement.
func NewServer() *serverImpl {
	return &serverImpl{
		BasedServer: &based.BasedServer{},
	}
}

// initServerForService initializes the server for service.
func (si *serverImpl) initServerForService(port string, beforeServing func(), cleanUp func()) {

	// Create a new server, and register some services.
	si.serverForService = iris.New()
	si.serverForService.Logger().SetLevel("disable")
	si.serverForService.Get("/ping", newPingHandler())
	si.serverForService.Post("/send", newSendHandler())

	// Start serving.
	beforeServing()
	go func() {
		core.Logger().Debug("Before listening...")
		err := si.serverForService.Listen(":"+port, iris.WithoutStartupLog, iris.WithoutServerError(iris.ErrServerClosed))
		if err != nil {
			core.Logger().Errorf("Failed to listen to the port %s! Please try another one. Error: %s.", port, err.Error())
		}
		cleanUp()
	}()
}

// initServerForShutdown initializes the server for shutdown.
func (si *serverImpl) initServerForShutdown(port string, cleanUp func()) {

	// Create a new server, and register some services.
	si.serverForShutdown = iris.New()
	si.serverForShutdown.Logger().SetLevel("disable")
	si.serverForShutdown.Post("/close", newCloseHandler(si))

	// Start serving.
	go func() {
		core.Logger().Debug("Before listening...")
		err := si.serverForShutdown.Listen(":"+port, iris.WithoutStartupLog, iris.WithoutServerError(iris.ErrServerClosed))
		if err != nil {
			core.Logger().Errorf("The port %s maybe used! Try to change another one! [%s]", port, err.Error())
		}
		cleanUp()
	}()
}

// InitServer initializes servers with given two ports.
func (si *serverImpl) Init(port string, closedPort string) *sync.WaitGroup {
	return si.BasedServer.Init(si.initServerForService, si.initServerForShutdown, port, closedPort)
}

// StopServer stops the running servers.
func (si *serverImpl) Stop(closedPort string) error {

	// Send a request to server.
	response, err := http.Post("http://127.0.0.1:"+closedPort+"/close", "application/json; charset=utf-8", nil)
	if err != nil {
		core.Logger().Errorf("Failed to request server. Error: %s.", err.Error())
		return err
	}

	// If StatusCode != OK, then something wrong happened.
	if response.StatusCode != iris.StatusOK {
		core.Logger().Errorf("Server doesn't response ok. StatusCode is [%d].", response.StatusCode)

		// Try to read body of response.
		defer response.Body.Close()
		buffer := &strings.Builder{}
		written, err := io.Copy(buffer, response.Body)
		if written == 0 || (err != nil && err != io.EOF) {
			core.Logger().Errorf("Failed to read response's body. Written byte is [%d]. Error: %s.", written, err.Error())
			return err
		}
		core.Logger().Error(buffer.String())
	}

	core.Logger().Info("Successfully closed the running servers.")
	return nil
}
