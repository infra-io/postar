// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package biz

import (
	"context"

	"github.com/go-logit/logit"

	"github.com/FishGoddess/errors"
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/model"
	pkgerrors "github.com/avinoplan/postar/pkg/errors"
	"github.com/panjf2000/ants/v2"
	"gopkg.in/gomail.v2"
)

// SMTPBiz is the biz of smtp.
type SMTPBiz struct {
	c    *configs.Config
	pool *ants.Pool // The pool of workers.
}

// NewSMTPBiz returns a new SMTPBiz.
func NewSMTPBiz(c *configs.Config, pool *ants.Pool) *SMTPBiz {
	return &SMTPBiz{
		c:    c,
		pool: pool,
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
	if email == nil {
		err := errors.New("email is nil")
		logit.Error("email is nil").Error("err", err).WithContext(ctx).End()
		return errors.BadRequest(err)
	}

	if options == nil {
		options = model.DefaultSendEmailOptions(sb.c)
		logit.Debug("options is nil, using default options").Any("options", options).WithContext(ctx).End()
	}

	ctx, cancel := context.WithTimeout(ctx, options.Timeout)
	defer cancel()

	errorCh := make(chan error, 1)
	err := sb.pool.Submit(func() {
		defer close(errorCh)
		errorCh <- sb.sendEmail(email)
	})

	if err != nil {
		logit.Error("submit email sending task to pool failed").Error("err", err).Any("email", email).WithContext(ctx).End()
		return pkgerrors.SendEmailFailedErr(err)
	}

	if options.Async {
		return nil
	}

	select {
	case err = <-errorCh:
		if err != nil {
			logit.Error("send email failed").Error("err", err).Any("email", email).WithContext(ctx).End()
			return pkgerrors.SendEmailFailedErr(err)
		}
	case <-ctx.Done():
		err = ctx.Err()
		if err != nil {
			logit.Error("send email timeout").Error("err", err).Any("email", email).WithContext(ctx).End()
			return errors.Timeout(err)
		}
	}

	return nil
}
