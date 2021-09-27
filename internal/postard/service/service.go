// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/17 00:06:09

package service

import (
	"context"
	"time"
)

// Email is an email.
type Email struct {
	To       []string // The receivers of one email.
	Subject  string   // The subject of one email.
	BodyType string   // The content type of body.
	Body     string   // The body of one email.
}

// SendEmailOptions is the options of sending one email.
type SendEmailOptions struct {
	Async   bool          // The mode of sending one email.
	Timeout time.Duration // The timeout of sending one email.
}

// ContextService is the service of context
type ContextService interface {
	// WrapContext wraps context with something and returns a new context.
	WrapContext(ctx context.Context) context.Context
}

// SmtpService is the service of smtp.
type SmtpService interface {
	// SendEmail sends email with options and returns an error if something wrong happens.
	SendEmail(ctx context.Context, email *Email, options *SendEmailOptions) error
}

// DefaultSendEmailOptions returns a default options for sending emails.
func DefaultSendEmailOptions() *SendEmailOptions {
	return &SendEmailOptions{
		Async:   false,
		Timeout: 5 * time.Second,
	}
}
