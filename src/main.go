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
// Created at 2020/07/08 23:41:46
package main

import (
	"strconv"

	"github.com/FishGoddess/logit"
	"github.com/avino-plan/postar/src/config"
	"github.com/avino-plan/postar/src/handlers"
	"github.com/avino-plan/postar/src/models"
	"github.com/avino-plan/postar/src/system"
	"github.com/kataras/iris/v12"
)

func main() {

	config.UseConfig(func(config *models.Config) {
		system.InitAllComponentsWith(config)
	})

	app := iris.New()
	app.Get("/ping", handlers.PingHandler)
	app.Post("/send", handlers.SendHandler)

	port := strconv.Itoa(5779)
	logit.Infof("Postar is running at port %s.", port)
	app.Listen(":" + port)
}
