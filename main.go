// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/11 00:22:42

package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/gomail.v2"
)

func main() {

	host := "smtp.office365.com"
	port := 587
	user := os.Args[1]
	password := os.Args[2]

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		msg := gomail.NewMessage()
		msg.SetHeader("From", user)
		msg.SetHeader("To", "")
		msg.SetHeader("Subject", "go smtp 发邮件")
		msg.SetBody("text/plain;charset=UTF-8", "发邮件啦！！！！")

		dialer := gomail.NewDialer(host, port, user, password)
		err := dialer.DialAndSend(msg)
		if err != nil {
			log.Println(err)
			writer.Write([]byte(err.Error()))
			return
		}
	})

	err := http.ListenAndServe(":5897", nil)
	if err != nil {
		panic(err)
	}
}
