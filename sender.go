// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/11 23:34:02

package main

import (
	"errors"
	"sync"
	"time"

	"gopkg.in/gomail.v2"
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
	SendEmail(email *Email, options *SendOptions) error
}

type SmtpSender struct {
	host        string
	port        int
	user        string
	sender      *gomail.Dialer
	messagePool *sync.Pool
}

func NewSmtpSender(host string, port int, user string, password string) *SmtpSender {

	return &SmtpSender{
		host:   host,
		port:   port,
		user:   user,
		sender: gomail.NewDialer(host, port, user, password),
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
	Logger().Debug("before sending email").Any("email", email).Any("options", options).End()

	errCh := make(chan error, 1)
	go func() {
		Logger().Debug("1").End()
		errCh <- sms.sender.DialAndSend(msg)
		Logger().Debug("2").End()
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
