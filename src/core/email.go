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
// Created at 2020/07/08 23:37:24

package core

// Email is the struct represents of a message including all information for sending.
type Email struct {
	To          string
	Subject     string
	ContentType string
	Body        string
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
