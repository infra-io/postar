// Copyright 2020 Ye Zi Jie. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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

	// Create one message including information for sending.
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
