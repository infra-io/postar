// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/08 23:37:24

package core

// Email is the struct represents of a message including all information for sending.
type Email struct {
	To          string `json:"to"`
	Subject     string `json:"subject"`
	ContentType string `json:"contentType"`
	Body        string `json:"body"`
}

// NewEmail returns an Email holder with given parameters.
func NewEmail(to string, subject string, contentType string, body string) *Email {
	return &Email{
		To:          to,
		Subject:     subject,
		ContentType: contentType,
		Body:        body,
	}
}
