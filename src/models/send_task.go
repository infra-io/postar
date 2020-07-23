// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/12 23:49:43

package models

// SendTask is the struct represents of all information of sending task.
type SendTask struct {
	To          string `json:"to"`
	Subject     string `json:"subject"`
	ContentType string `json:"contentType"`
	Body        string `json:"body"`
}

// NewEmptySendTask returns an empty SendTask holder.
func NewEmptySendTask() *SendTask {
	return &SendTask{}
}
