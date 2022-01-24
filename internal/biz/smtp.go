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
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/model"
	"github.com/avinoplan/postar/pkg/errors"
	"github.com/panjf2000/ants/v2"
	"gopkg.in/gomail.v2"
)

// SMTPBiz is the biz of smtp.
type SMTPBiz struct {
	c      *configs.Config
	logger *logit.Logger
	pool   *ants.Pool // The pool of workers.
}

// NewSMTPBiz returns a new SMTPBiz.
func NewSMTPBiz(c *configs.Config, logger *logit.Logger, pool *ants.Pool) *SMTPBiz {
	return &SMTPBiz{
		c:      c,
		logger: logger,
		pool:   pool,
	}
}

// sendEmail sends email and returns an error if something wrong happens.
func (sb *SMTPBiz) sendEmail(email *model.Email) error {
	msg := gomail.NewMessage()
	msg.SetHeader("Subject", email.Subject)
	msg.SetHeader("From", sb.c.SMTPUser())
	msg.SetHeader("To", email.Receivers...)
	msg.SetBody(email.BodyType, email.Body)
	return gomail.NewDialer(sb.c.SMTPHost(), sb.c.SMTPPort(), sb.c.SMTPUser(), sb.c.SMTPPassword()).DialAndSend(msg)
}

// SendEmail sends email to somewhere.
func (sb *SMTPBiz) SendEmail(ctx context.Context, email *model.Email, options *model.SendEmailOptions) error {
	logger := logit.FromContext(ctx)

	if options == nil {
		options = model.DefaultSendEmailOptions(sb.c)
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
