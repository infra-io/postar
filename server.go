// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/12 00:26:28

package main

import (
	"encoding/json"
	"net/http"
)

type Server interface {
	Serve(address string) error
	Close() error
}

type SendRequest struct {
	Email       *Email       `json:"email"`
	SendOptions *SendOptions `json:"sendOptions"`
}

func newSendRequest() *SendRequest {
	sendOptions := DefaultSendOptions()
	return &SendRequest{
		Email:       nil,
		SendOptions: &sendOptions,
	}
}

type HttpServer struct {
	sender Sender
}

func NewHttpServer(sender Sender) *HttpServer {
	return &HttpServer{
		sender: sender,
	}
}

func getSendRequestFrom(request *http.Request) (*SendRequest, error) {
	defer request.Body.Close()
	sendRequest := newSendRequest()
	return sendRequest, json.NewDecoder(request.Body).Decode(sendRequest)
}

func (hs *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	sendRequest, err := getSendRequestFrom(request)
	if err != nil {
		Logger().Error("get email from request failed").Error("err", err).End()
		writer.Write([]byte(err.Error()))
		return
	}

	err = hs.sender.SendEmail(sendRequest.Email, sendRequest.SendOptions)
	if err != nil {
		Logger().Error("send email failed").Error("err", err).Any("sendRequest.Email", sendRequest.Email).Any("sendRequest.SendOptions", sendRequest.SendOptions).End()
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Write([]byte("ok"))
}

func (hs *HttpServer) Serve(address string) error {
	return http.ListenAndServe(address, hs)
}

func (hs *HttpServer) Close() error {
	return nil
}
