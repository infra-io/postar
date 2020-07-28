// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
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

// closeHandlerOf returns a handler which is for closing the server.
func newCloseHandler(si *serverImpl) func(ctx iris.Context) {
	return func(ctx iris.Context) {

		ctxBackground := context.Background()

		// Close the server for service.
		if si.serverForService != nil {
			err := si.serverForService.Shutdown(ctxBackground)
			if err != nil {
				core.Logger().Errorf("Failed to close server for service! Try to kill it? [%s].", err.Error())
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Write(models.FailedToCloseServerResponse())
				return
			}
		}

		ctx.Write(models.ServerIsClosingResponse())
		ctx.ResponseWriter().Flush()

		// Close the server for shutdown.
		if si.serverForShutdown != nil {
			// This server should be closed after finishing flushing all data to clients, or
			// you will get a Connection Reset error, so do it after 3 seconds.
			time.AfterFunc(3*time.Second, func() {
				err := si.serverForShutdown.Shutdown(ctxBackground)
				if err != nil {
					core.Logger().Errorf("Failed to close server for shutdown! Exit with code 0. [%s].", err.Error())
					os.Exit(0) // Return 0 if failed to close serverForShutdown.
				}
			})
		}
	}
}

// newPingHandler returns a handler which tells you if postar is ready.
func newPingHandler() func(ctx iris.Context) {
	return func(ctx iris.Context) {
		ctx.Write([]byte(`<body style="text-align: center;"><h1>Pong!</h1><h3>- Postar is ready! -</h3><p>- The version is ` + core.Version + ` -</p></body>`))
	}
}

// newSendHandler returns a handler which is for sending emails.
func newSendHandler() func(ctx iris.Context) {
	return func(ctx iris.Context) {

		// Parse send task from request.
		sendTask := models.NewSendTaskWithDefaultOptions()
		err := ctx.ReadJSON(&sendTask)
		if err != nil {
			core.Logger().Errorf("The error is %s.", err.Error())
			ctx.StatusCode(400)
			ctx.Header("Content-Type", "application/json; charset=utf-8")
			ctx.Write(models.WrongRequestBodyResponse())
			return
		}

		// Try to send this email.
		if sendTask.Options.Sync {
			err = core.SendSync(sendTask.Email)
			if err != nil {
				core.Logger().Errorf("The error is %s. The information of sending task is %+v.", err.Error(), sendTask)
				ctx.StatusCode(500)
				ctx.Header("Content-Type", "application/json; charset=utf-8")
				ctx.Write(models.FailedToSendEmailResponse())
				return
			}
		} else {
			core.SendAsync(sendTask.Email)
		}

		// Successfully sent.
		core.Logger().Debugf("Email %+v successfully sent.", sendTask.Email)
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.Write(models.EmailSuccessfullySentResponse())
	}
}
