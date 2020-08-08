// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/27 00:00:54

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/avino-plan/postar/server/grpc/services"
	"google.golang.org/grpc"
)

func main() {

	// 初始化测试地址
	ip := "127.0.0.1"
	if len(os.Args) > 1 {
		ip = os.Args[1]
	}

	// Connect to the remote server.
	conn, err := grpc.Dial(ip+":5779", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create a new postar client.
	postarClient := services.NewPostarServiceClient(conn)

	// Send a request to server.
	resp, err := postarClient.Ping(context.Background(), &services.EmptyRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp.Data))

	// Send a request to server.
	resp, err = postarClient.Send(context.Background(), &services.SendTask{
		Email: &services.Email{
			To:          "fishinlove@163.com",
			Subject:     "grpc 测试 postar 运行情况",
			ContentType: "text/html; charset=utf-8",
			Body:        "<h1>哈喽！来自 <span style=\"color: #123456;\">postar<span> 的问候！</h1>",
		},
		Options: &services.SendOptions{Sync: false},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp.Data))

	// =====================================================================

	// Connect to the remote server.
	conn, err = grpc.Dial(ip+":5780", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create a new close client.
	closeClient := services.NewCloseServiceClient(conn)

	// Send a request to server.
	resp, err = closeClient.Close(context.Background(), &services.EmptyRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp.Data))
}
