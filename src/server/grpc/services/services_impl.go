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

	// targets are the servers that will be closed by this service.
	targets []*grpc.Server
}

// NewCloseService returns a CloseServiceImpl holder with targets servers.
func NewCloseService(targets ...*grpc.Server) *CloseServiceImpl {
	return &CloseServiceImpl{
		targets: targets,
	}
}

// Close is the main method that CloseService provides.
func (sci *CloseServiceImpl) Close(ctx context.Context, request *EmptyRequest) (*Result, error) {
	for _, target := range sci.targets {
		target.GracefulStop()
	}
	return &Result{
		Data: models.ServerIsClosingResponse(),
	}, nil
}
