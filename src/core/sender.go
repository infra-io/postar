// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/13 23:23:24

package core

import (
	"gopkg.in/gomail.v2"
)

// sender is for sending an email.
type sender struct {
	dialer *gomail.Dialer
}

// Send sends the email and returns an error if failed.
func (s *sender) Send(email *Email) error {

	// Create one message including information to be sent.
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.dialer.Username)
	msg.SetHeader("To", email.To)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody(email.ContentType, email.Body)

	// Dial and send this message.
	//s.dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return s.dialer.DialAndSend(msg)
}

// newSender returns a sender with given parameters.
func newSender(host string, port int, username string, password string) *sender {
	return &sender{
		dialer: gomail.NewDialer(host, port, username, password),
	}
}
