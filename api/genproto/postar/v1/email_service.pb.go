// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: postar/v1/email_service.proto

package postarv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *Email `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postar_v1_email_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postar_v1_email_service_proto_msgTypes[0]
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
	return file_postar_v1_email_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailRequest) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

type SendEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendEmailResponse) Reset() {
	*x = SendEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postar_v1_email_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailResponse) ProtoMessage() {}

func (x *SendEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postar_v1_email_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailResponse.ProtoReflect.Descriptor instead.
func (*SendEmailResponse) Descriptor() ([]byte, []int) {
	return file_postar_v1_email_service_proto_rawDescGZIP(), []int{1}
}

var File_postar_v1_email_service_proto protoreflect.FileDescriptor

var file_postar_v1_email_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3a, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x13, 0x0a, 0x11, 0x53,
	0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x7d, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6d, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x42,
	0xa3, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x11, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x61,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x61,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postar_v1_email_service_proto_rawDescOnce sync.Once
	file_postar_v1_email_service_proto_rawDescData = file_postar_v1_email_service_proto_rawDesc
)

func file_postar_v1_email_service_proto_rawDescGZIP() []byte {
	file_postar_v1_email_service_proto_rawDescOnce.Do(func() {
		file_postar_v1_email_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_postar_v1_email_service_proto_rawDescData)
	})
	return file_postar_v1_email_service_proto_rawDescData
}

var file_postar_v1_email_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_postar_v1_email_service_proto_goTypes = []interface{}{
	(*SendEmailRequest)(nil),  // 0: postar.v1.SendEmailRequest
	(*SendEmailResponse)(nil), // 1: postar.v1.SendEmailResponse
	(*Email)(nil),             // 2: postar.v1.Email
}
var file_postar_v1_email_service_proto_depIdxs = []int32{
	2, // 0: postar.v1.SendEmailRequest.email:type_name -> postar.v1.Email
	0, // 1: postar.v1.EmailService.SendEmail:input_type -> postar.v1.SendEmailRequest
	1, // 2: postar.v1.EmailService.SendEmail:output_type -> postar.v1.SendEmailResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_postar_v1_email_service_proto_init() }
func file_postar_v1_email_service_proto_init() {
	if File_postar_v1_email_service_proto != nil {
		return
	}
	file_postar_v1_email_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_postar_v1_email_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_postar_v1_email_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailResponse); i {
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
			RawDescriptor: file_postar_v1_email_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_postar_v1_email_service_proto_goTypes,
		DependencyIndexes: file_postar_v1_email_service_proto_depIdxs,
		MessageInfos:      file_postar_v1_email_service_proto_msgTypes,
	}.Build()
	File_postar_v1_email_service_proto = out.File
	file_postar_v1_email_service_proto_rawDesc = nil
	file_postar_v1_email_service_proto_goTypes = nil
	file_postar_v1_email_service_proto_depIdxs = nil
}
