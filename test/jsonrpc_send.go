// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/26 00:55:32

package main

import (
	"fmt"
	stdJsonRPC "net/rpc/jsonrpc"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/models"
	"github.com/avino-plan/postar/src/server/jsonrpc"
)

// TestServerImpl_Stop tests ServerImpl.Stop.
func main() {

	// Connect to the remote server.
	postarConn, err := stdJsonRPC.Dial("tcp", "127.0.0.1:5779")
	if err != nil {
		panic(err)
	}
	defer postarConn.Close()

	resp := &jsonrpc.Result{}
	err = postarConn.Call("PostarService.Ping", &jsonrpc.EmptyRequest{}, resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp.Data))

	// Send a request to server.
	sendTask := models.NewSendTaskWithDefaultOptions()
	sendTask.Email = &core.Email{
		To:          "fishinlove@163.com",
		Subject:     "jsonrpc 测试 postar 运行情况",
		ContentType: "text/html; charset=utf-8",
		Body:        "<h1>哈喽！来自 <span style=\"color: #123456;\">postar<span> 的问候！</h1>",
	}
	sendTask.Options.Sync = true
	err = postarConn.Call("PostarService.Send", sendTask, resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp.Data))

	// ============================================================

	// Connect to the remote server.
	closeConn, err := stdJsonRPC.Dial("tcp", "127.0.0.1:5780")
	if err != nil {
		panic(err)
	}
	defer closeConn.Close()

	// Send a request to server.
	req := &jsonrpc.EmptyRequest{}
	err = closeConn.Call("CloseService.Close", req, resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp.Data))
}
