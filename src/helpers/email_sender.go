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
// Created at 2020/07/08 23:36:30

package helpers

import (
	"log"

	"github.com/avino-plan/postar/src/models"
	"gopkg.in/gomail.v2"
)

// 发送邮件
func SendEmail(email *models.Email) {
	go sendEmailInternal(email)
}

// 内部发送邮件的方法
func sendEmailInternal(email *models.Email) {

	// 定义一个邮件信息
	msg := gomail.NewMessage()
	msg.SetHeader("From", "smtp.user")
	msg.SetHeader("To", email.To)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody(email.ContentType, email.Body)

	// 连接并发送
	d := gomail.NewDialer("smtp.host", 537, "smtp.user", "smtp.password")
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(msg)
	if err != nil {
		log.Printf("邮件发送失败！邮件信息是：%s，错误原因：%s\n", email, err)
	}
}
