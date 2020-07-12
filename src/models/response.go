// Copyright 2020 Ye Zi Jie. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/12 23:05:37

package models

import "encoding/json"

// 响应结构体
type response struct {

	// code 是状态码
	code int

	// msg 是信息
	msg string

	// data 是携带的数据
	data interface{}
}

// newResponse 返回一个完整的响应结构体对象
func newResponse(code int, msg string, data interface{}) *response {
	return &response{
		code: code,
		msg:  msg,
		data: data,
	}
}

// ======================================= all responses =======================================

// 优化点：把这些常用的响应结果做成单例（sync.Once）甚至是提前序列化成 Json 的字节数组
// 一开始考虑使用 sync.Once 做单例的方式，但是想想，每一个对象都要使用单独的 once 对象，
// 这还是浪费了一些内存，所以最后选择一步到位，将常用的响应结果序列化成 Json 字节数组！
var (
	// 邮件成功发送的响应状态码和信息
	emailSuccessfullySentResponse, _ = json.Marshal(newResponse(0, "Email successfully sent!", nil))

	// 请求体错误的响应结果
	wrongRequestBodyResponse, _ = json.Marshal(newResponse(100, "Please check your request body! It should be a Json string!", nil))

	// 邮件发送失败的响应状态码和信息
	failedToSendEmailResponse, _ = json.Marshal(newResponse(200, "Failed to send this email, maybe the email has something wrong!", nil))
)

// EmailSuccessfullySentResponse 返回一个邮件成功发送的响应结果
func EmailSuccessfullySentResponse() []byte {
	return emailSuccessfullySentResponse
}

// WrongRequestBodyResponse 返回一个请求体错误的响应结果
func WrongRequestBodyResponse() []byte {
	return wrongRequestBodyResponse
}

// FailedToSendEmailResponse 返回一个邮件发送失败的响应结果
func FailedToSendEmailResponse() []byte {
	return failedToSendEmailResponse
}
