// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 01:42:33

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: postard.proto

package postard

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ResponseCodes is all codes of response.
type ResponseCodes int32

const (
	ResponseCodes_OK                  ResponseCodes = 0
	ResponseCodes_InternalServerError ResponseCodes = 50000
	ResponseCodes_TimeoutError        ResponseCodes = 50001
)

// Enum value maps for ResponseCodes.
var (
	ResponseCodes_name = map[int32]string{
		0:     "OK",
		50000: "InternalServerError",
		50001: "TimeoutError",
	}
	ResponseCodes_value = map[string]int32{
		"OK":                  0,
		"InternalServerError": 50000,
		"TimeoutError":        50001,
	}
)

func (x ResponseCodes) Enum() *ResponseCodes {
	p := new(ResponseCodes)
	*p = x
	return p
}

func (x ResponseCodes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseCodes) Descriptor() protoreflect.EnumDescriptor {
	return file_postard_proto_enumTypes[0].Descriptor()
}

func (ResponseCodes) Type() protoreflect.EnumType {
	return &file_postard_proto_enumTypes[0]
}

func (x ResponseCodes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseCodes.Descriptor instead.
func (ResponseCodes) EnumDescriptor() ([]byte, []int) {
	return file_postard_proto_rawDescGZIP(), []int{0}
}

// PostardResponse is the response of Postard.
type PostardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ResponseCodes `protobuf:"varint,1,opt,name=code,proto3,enum=github.com.avinoplan.postar.api.postard.ResponseCodes" json:"code,omitempty"` // 0 is ok.
	Msg     string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`                                                               // For messaging.
	TraceId string        `protobuf:"bytes,3,opt,name=traceId,proto3" json:"traceId,omitempty"`                                                       // For tracing.
	Data    *anypb.Any    `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`                                                             // Any data.
}

func (x *PostardResponse) Reset() {
	*x = PostardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostardResponse) ProtoMessage() {}

func (x *PostardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostardResponse.ProtoReflect.Descriptor instead.
func (*PostardResponse) Descriptor() ([]byte, []int) {
	return file_postard_proto_rawDescGZIP(), []int{0}
}

func (x *PostardResponse) GetCode() ResponseCodes {
	if x != nil {
		return x.Code
	}
	return ResponseCodes_OK
}

func (x *PostardResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PostardResponse) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *PostardResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

// Email wraps all information of using smtp service.
type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To       []string `protobuf:"bytes,1,rep,name=to,proto3" json:"to,omitempty"`             // The receivers of one email.
	Subject  string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`   // The subject of one email.
	BodyType string   `protobuf:"bytes,3,opt,name=bodyType,proto3" json:"bodyType,omitempty"` // The content type of body.
	Body     string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`         // The body of one email.
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_postard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_postard_proto_rawDescGZIP(), []int{1}
}

func (x *Email) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Email) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Email) GetBodyType() string {
	if x != nil {
		return x.BodyType
	}
	return ""
}

func (x *Email) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

// SendEmailOptions is the options of sending emails.
type SendEmailOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Async   bool  `protobuf:"varint,1,opt,name=async,proto3" json:"async,omitempty"`     // If need sending emails asynchronously.
	Timeout int64 `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"` // Sending timeout.
}

func (x *SendEmailOptions) Reset() {
	*x = SendEmailOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailOptions) ProtoMessage() {}

func (x *SendEmailOptions) ProtoReflect() protoreflect.Message {
	mi := &file_postard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailOptions.ProtoReflect.Descriptor instead.
func (*SendEmailOptions) Descriptor() ([]byte, []int) {
	return file_postard_proto_rawDescGZIP(), []int{2}
}

func (x *SendEmailOptions) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *SendEmailOptions) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

// SendEmailRequest is the request of SendEmail.
type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   *Email            `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`     // Sending email.
	Options *SendEmailOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"` // Sending options.
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_postard_proto_rawDescGZIP(), []int{3}
}

func (x *SendEmailRequest) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *SendEmailRequest) GetOptions() *SendEmailOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

var File_postard_proto protoreflect.FileDescriptor

var file_postard_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x6e,
	0x6f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x6e, 0x6f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x42, 0x0a, 0x10,
	0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x22, 0xad, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x76, 0x69, 0x6e, 0x6f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x53, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x6e, 0x6f, 0x70,
	0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2a, 0x46, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x13, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0xd0, 0x86, 0x03, 0x12, 0x12, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0xd1, 0x86, 0x03, 0x32, 0x8c, 0x01, 0x0a, 0x07, 0x50, 0x6f, 0x73,
	0x74, 0x61, 0x72, 0x64, 0x12, 0x80, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x39, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x76, 0x69, 0x6e, 0x6f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x6e, 0x6f,
	0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x69, 0x6e, 0x6f, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postard_proto_rawDescOnce sync.Once
	file_postard_proto_rawDescData = file_postard_proto_rawDesc
)

func file_postard_proto_rawDescGZIP() []byte {
	file_postard_proto_rawDescOnce.Do(func() {
		file_postard_proto_rawDescData = protoimpl.X.CompressGZIP(file_postard_proto_rawDescData)
	})
	return file_postard_proto_rawDescData
}

var file_postard_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_postard_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_postard_proto_goTypes = []interface{}{
	(ResponseCodes)(0),       // 0: github.com.avinoplan.postar.api.postard.ResponseCodes
	(*PostardResponse)(nil),  // 1: github.com.avinoplan.postar.api.postard.PostardResponse
	(*Email)(nil),            // 2: github.com.avinoplan.postar.api.postard.Email
	(*SendEmailOptions)(nil), // 3: github.com.avinoplan.postar.api.postard.SendEmailOptions
	(*SendEmailRequest)(nil), // 4: github.com.avinoplan.postar.api.postard.SendEmailRequest
	(*anypb.Any)(nil),        // 5: google.protobuf.Any
}
var file_postard_proto_depIdxs = []int32{
	0, // 0: github.com.avinoplan.postar.api.postard.PostardResponse.code:type_name -> github.com.avinoplan.postar.api.postard.ResponseCodes
	5, // 1: github.com.avinoplan.postar.api.postard.PostardResponse.data:type_name -> google.protobuf.Any
	2, // 2: github.com.avinoplan.postar.api.postard.SendEmailRequest.email:type_name -> github.com.avinoplan.postar.api.postard.Email
	3, // 3: github.com.avinoplan.postar.api.postard.SendEmailRequest.options:type_name -> github.com.avinoplan.postar.api.postard.SendEmailOptions
	4, // 4: github.com.avinoplan.postar.api.postard.Postard.SendEmail:input_type -> github.com.avinoplan.postar.api.postard.SendEmailRequest
	1, // 5: github.com.avinoplan.postar.api.postard.Postard.SendEmail:output_type -> github.com.avinoplan.postar.api.postard.PostardResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_postard_proto_init() }
func file_postard_proto_init() {
	if File_postard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_postard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostardResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_postard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_postard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_postard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_postard_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_postard_proto_goTypes,
		DependencyIndexes: file_postard_proto_depIdxs,
		EnumInfos:         file_postard_proto_enumTypes,
		MessageInfos:      file_postard_proto_msgTypes,
	}.Build()
	File_postard_proto = out.File
	file_postard_proto_rawDesc = nil
	file_postard_proto_goTypes = nil
	file_postard_proto_depIdxs = nil
}
