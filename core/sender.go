// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/13 23:23:24

package core

import (
	"sync"

	"gopkg.in/gomail.v2"
)

// sender is for sending an email.
type sender struct {

	// dialer is the real sender inside.
	dialer *gomail.Dialer

	// messagePool can reuse message holders.
	messagePool *sync.Pool
}

// getMessage returns a message for use.
func (s *sender) getMessage() *gomail.Message {
	return s.messagePool.Get().(*gomail.Message)
}

// putMessage resets and releases the msg.
func (s *sender) putMessage(msg *gomail.Message) {
	msg.Reset()
	s.messagePool.Put(msg)
}

// Send sends the email and returns an error if failed.
func (s *sender) Send(email *Email) error {

	// Create one message including information to be sent.
	msg := s.getMessage()
	defer s.putMessage(msg)

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
		messagePool: &sync.Pool{
			New: func() interface{} {
				return gomail.NewMessage()
			},
		},
	}
}
