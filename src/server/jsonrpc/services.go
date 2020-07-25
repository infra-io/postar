// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/25 21:49:26

package jsonrpc

import (
	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/models"
)

// SendTaskResult represents the result of a sending task.
type SendTaskResult struct {

	// result is the result data.
	result []byte
}

// PostarService is the main service of postar.
type PostarService struct{}

// Send will do the sendTask and return error if failed.
// The result will be stored in sendTaskResult.
func (ps *PostarService) Send(sendTask *models.SendTask, sendTaskResult *SendTaskResult) error {

	// Try to send this email.
	if sendTask.Options.Sync {
		err := core.SendSync(sendTask.Email)
		if err != nil {
			sendTaskResult.result = models.FailedToSendEmailResponse()
			return err
		}
	} else {
		core.SendAsync(sendTask.Email)
	}

	// Successfully sent.
	sendTaskResult.result = models.EmailSuccessfullySentResponse()
	return nil
}
