// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 02:05:37

package service

import (
	"context"
	"time"

	"github.com/avino-plan/postar/internal/pkg/concurrency"
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
)

// SmtpServiceImpl is the service of smtp.
type SmtpServiceImpl struct {
	host     string            // The host of smtp server.
	port     int               // The port of smtp server.
	user     string            // The user of smtp server.
	password string            // The password of smtp server.
	pool     *concurrency.Pool // The pool of workers.
}

// NewSmtpService returns a new SmtpServer.
func NewSmtpService(host string, port int, user string, password string, pool *concurrency.Pool) SmtpService {
	return &SmtpServiceImpl{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		pool:     pool,
	}
}

// sendEmail sends email and returns an error if something wrong happens.
func (ss *SmtpServiceImpl) sendEmail(email *Email) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", ss.user)
	msg.SetHeader("To", email.To...)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody(email.BodyType, email.Body)
	return gomail.NewDialer(ss.host, ss.port, ss.user, ss.password).DialAndSend(msg)
}

// SendEmail sends email to somewhere.
func (ss *SmtpServiceImpl) SendEmail(ctx context.Context, email *Email, options *SendEmailOptions) error {
	if options == nil {
		options = DefaultSendEmailOptions()
	}

	errorCh := ss.pool.Go(ctx, func(ctx context.Context) error {
		return ss.sendEmail(email)
	})
	if options.Async {
		return nil
	}

	timer := time.NewTimer(options.Timeout)
	select {
	case err := <-errorCh:
		timer.Stop()
		return errors.Wrap(err, "send email failed")
	case <-timer.C:
		return errors.New("send email timeout")
	}
}
