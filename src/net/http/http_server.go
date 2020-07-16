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
	"context"
	"os"
	"sync"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/models"
	"github.com/kataras/iris/v12"
)

var (
	// serverForService is for main service and serverForShutdown is for closed service.
	serverForService  *iris.Application
	serverForShutdown *iris.Application
)

// InitServer initializes servers with given two ports.
func InitServer(port string, closedPort string) *sync.WaitGroup {
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

	core.Logger().Infof("The main service is using port %s. The closed service is using port %s.", port, closedPort)
	return wg
}

// initServerForService initializes the server for service.
func initServerForService(port string, afterListening func()) {
	serverForService = iris.New()
	serverForService.Logger().SetLevel("disable")
	serverForService.Get("/", pingHandler)
	serverForService.Post("/send", sendHandler)
	err := serverForService.Listen(":"+port, iris.WithoutStartupLog)
	if err != nil {
		core.Logger().Errorf("The port %s maybe used! Try to change another one!", port)
	}
	afterListening()
}

// initServerForShutdown initializes the server for shutdown.
func initServerForShutdown(port string, afterListening func()) {
	serverForShutdown = iris.New()
	serverForService.Logger().SetLevel("disable")
	serverForShutdown.Post("/close", closeHandler)
	err := serverForShutdown.Listen(":"+port, iris.WithoutStartupLog)
	if err != nil {
		core.Logger().Errorf("The port %s maybe used! Try to change another one!", port)
	}
	afterListening()
}

// closeHandler handles the service of closing the server.
func closeHandler(ctx iris.Context) {

	ctxBackground := context.Background()

	// Close the server for service.
	if serverForService != nil {
		err := serverForService.Shutdown(ctxBackground)
		if err != nil {
			core.Logger().Errorf("Failed to close server for service! Try to kill it? [%s].", err.Error())
			ctx.Write(models.ServerIsClosingResponse())
			return
		}
	}

	// Close the server for shutdown.
	if serverForShutdown != nil {
		defer func() {
			err := serverForShutdown.Shutdown(ctxBackground)
			if err != nil {
				core.Logger().Errorf("Failed to close server for shutdown! Try to kill it? [%s].", err.Error())
				os.Exit(0) // Return 0 if failed to close serverForShutdown.
			}
		}()
	}

	ctx.Write(models.ServerIsClosingResponse())
}
