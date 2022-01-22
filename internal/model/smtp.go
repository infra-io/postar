// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/17 00:06:09

package model

import "time"

// Email is an email.
type Email struct {
	To       []string // Receivers.
	Subject  string   // Subject.
	BodyType string   // Body type.
	Body     string   // Body.
}

// SendEmailOptions is the options of sending one email.
type SendEmailOptions struct {
	Async   bool          // The mode of sending one email.
	Timeout time.Duration // The timeout of sending one email.
}

// DefaultSendEmailOptions returns a default options for sending emails.
func DefaultSendEmailOptions() *SendEmailOptions {
	return &SendEmailOptions{
		Async:   false,
		Timeout: 5 * time.Second,
	}
}
