// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/16 02:19:43

package model

var (
	SuccessfulResponse              = &Response{Code: 0, Msg: "ok"}
	RequestMethodNotAllowedResponse = &Response{Code: -1001, Msg: "request method not allowed"}
)

var (
	GetSendRequestFailedResponse = &Response{Code: -10001, Msg: "get send request failed"}
	SendEmailTimeoutResponse     = &Response{Code: -10002, Msg: "send email timeout"}
	SendEmailFailedResponse      = &Response{Code: -10003, Msg: "send email failed"}
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewSuccessfulResponse(data interface{}) *Response {
	return &Response{
		Code: SuccessfulResponse.Code,
		Msg:  SuccessfulResponse.Msg,
		Data: data,
	}
}
