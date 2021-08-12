// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/12 00:26:28

package server

import (
	"encoding/json"
	"net/http"

	"github.com/avino-plan/postar/module"
	"github.com/avino-plan/postar/module/sender"
)

type SendRequest struct {
	Email       *sender.Email       `json:"email"`
	SendOptions *sender.SendOptions `json:"sendOptions"`
}

func newSendRequest() *SendRequest {
	sendOptions := sender.DefaultSendOptions()
	return &SendRequest{
		Email:       nil,
		SendOptions: &sendOptions,
	}
}

type HttpServer struct {
	address string
	sender  sender.Sender
}

func newHttpServer() Server {
	return &HttpServer{}
}

func getSendRequestFrom(request *http.Request) (*SendRequest, error) {
	defer request.Body.Close()
	sendRequest := newSendRequest()
	return sendRequest, json.NewDecoder(request.Body).Decode(sendRequest)
}

func (hs *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	sendRequest, err := getSendRequestFrom(request)
	if err != nil {
		module.Logger().Error("get email from request failed").Error("err", err).End()
		writer.Write([]byte(err.Error()))
		return
	}

	err = hs.sender.SendEmail(sendRequest.Email, sendRequest.SendOptions)
	if err != nil {
		module.Logger().Error("send email failed").Error("err", err).Any("sendRequest.Email", sendRequest.Email).Any("sendRequest.SendOptions", sendRequest.SendOptions).End()
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Write([]byte("ok"))
}

func (hs *HttpServer) Configure(config *module.Config) error {
	hs.address = config.Server.Address
	return nil
}

func (hs *HttpServer) SetSender(sender sender.Sender) {
	hs.sender = sender
}

func (hs *HttpServer) Serve() error {
	go func() {
		err := http.ListenAndServe(hs.address, hs)
		if err != nil {
			panic(err)
		}
	}()
	return nil
}

func (hs *HttpServer) Close() error {
	return nil
}
