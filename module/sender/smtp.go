// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/11 23:34:02

package sender

import (
	"errors"
	"sync"
	"time"

	"github.com/avino-plan/postar/module"
	"gopkg.in/gomail.v2"
)

type SmtpSender struct {
	host        string
	port        int
	user        string
	sender      *gomail.Dialer
	messagePool *sync.Pool
}

func newSmtpSender() Sender {
	return &SmtpSender{
		messagePool: &sync.Pool{
			New: func() interface{} {
				return gomail.NewMessage()
			},
		},
	}
}

func (sms *SmtpSender) getMessage() *gomail.Message {
	return sms.messagePool.Get().(*gomail.Message)
}

func (sms *SmtpSender) releaseMessage(msg *gomail.Message) {
	msg.Reset()
	sms.messagePool.Put(msg)
}

func (sms *SmtpSender) Configure(config *module.Config) error {
	sms.host = config.Sender.Host
	sms.port = config.Sender.Port
	sms.user = config.Sender.User
	sms.sender = gomail.NewDialer(sms.host, sms.port, sms.user, config.Sender.Password)
	return nil
}

func (sms *SmtpSender) SendEmail(email *Email, options *SendOptions) error {

	if email == nil {
		return nil
	}

	if options == nil {
		opts := DefaultSendOptions()
		options = &opts
	}

	msg := sms.getMessage()
	defer sms.releaseMessage(msg)
	msg.SetHeader("From", sms.user)
	msg.SetHeader("To", email.To...)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody("text/plain;charset=utf-8", email.Content)
	module.Logger().Debug("before sending email").Any("email", email).Any("options", options).End()

	errCh := make(chan error, 1)
	go func() {
		module.Logger().Debug("1").End()
		errCh <- sms.sender.DialAndSend(msg)
		module.Logger().Debug("2").End()
	}()

	if options.Async {
		return nil
	}

	select {
	case err := <-errCh:
		return err
	case <-time.After(time.Duration(options.SendTimeout) * time.Millisecond):
		return errors.New("send timeout")
	}
}

func (sms *SmtpSender) Close() error {
	return nil
}
