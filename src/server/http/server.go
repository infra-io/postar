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
	"github.com/kataras/iris/v12"
)

// ServerImpl is an implement of Server interface which provides http functions.
type serverImpl struct {

	// serverForService is for main service.
	serverForService *iris.Application

	// serverForShutdown is for closed service.
	serverForShutdown *iris.Application

	// wg is for waiting these servers.
	wg *sync.WaitGroup
}

// NewServer returns an empty server implement.
func NewServer() *serverImpl {
	return &serverImpl{}
}

// initServerForService initializes the server for service.
func (si *serverImpl) initServerForService(port string, afterListening func()) {
	si.serverForService = iris.New()
	si.serverForService.Logger().SetLevel("disable")
	si.serverForService.Get("/", newPingHandler())
	si.serverForService.Post("/send", newSendHandler())
	err := si.serverForService.Listen(":"+port, iris.WithoutStartupLog, iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		core.Logger().Errorf("Failed to listen to the port %s! Please try another one. Error: %s.", port, err.Error())
	}
	afterListening()
}

// initServerForShutdown initializes the server for shutdown.
func (si *serverImpl) initServerForShutdown(port string, afterListening func()) {
	si.serverForShutdown = iris.New()
	si.serverForShutdown.Logger().SetLevel("disable")
	si.serverForShutdown.Post("/close", newCloseHandler(si))
	err := si.serverForShutdown.Listen(":"+port, iris.WithoutStartupLog, iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		core.Logger().Errorf("The port %s maybe used! Try to change another one! [%s]", port, err.Error())
	}
	afterListening()
}

// InitServer initializes servers with given two ports.
func (si *serverImpl) Init(port string, closedPort string) *sync.WaitGroup {

	// Create a wait group to wait these servers.
	si.wg = &sync.WaitGroup{}

	// Notice that wg.Add must be executed before wg.Done, so they can't code in go func.
	si.wg.Add(1)
	go si.initServerForService(port, func() {
		core.Logger().Debug("Add 1 to wg in initServerForService...")
		si.wg.Done()
	})

	si.wg.Add(1)
	go si.initServerForShutdown(closedPort, func() {
		core.Logger().Debug("Add 1 to wg in initServerForShutdown...")
		si.wg.Done()
	})

	core.Logger().Infof("The main service is using port %s.", port)
	core.Logger().Infof("The closed service is using port %s.", closedPort)
	return si.wg
}

// StopServer stops running servers.
func (si *serverImpl) Stop() error {

	// Send a request to server.
	response, err := http.Post("http://localhost:"+core.ServerClosedPort()+"/close", "application/json; charset=utf-8", nil)
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
