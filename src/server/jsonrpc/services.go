// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/25 21:49:26

package jsonrpc

import (
	"net"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/models"
)

// Result represents the result of one call.
type Result struct {

	// Data is the result data.
	Data []byte `json:"data"`
}

// PostarService is the main service of postar.
type PostarService struct{}

// Send will do the sendTask and return error if failed.
// The result will be stored in sendTaskResult.
func (ps *PostarService) Send(sendTask *models.SendTask, result *Result) error {

	// Try to send this email.
	if sendTask.Options.Sync {
		err := core.SendSync(sendTask.Email)
		if err != nil {
			result.Data = models.FailedToSendEmailResponse()
			return err
		}
	} else {
		core.SendAsync(sendTask.Email)
	}

	// Successfully sent.
	result.Data = models.EmailSuccessfullySentResponse()
	return nil
}

// =================================== close service ===================================

// TODO document
type CloseRequest struct{}

type CloseService struct {
	target net.Listener
}

func NewCloseService(target net.Listener) *CloseService {
	return &CloseService{
		target: target,
	}
}

func (cs *CloseService) Close(request *CloseRequest, result *Result) error {

	// Close target.
	err := cs.target.Close()
	if err != nil {
		core.Logger().Errorf("Failed to close server for service! Try to kill it? [%s].", err.Error())
		result.Data = models.FailedToCloseServerResponse()
		return err
	}

	result.Data = models.ServerIsClosingResponse()
	return nil
}
