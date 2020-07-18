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
// Created at 2020/07/13 00:42:12

package http

import (
	"context"
	"os"
	"time"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/models"
	"github.com/kataras/iris/v12"
)

// closeHandler handles the service of closing the server.
func closeHandler(ctx iris.Context) {

	ctxBackground := context.Background()

	// Close the server for service.
	if serverForService != nil {
		err := serverForService.Shutdown(ctxBackground)
		if err != nil {
			core.Logger().Errorf("Failed to close server for service! Try to kill it? [%s].", err.Error())
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Write(models.FailedToCloseServerResponse())
			return
		}
	}

	core.Logger().Info("Server for service has been closed! Have a good day :)")
	ctx.Write(models.ServerIsClosingResponse())
	ctx.ResponseWriter().Flush()

	// Close the server for shutdown.
	if serverForShutdown != nil {
		// This server should be closed after finishing flushing all data to clients, or
		// you will get a Connection Reset error, so do it after 3 seconds.
		time.AfterFunc(3*time.Second, func() {
			err := serverForShutdown.Shutdown(ctxBackground)
			if err != nil {
				core.Logger().Errorf("Failed to close server for shutdown! Exit with code 0. [%s].", err.Error())
				os.Exit(0) // Return 0 if failed to close serverForShutdown.
			}
		})
	}
}

// pingHandler tells you if postar is ready or not.
func pingHandler(ctx iris.Context) {
	ctx.Write([]byte(`<body style="text-align: center;"><h1>Pong!</h1><h3>- Postar is ready! -</h3><p>- The version is ` + core.Version + ` -</p></body>`))
}

// sendHandler handles the service of sending emails.
func sendHandler(ctx iris.Context) {

	// Parse send task from request.
	sendTask := models.NewEmptySendTask()
	err := ctx.ReadJSON(&sendTask)
	if err != nil {
		core.Logger().Errorf("The error is %s.", err.Error())
		ctx.StatusCode(400)
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.Write(models.WrongRequestBodyResponse())
		return
	}

	// Try to send this email.
	email := core.NewEmail(sendTask.To, sendTask.Subject, sendTask.ContentType, sendTask.Body)
	err = core.Send(email)
	if err != nil {
		core.Logger().Errorf("The error is %s. The information of this email is %+v.", err.Error(), sendTask)
		ctx.StatusCode(500)
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.Write(models.FailedToSendEmailResponse())
		return
	}

	// Successfully.
	core.Logger().Debugf("Email %+v successfully sent.", email)
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Write(models.EmailSuccessfullySentResponse())
}
