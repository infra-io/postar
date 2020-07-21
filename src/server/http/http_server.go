// Copyright 2020 Ye Zi Jie. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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

var (
	// serverForService is for main service and serverForShutdown is for closed service.
	serverForService  *iris.Application
	serverForShutdown *iris.Application
)

// InitServer initializes servers with given two ports.
func InitServer(port string, closedPort string) *sync.WaitGroup {

	// Create a wait group to wait these servers.
	wg := &sync.WaitGroup{}
	go func() {
		wg.Add(1)
		initServerForService(port, func() {
			wg.Done()
		})
	}()

	go func() {
		wg.Add(1)
		initServerForShutdown(closedPort, func() {
			wg.Done()
		})
	}()

	core.Logger().Infof("The main service is using port %s.", port)
	core.Logger().Infof("The closed service is using port %s.", closedPort)
	return wg
}

// initServerForService initializes the server for service.
func initServerForService(port string, afterListening func()) {
	serverForService = iris.New()
	serverForService.Logger().SetLevel("disable")
	serverForService.Get("/", pingHandler)
	serverForService.Get("/ping", pingHandler)
	serverForService.Post("/send", sendHandler)
	err := serverForService.Listen(":"+port, iris.WithoutStartupLog, iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		core.Logger().Errorf("The port %s maybe used! Try to change another one! [%s]", port, err.Error())
	}
	afterListening()
}

// initServerForShutdown initializes the server for shutdown.
func initServerForShutdown(port string, afterListening func()) {
	serverForShutdown = iris.New()
	serverForShutdown.Logger().SetLevel("disable")
	serverForShutdown.Post("/close", closeHandler)
	err := serverForShutdown.Listen(":"+port, iris.WithoutStartupLog, iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		core.Logger().Errorf("The port %s maybe used! Try to change another one! [%s]", port, err.Error())
	}
	afterListening()
}

// StopServer stops running servers.
func StopServer() error {

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

	return nil
}
