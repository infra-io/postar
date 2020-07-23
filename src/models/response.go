// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/12 23:05:37

package models

import "encoding/json"

// response is the struct represents of the response of request.
type response struct {

	// code is the status code.
	Code int `json:"code"`

	// msg is the message.
	Msg string `json:"msg"`

	// data is the extra thing that returns.
	Data interface{} `json:"data"`
}

// newResponse returns a response holder.
func newResponse(code int, msg string, data interface{}) *response {
	return &response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// ======================================= all responses =======================================

// Optimization point: At first, we want to use singleton pattern to code these,
// however, every responses will have a sync.Once holder which was originally a waste.
// So we just serialize them in bytes and return bytes directly.
var (
	// emailSuccessfullySentResponse means email successfully sent.
	emailSuccessfullySentResponse, _ = json.Marshal(newResponse(0, "Email successfully sent!", nil))

	// serverIsClosingResponse means server is closing.
	serverIsClosingResponse, _ = json.Marshal(newResponse(0, "Server is closing! Please wait a few moments.", nil))

	// wrongRequestBodyResponse means you should check your request body.
	wrongRequestBodyResponse, _ = json.Marshal(newResponse(100, "Please check your request body! It should be a Json string!", nil))

	// failedToSendEmailResponse means failed to send this email.
	failedToSendEmailResponse, _ = json.Marshal(newResponse(200, "Failed to send this email, maybe the email has something wrong?", nil))

	// failedToCloseServerResponse means failed to close server.
	failedToCloseServerResponse, _ = json.Marshal(newResponse(300, "Failed to close server! Try to kill it?", nil))
)

// EmailSuccessfullySentResponse returns emailSuccessfullySentResponse.
func EmailSuccessfullySentResponse() []byte {
	return emailSuccessfullySentResponse
}

// ServerIsClosingResponse returns serverIsClosingResponse.
func ServerIsClosingResponse() []byte {
	return serverIsClosingResponse
}

// WrongRequestBodyResponse returns wrongRequestBodyResponse.
func WrongRequestBodyResponse() []byte {
	return wrongRequestBodyResponse
}

// FailedToSendEmailResponse returns failedToSendEmailResponse.
func FailedToSendEmailResponse() []byte {
	return failedToSendEmailResponse
}

// FailedToCloseServerResponse returns failedToCloseServerResponse.
func FailedToCloseServerResponse() []byte {
	return failedToCloseServerResponse
}
