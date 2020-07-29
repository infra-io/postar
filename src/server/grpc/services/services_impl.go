// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/26 22:12:20

package services

import (
	"context"
	"time"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/models"
	"google.golang.org/grpc"
)

// PostarServiceImpl is the implement of PostarService.
type PostarServiceImpl struct{}

// Ping returns if postar is ready.
func (psi *PostarServiceImpl) Ping(ctx context.Context, request *EmptyRequest) (*Result, error) {
	return &Result{
		Data: []byte("Pong! Postar is ready! The version is " + core.Version + "."),
	}, nil
}

// Send will do the sendTask and return error if failed.
func (psi *PostarServiceImpl) Send(ctx context.Context, sendTask *SendTask) (*Result, error) {

	// Convert email object.
	email := &core.Email{
		To:          sendTask.Email.To,
		Subject:     sendTask.Email.Subject,
		ContentType: sendTask.Email.ContentType,
		Body:        sendTask.Email.Body,
	}

	// Try to send this email.
	if sendTask.Options.Sync {
		err := core.SendSync(email)
		if err != nil {
			core.Logger().Errorf("The error is %s. The information of sending task is {%+v, %+v}.", err.Error(), *sendTask.Email, *sendTask.Options)
			return &Result{
				Data: models.FailedToSendEmailResponse(),
			}, err
		}
	} else {
		core.SendAsync(email)
	}

	// Successfully sent.
	return &Result{
		Data: models.EmailSuccessfullySentResponse(),
	}, nil
}

// =================================== close service ===================================

// CloseServiceImpl is the close service of postar.
type CloseServiceImpl struct {

	// target is the server that will be closed by this service.
	target *grpc.Server

	// closeServer should be close "inelegantly".
	closeServer *grpc.Server
}

// NewCloseService returns a CloseServiceImpl holder with targets servers.
func NewCloseService(target *grpc.Server, closeServer *grpc.Server) *CloseServiceImpl {
	return &CloseServiceImpl{
		target:      target,
		closeServer: closeServer,
	}
}

// Close is the main method that CloseService provides.
func (sci *CloseServiceImpl) Close(ctx context.Context, request *EmptyRequest) (*Result, error) {

	// Stop the target server gracefully.
	sci.target.GracefulStop()

	// The close server should be delayed, and stopped "inelegantly".
	time.AfterFunc(3*time.Second, func() {
		sci.closeServer.Stop()
	})
	return &Result{
		Data: models.ServerIsClosingResponse(),
	}, nil
}
