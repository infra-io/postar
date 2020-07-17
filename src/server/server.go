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
// Created at 2020/07/15 23:44:43

package server

import (
	"sync"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/server/http"
)

var (
	// servers stores all servers that can be used.
	servers = map[string]func(port string, closedPort string) *sync.WaitGroup{
		"http": http.InitServer,
	}
)

// RunServer runs a server for service and shutdown.
// Notice that the returning value is *sync.WaitGroup, so you can use it to
// block your main goroutine before closing the server.
func RunServer() *sync.WaitGroup {
	initServer, ok := servers[core.ServerType()]
	if !ok {
		core.Logger().Errorf("The server type %s doesn't exist! Try 'http'?", core.ServerType())
	}
	return initServer(core.ServerPort(), core.ServerClosedPort())
}
