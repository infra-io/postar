// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/12 23:20:22

package sender

import (
	"errors"
	"fmt"

	"github.com/avino-plan/postar/module"
)

var (
	timeoutErr = errors.New("timeout")

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
	Async   bool `json:"async"`
	Timeout int  `json:"timeout"` // 发送超时，单位为 ms
}

func DefaultSendOptions() SendOptions {
	return SendOptions{
		Async:   false,
		Timeout: 10000,
	}
}

type Sender interface {
	module.Configurer
	SendEmail(email *Email, options *SendOptions) error
	Close() error
}

func Initialize(config *module.Config) (Sender, error) {

	newSender, ok := senders[config.Global.SenderType]
	if !ok {
		return nil, fmt.Errorf("sender type %s not found", config.Global.SenderType)
	}

	sender := newSender()
	return sender, sender.Configure(config)
}

func IsTimeout(err error) bool {
	return err == timeoutErr
}
