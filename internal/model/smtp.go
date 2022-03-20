// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

import (
	"time"

	"github.com/avino-plan/postar/configs"
)

type Email struct {
	Subject   string   // Subject.
	Receivers []string // Receivers.
	BodyType  string   // Body type.
	Body      string   // Body.
}

func NewEmail() *Email {
	return new(Email)
}

// SendEmailOptions is the options of sending one email.
type SendEmailOptions struct {
	Async   bool          // The mode of sending one email.
	Timeout time.Duration // The timeout of sending one email.
}

// DefaultSendEmailOptions returns a default options for sending emails.
func DefaultSendEmailOptions(c *configs.Config) *SendEmailOptions {
	if c == nil {
		return &SendEmailOptions{
			Async:   false,
			Timeout: 10 * time.Second,
		}
	}

	return &SendEmailOptions{
		Async:   c.TaskAsync(),
		Timeout: c.TaskTimeout(),
	}
}
