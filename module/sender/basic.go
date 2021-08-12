// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/12 23:20:22

package sender

import "github.com/avino-plan/postar/module"

var (
	senders = map[string]func() Sender{
		"smtp": newSmtpSender,
	}
)

type Email struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Content string   `json:"content"`
}

type SendOptions struct {
	Async       bool `json:"async"`
	SendTimeout int  `json:"sendTimeout"` // 发送超时，单位为 ms
}

func DefaultSendOptions() SendOptions {
	return SendOptions{
		Async:       false,
		SendTimeout: 5000,
	}
}

type Sender interface {
	module.Configurer
	SendEmail(email *Email, options *SendOptions) error
	Close() error
}

func Initialize(config *module.Config) (Sender, error) {
	sender := senders[config.Global.SenderType]()
	return sender, sender.Configure(config)
}
