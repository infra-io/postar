// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 02:05:37

package biz

import (
	"context"

	"github.com/FishGoddess/logit"
	"github.com/avino-plan/postar/internal/postard/model"
	"github.com/avino-plan/postar/pkg/concurrency"
	"github.com/avino-plan/postar/pkg/errors"
	"gopkg.in/gomail.v2"
)

// SmtpBiz is the biz of smtp.
type SmtpBiz struct {
	pool     *concurrency.Pool // The pool of workers.
	host     string            // The host of smtp server.
	port     int               // The port of smtp server.
	user     string            // The user of smtp server.
	password string            // The password of smtp server.
}

// NewSmtpBiz returns a new SmtpBiz.
func NewSmtpBiz(pool *concurrency.Pool, host string, port int, user string, password string) *SmtpBiz {
	return &SmtpBiz{
		pool:     pool,
		host:     host,
		port:     port,
		user:     user,
		password: password,
	}
}

// sendEmail sends email and returns an error if something wrong happens.
func (sb *SmtpBiz) sendEmail(email *model.Email) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", sb.user)
	msg.SetHeader("To", email.To...)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody(email.BodyType, email.Body)
	return gomail.NewDialer(sb.host, sb.port, sb.user, sb.password).DialAndSend(msg)
}

// SendEmail sends email to somewhere.
func (sb *SmtpBiz) SendEmail(ctx context.Context, email *model.Email, options *model.SendEmailOptions) error {
	logger := logit.FromContext(ctx)

	if options == nil {
		options = model.DefaultSendEmailOptions()
		logger.Debug("options is nil, using DefaultSendEmailOptions()").Any("options", options).End()
	}

	ctx, cancel := context.WithTimeout(ctx, options.Timeout)
	defer cancel()

	errorCh := sb.pool.Go(ctx, func(ctx context.Context) error { return sb.sendEmail(email) })
	if options.Async {
		return nil
	}

	select {
	case e := <-errorCh:
		if e != nil {
			logger.Error("send email failed").Error("e", e).End()
			return errors.SendEmailFailedErr(e)
		}
	case <-ctx.Done():
		e := ctx.Err()
		if e != nil {
			logger.Error("send email timeout").Error("e", e).End()
			return errors.SendTimeoutErr(e)
		}
	}
	return nil
}
