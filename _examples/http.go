// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 02:05:02

package main

import (
	"bytes"
	"fmt"
	"github.com/avinoplan/postar/api"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "http://127.0.0.1:5897/sendEmail"

	emailReq := &api.SendEmailRequest{
		Email:   &api.Email{
			Receivers: []string{os.Getenv("POSTAR_RECEIVER")},
			Subject:   "测试邮件",
			BodyType:  "text/html",
			Body:      "<p>邮件内容</p>",
		},
		Options: nil,
	}
	fmt.Printf("client req: %+v\n", emailReq)

	marshaled, err := proto.Marshal(emailReq)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(url, "application/octet-stream", bytes.NewReader(marshaled))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	emailRsp := new(api.SendEmailResponse)
	err = proto.Unmarshal(body, emailRsp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("server rsp: %+v\n", emailRsp)
}
