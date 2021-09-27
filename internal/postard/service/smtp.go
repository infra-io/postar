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

	"github.com/FishGoddess/logit"
	"github.com/avino-plan/postar/internal/pkg/concurrency"
	"gopkg.in/gomail.v2"
)

// smtpServiceImpl is the service of smtp.
type smtpServiceImpl struct {
	pool     *concurrency.Pool // The pool of workers.
	host     string            // The host of smtp server.
	port     int               // The port of smtp server.
	user     string            // The user of smtp server.
	password string            // The password of smtp server.
}

// NewSmtpService returns a new SmtpService.
func NewSmtpService(pool *concurrency.Pool, host string, port int, user string, password string) SmtpService {
	return &smtpServiceImpl{
		pool:     pool,
		host:     host,
		port:     port,
		user:     user,
		password: password,
	}
}

// sendEmail sends email and returns an error if something wrong happens.
func (ss *smtpServiceImpl) sendEmail(email *Email) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", ss.user)
	msg.SetHeader("To", email.To...)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody(email.BodyType, email.Body)
	return gomail.NewDialer(ss.host, ss.port, ss.user, ss.password).DialAndSend(msg)
}

// SendEmail sends email to somewhere.
func (ss *smtpServiceImpl) SendEmail(ctx context.Context, email *Email, options *SendEmailOptions) error {
	logger := logit.FromContext(ctx)

	if options == nil {
		options = DefaultSendEmailOptions()
		logger.Debug("options is nil, using DefaultSendEmailOptions()").Any("options", options).End()
	}

	ctx, cancel := context.WithTimeout(ctx, options.Timeout)
	defer cancel()

	errorCh := ss.pool.Go(ctx, func(ctx context.Context) error { return ss.sendEmail(email) })
	if options.Async {
		return nil
	}

	select {
	case e := <-errorCh:
		if e != nil {
			logger.Error("send email failed").Error("e", e).End()
			return errSendEmailFailed
		}
	case <-ctx.Done():
		e := ctx.Err()
		if e != nil {
			logger.Error("send email timeout").Error("e", e).End()
			return errSendTimeout
		}
	}
	return nil
}
