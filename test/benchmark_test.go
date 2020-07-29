// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/29 00:02:13

package main

import (
	"bytes"
	"context"
	"net/http"
	stdJsonRPC "net/rpc/jsonrpc"
	"testing"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/models"
	"github.com/avino-plan/postar/src/server/grpc/services"
	"github.com/avino-plan/postar/src/server/jsonrpc"
	"google.golang.org/grpc"
)

// Benchmark the http server.
// BenchmarkHttpServer-8               3165            386013 ns/op           14838 B/op         89 allocs/op
func BenchmarkHttpServer(b *testing.B) {

	body := []byte(`
{
 "email": {
   "to": "fishinlove@163.com",
   "subject": "测试 postar 运行情况",
   "contentType": "text/html; charset=utf-8",
   "body": "<h1>哈喽！来自 <span style=\"color: #123456;\">postar<span> 的问候！</h1>"
 },
 "options": {
   "sync": true
 }
}
   `)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := http.Post("http://127.0.0.1:5779/send", "application/json; charset=utf-8", bytes.NewBuffer(body))
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}

// Benchmark the jsonrpc server.
// BenchmarkJsonRPCServer-8           17462             69567 ns/op             712 B/op         15 allocs/op
func BenchmarkJsonRPCServer(b *testing.B) {

	// Connect to the remote server.
	postarConn, err := stdJsonRPC.Dial("tcp", "127.0.0.1:5779")
	if err != nil {
		panic(err)
	}
	defer postarConn.Close()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		sendTask := models.NewSendTaskWithDefaultOptions()
		sendTask.Email = &core.Email{
			To:          "fishinlove@163.com",
			Subject:     "jsonrpc 测试 postar 运行情况",
			ContentType: "text/html; charset=utf-8",
			Body:        "<h1>哈喽！来自 <span style=\"color: #123456;\">postar<span> 的问候！</h1>",
		}
		sendTask.Options.Sync = true

		resp := &jsonrpc.Result{}
		err = postarConn.Call("PostarService.Send", sendTask, resp)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark the grpc server.
// BenchmarkGRPCServer-8              10000            132845 ns/op            5248 B/op         98 allocs/op
func BenchmarkGRPCServer(b *testing.B) {

	conn, err := grpc.Dial("127.0.0.1:5779", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	postarClient := services.NewPostarServiceClient(conn)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := postarClient.Send(context.Background(), &services.SendTask{
			Email: &services.Email{
				To:          "fishinlove@163.com",
				Subject:     "grpc 测试 postar 运行情况",
				ContentType: "text/html; charset=utf-8",
				Body:        "<h1>哈喽！来自 <span style=\"color: #123456;\">postar<span> 的问候！</h1>",
			},
			Options: &services.SendOptions{Sync: false},
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}
