// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"os"

	postarapi "github.com/avino-plan/api/go-out/postar"
	"google.golang.org/grpc"
)

func main() {
	// We recommend you to use client in:
	//
	// [Gitee](https://gitee.com/avino-plan/postar-client)
	// or
	// [Github](https://github.com/avino-plan/postar-client)
	//
	// However, you can use postar service directly in this way below.
	conn, err := grpc.Dial("127.0.0.1:5897", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	req := &postarapi.SendEmailRequest{
		Email: &postarapi.Email{
			Receivers: []string{os.Getenv("POSTAR_RECEIVER")},
			Subject:   "测试邮件",
			BodyType:  "text/html",
			Body:      "<p>邮件内容</p>",
		},
		Options: nil,
	}
	fmt.Printf("client req: %+v\n", req)

	client := postarapi.NewPostarServiceClient(conn)
	rsp, err := client.SendEmail(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("server rsp: %+v\n", rsp)
}
