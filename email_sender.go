// 邮件发送工具箱
//
// Author: Fish
// Email: fishgoddess@qq.com
// Created by 2019/07/11 10:24:38
package helper

import (
	"log"

	"gopkg.in/gomail.v2"
)

// 邮件结构体
type Email struct {
	To          string
	Subject     string
	ContentType string
	Body        string
}

// 创建邮件对象
func NewEmail(to string, subject string, contentType string, body string) Email {
	return Email{
		To:          to,
		Subject:     subject,
		ContentType: contentType,
		Body:        body,
	}
}

// 创建 HTML 类型邮件
func NewHtmlEmail(to string, subject string, body string) Email {
	return NewEmail(to, subject, "text/html;charset=utf-8", body)
}

// 发送邮件
func SendEmail(email Email) {
	go sendEmailInternal(email)
}

// 内部发送邮件的方法
func sendEmailInternal(email Email) {

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
