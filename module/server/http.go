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

	"github.com/avino-plan/postar/base/model"
	"github.com/avino-plan/postar/module"
	"github.com/avino-plan/postar/module/sender"
)

type HttpServer struct {
	address string
	sender  sender.Sender
}

func newHttpServer() Server {
	return &HttpServer{}
}

func (hs *HttpServer) Configure(config *module.Config) error {
	hs.address = config.Server.Address
	return nil
}

func (hs *HttpServer) ConfigureSender(sender sender.Sender) {
	hs.sender = sender
}

func (hs *HttpServer) writeResponse(writer http.ResponseWriter, statusCode int, response *model.Response) {

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	marshaled, err := json.Marshal(response)
	if err != nil {
		module.Logger().Error("marshal response to Json failed").Error("err", err).End()
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"code": -1, "msg": "marshal response to Json failed", "data": null}`))
		return
	}

	writer.WriteHeader(statusCode)
	writer.Write(marshaled)
}

func (hs *HttpServer) rootHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		module.Logger().Error("request method not allowed").String("request.Method", request.Method).End()
		hs.writeResponse(writer, http.StatusMethodNotAllowed, model.RequestMethodNotAllowedResponse)
		return
	}
	hs.writeResponse(writer, http.StatusOK, model.NewSuccessfulResponse(map[string]interface{}{
		"service":      "postar",
		"introduction": "You know, for sending emails!",
		"version":      module.Version,
	}))
}

func (hs *HttpServer) getSendRequestFrom(request *http.Request) (*SendRequest, error) {
	defer request.Body.Close()
	sendRequest := newSendRequest()
	return sendRequest, json.NewDecoder(request.Body).Decode(sendRequest)
}

func (hs *HttpServer) sendEmailHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "PUT" {
		module.Logger().Error("request method not allowed").String("request.Method", request.Method).End()
		hs.writeResponse(writer, http.StatusMethodNotAllowed, model.RequestMethodNotAllowedResponse)
		return
	}

	sendRequest, err := hs.getSendRequestFrom(request)
	if err != nil {
		module.Logger().Error("get send request failed").Error("err", err).End()
		hs.writeResponse(writer, http.StatusBadRequest, model.GetSendRequestFailedResponse)
		return
	}

	err = hs.sender.SendEmail(sendRequest.Email, sendRequest.Options)
	if sender.IsTimeout(err) {
		module.Logger().Error("send timeout").Error("err", err).Any("email", sendRequest.Email).Any("options", sendRequest.Options).End()
		hs.writeResponse(writer, http.StatusInternalServerError, model.SendEmailTimeoutResponse)
		return
	}

	if err != nil {
		module.Logger().Error("send email failed").Error("err", err).Any("email", sendRequest.Email).Any("options", sendRequest.Options).End()
		hs.writeResponse(writer, http.StatusInternalServerError, model.SendEmailFailedResponse)
		return
	}

	hs.writeResponse(writer, http.StatusOK, model.SuccessfulResponse)
}

func (hs *HttpServer) Serve() error {
	handlers := http.NewServeMux()
	handlers.HandleFunc("/", hs.rootHandler)
	handlers.HandleFunc("/send", hs.sendEmailHandler)
	go func() {
		err := http.ListenAndServe(hs.address, handlers)
		if err != nil {
			panic(err)
		}
	}()
	return nil
}

func (hs *HttpServer) Close() error {
	return nil
}
