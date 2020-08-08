// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/12 23:49:43

package models

import "github.com/avino-plan/postar/core"

// SendTask is the struct represents of all information of sending task.
type SendTask struct {

	// Email is the email which will be sent.
	Email *core.Email `json:"email"`

	// Options are some settings of sending task.
	Options *SendOptions `json:"options"`
}

// sendOptions are some settings of sending task.
type SendOptions struct {

	// Sync means the send task is synchronous, default is asynchronous.
	Sync bool `json:"sync"`
}

// NewEmptySendTask returns an empty SendTask holder.
func NewSendTaskWithDefaultOptions() *SendTask {
	return &SendTask{
		Options: &SendOptions{
			Sync: false,
		},
	}
}
