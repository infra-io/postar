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

	"github.com/avino-plan/postar/pkg/concurrency"
	"gopkg.in/gomail.v2"
)

// smtpService is the service of smtp.
type smtpService struct {
	host     string            // The host of smtp server.
	port     int               // The port of smtp server.
	user     string            // The user of smtp server.
	password string            // The password of smtp server.
	pool     *concurrency.Pool // The pool of workers.
}

// NewSmtpService returns a new SmtpServer.
func NewSmtpService(host string, port int, user string, password string, pool *concurrency.Pool) SmtpService {
	return &smtpService{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		pool:     pool,
	}
}

// sendEmail sends email and returns an error if something wrong happens.
func (ss *smtpService) sendEmail(email *Email) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", ss.user)
	msg.SetHeader("To", email.To...)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody(email.BodyType, email.Body)
	return gomail.NewDialer(ss.host, ss.port, ss.user, ss.password).DialAndSend(msg)
}

// SendEmail send
func (ss *smtpService) SendEmail(ctx context.Context, email *Email, options *SendEmailOptions) error {

	done := make(chan error, 1)
	err := ss.pool.Go(ctx, func(ctx context.Context) {
		done <- ss.sendEmail(email)
	})
	if err != nil && concurrency.IsEnqueueTimeout(err) {
		// TODO wraps err
		return err
	}

	if options.Async {
		return nil
	}

	timer := time.NewTimer(options.Timeout)
	select {
	case err = <-done:
		return err
	case <-timer.C:
		// TODO timeout error
		return nil
	}
}
