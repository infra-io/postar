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
// Created at 2020/07/13 23:00:26

package system

import (
	"github.com/avino-plan/postar/src/models"
	"gopkg.in/gomail.v2"
)

func InitAllComponentsWith(config *models.Config) {
	initEmailDialerWith(config.Smtp.Host, config.Smtp.Port, config.Smtp.Username, config.Smtp.Password)
}

// SendEmail 可以发送一封邮件。
func SendEmail(email *models.Email) error {

	// 定义一个邮件信息
	msg := gomail.NewMessage()
	msg.SetHeader("From", emailDialer.Username)
	msg.SetHeader("To", email.To)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody(email.ContentType, email.Body)

	// 连接并发送
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return emailDialer.DialAndSend(msg)
}
