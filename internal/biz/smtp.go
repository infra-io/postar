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
	"github.com/avinoplan/postar/internal/model"
	"github.com/avinoplan/postar/pkg/errors"
	"github.com/panjf2000/ants/v2"
	"gopkg.in/gomail.v2"
)

// SmtpBiz is the biz of smtp.
type SmtpBiz struct {
	pool     *ants.Pool // The pool of workers.
	host     string     // The host of smtp server.
	port     int        // The port of smtp server.
	user     string     // The user of smtp server.
	password string     // The password of smtp server.
}

// NewSmtpBiz returns a new SmtpBiz.
func NewSmtpBiz(pool *ants.Pool, host string, port int, user string, password string) *SmtpBiz {
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
		logger.Debug("options is nil, using default options").Any("options", options).End()
	}

	ctx, cancel := context.WithTimeout(ctx, options.Timeout)
	defer cancel()

	errorCh := make(chan error, 1)
	err := sb.pool.Submit(func() {
		defer close(errorCh)
		errorCh <- sb.sendEmail(email)
	})

	if err != nil {
		logger.Error("submit email sending task to pool failed").Error("err", err).End()
		return errors.SendEmailFailedErr(err)
	}

	if options.Async {
		return nil
	}

	select {
	case err = <-errorCh:
		if err != nil {
			logger.Error("send email failed").Error("err", err).End()
			return errors.SendEmailFailedErr(err)
		}
	case <-ctx.Done():
		err = ctx.Err()
		if err != nil {
			logger.Error("send email timeout").Error("err", err).End()
			return errors.TimeoutErr(err)
		}
	}

	return nil
}
