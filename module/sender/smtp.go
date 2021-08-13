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

type smtpRequest struct {
	message *gomail.Message
	errorCh chan error
}

func (sr *smtpRequest) reset() {
	sr.message.Reset()
	sr.errorCh = make(chan error, 1)
}

type SmtpSender struct {
	host        string
	port        int
	user        string
	dialer      *gomail.Dialer
	requestCh   chan *smtpRequest
	requestPool *sync.Pool
}

func newSmtpSender() Sender {
	return &SmtpSender{
		requestPool: &sync.Pool{
			New: func() interface{} {
				return &smtpRequest{
					message: gomail.NewMessage(),
					errorCh: make(chan error, 1),
				}
			},
		},
	}
}

func (sms *SmtpSender) newRequest() *smtpRequest {
	return sms.requestPool.Get().(*smtpRequest)
}

func (sms *SmtpSender) releaseRequest(request *smtpRequest) {
	request.reset()
	sms.requestPool.Put(request)
}

func (sms *SmtpSender) Configure(config *module.Config) error {

	sms.host = config.Sender.Host
	sms.port = config.Sender.Port
	sms.user = config.Sender.User
	sms.dialer = gomail.NewDialer(sms.host, sms.port, sms.user, config.Sender.Password)
	sms.requestCh = make(chan *smtpRequest, config.Sender.RequestChannelSize)

	for i := 0; i < config.Sender.WorkerNumber; i++ {
		go func() {
			for request := range sms.requestCh {
				request.errorCh <- sms.dialer.DialAndSend(request.message)
				sms.releaseRequest(request)
			}
		}()
	}
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

	request := sms.newRequest()
	request.message.SetHeader("From", sms.user)
	request.message.SetHeader("To", email.To...)
	request.message.SetHeader("Subject", email.Subject)
	request.message.SetBody("text/plain;charset=utf-8", email.Content)
	module.Logger().Debug("before sending email").Any("email", email).Any("options", options).End()

	sms.requestCh <- request
	if options.Async {
		return nil
	}

	select {
	case err := <-request.errorCh:
		return err
	case <-time.After(time.Duration(options.Timeout) * time.Millisecond):
		return errors.New("send timeout")
	}
}

func (sms *SmtpSender) Close() error {
	close(sms.requestCh)
	return nil
}
