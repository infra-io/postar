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
	"net/rpc/jsonrpc"
)

// SendTask is the struct represents of all information of sending task.
type SendTask struct {

	// Email is the email which will be sent.
	Email *Email `json:"email"`

	// Options are some settings of sending task.
	Options *SendOptions `json:"options"`
}

// Email is the struct represents of a message including all information for sending.
type Email struct {
	To          string `json:"to"`
	Subject     string `json:"subject"`
	ContentType string `json:"contentType"`
	Body        string `json:"body"`
}

// sendOptions are some settings of sending task.
type SendOptions struct {

	// Sync means the send task is synchronous, default is asynchronous.
	Sync bool `json:"sync"`
}

// NewEmptySendTask returns an empty SendTask holder.
func NewSendTaskWithDefaultOptions() *SendTask {
	return &SendTask{
		Options: &SendOptions{
			Sync: false,
		},
	}
}

// Result represents the result of one call.
type Result struct {

	// Data is the result data.
	Data []byte `json:"data"`
}

// TestServerImpl_Stop tests ServerImpl.Stop.
func main() {

	// Connect to the remote server.
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:5779")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send a request to server.
	req := NewSendTaskWithDefaultOptions()
	req.Email = &Email{
		To:          "fishgoddess@qq.com",
		Subject:     "jsonrpc 测试 postar 运行情况",
		ContentType: "text/html; charset=utf-8",
		Body:        "<h1>哈喽！来自 <span style=\"color: #123456;\">postar<span> 的问候！</h1>",
	}
	resp := &Result{}
	err = conn.Call("PostarService.Send", req, resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp.Data))
}
