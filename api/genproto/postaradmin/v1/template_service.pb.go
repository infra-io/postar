// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: postaradmin/v1/template_service.proto

package postaradminv1

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

type CreateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *CreateTemplateRequest) Reset() {
	*x = CreateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateRequest) ProtoMessage() {}

func (x *CreateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTemplateRequest) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

type CreateTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *CreateTemplateResponse) Reset() {
	*x = CreateTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateResponse) ProtoMessage() {}

func (x *CreateTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateResponse.ProtoReflect.Descriptor instead.
func (*CreateTemplateResponse) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTemplateResponse) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

type UpdateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *UpdateTemplateRequest) Reset() {
	*x = UpdateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateRequest) ProtoMessage() {}

func (x *UpdateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateRequest.ProtoReflect.Descriptor instead.
func (*UpdateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTemplateRequest) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

type UpdateTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *UpdateTemplateResponse) Reset() {
	*x = UpdateTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateResponse) ProtoMessage() {}

func (x *UpdateTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateResponse.ProtoReflect.Descriptor instead.
func (*UpdateTemplateResponse) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTemplateResponse) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

type GetTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId int64 `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
}

func (x *GetTemplateRequest) Reset() {
	*x = GetTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateRequest) ProtoMessage() {}

func (x *GetTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateRequest.ProtoReflect.Descriptor instead.
func (*GetTemplateRequest) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTemplateRequest) GetTemplateId() int64 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

type GetTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *GetTemplateResponse) Reset() {
	*x = GetTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateResponse) ProtoMessage() {}

func (x *GetTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateResponse.ProtoReflect.Descriptor instead.
func (*GetTemplateResponse) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetTemplateResponse) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

type ListTemplatesFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId     int32         `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	TemplateId    int64         `protobuf:"varint,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	TemplateName  string        `protobuf:"bytes,3,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	TemplateState TemplateState `protobuf:"varint,4,opt,name=template_state,json=templateState,proto3,enum=postaradmin.v1.TemplateState" json:"template_state,omitempty"`
	EmailSubject  string        `protobuf:"bytes,5,opt,name=email_subject,json=emailSubject,proto3" json:"email_subject,omitempty"`
}

func (x *ListTemplatesFilter) Reset() {
	*x = ListTemplatesFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplatesFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplatesFilter) ProtoMessage() {}

func (x *ListTemplatesFilter) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplatesFilter.ProtoReflect.Descriptor instead.
func (*ListTemplatesFilter) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListTemplatesFilter) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *ListTemplatesFilter) GetTemplateId() int64 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *ListTemplatesFilter) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *ListTemplatesFilter) GetTemplateState() TemplateState {
	if x != nil {
		return x.TemplateState
	}
	return TemplateState_TEMPLATE_STATE_UNSPECIFIED
}

func (x *ListTemplatesFilter) GetEmailSubject() string {
	if x != nil {
		return x.EmailSubject
	}
	return ""
}

type ListTemplatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter    *ListTemplatesFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	PageToken string               `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize  int32                `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListTemplatesRequest) Reset() {
	*x = ListTemplatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplatesRequest) ProtoMessage() {}

func (x *ListTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplatesRequest.ProtoReflect.Descriptor instead.
func (*ListTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListTemplatesRequest) GetFilter() *ListTemplatesFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListTemplatesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListTemplatesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Templates     []*Template `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
	NextPageToken string      `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTemplatesResponse) Reset() {
	*x = ListTemplatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplatesResponse) ProtoMessage() {}

func (x *ListTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplatesResponse.ProtoReflect.Descriptor instead.
func (*ListTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListTemplatesResponse) GetTemplates() []*Template {
	if x != nil {
		return x.Templates
	}
	return nil
}

func (x *ListTemplatesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type DeleteTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId int64 `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
}

func (x *DeleteTemplateRequest) Reset() {
	*x = DeleteTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateRequest) ProtoMessage() {}

func (x *DeleteTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateRequest.ProtoReflect.Descriptor instead.
func (*DeleteTemplateRequest) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteTemplateRequest) GetTemplateId() int64 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

type DeleteTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTemplateResponse) Reset() {
	*x = DeleteTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postaradmin_v1_template_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateResponse) ProtoMessage() {}

func (x *DeleteTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postaradmin_v1_template_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateResponse.ProtoReflect.Descriptor instead.
func (*DeleteTemplateResponse) Descriptor() ([]byte, []int) {
	return file_postaradmin_v1_template_service_proto_rawDescGZIP(), []int{10}
}

var File_postaradmin_v1_template_service_proto protoreflect.FileDescriptor

var file_postaradmin_v1_template_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a,
	0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x22, 0x4e, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x22, 0x4d, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x08,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x22, 0x4e, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x8f,
	0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x77, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd9, 0x05,
	0x0a, 0x0f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x8a, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2d, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x8a,
	0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x25, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61,
	0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x32, 0x1e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x8c, 0x01, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12, 0x2c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x84, 0x01, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2d, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x95, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x2a, 0x2c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0xc9, 0x01, 0x0a, 0x12, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x42, 0x14, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postaradmin_v1_template_service_proto_rawDescOnce sync.Once
	file_postaradmin_v1_template_service_proto_rawDescData = file_postaradmin_v1_template_service_proto_rawDesc
)

func file_postaradmin_v1_template_service_proto_rawDescGZIP() []byte {
	file_postaradmin_v1_template_service_proto_rawDescOnce.Do(func() {
		file_postaradmin_v1_template_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_postaradmin_v1_template_service_proto_rawDescData)
	})
	return file_postaradmin_v1_template_service_proto_rawDescData
}

var file_postaradmin_v1_template_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_postaradmin_v1_template_service_proto_goTypes = []interface{}{
	(*CreateTemplateRequest)(nil),  // 0: postaradmin.v1.CreateTemplateRequest
	(*CreateTemplateResponse)(nil), // 1: postaradmin.v1.CreateTemplateResponse
	(*UpdateTemplateRequest)(nil),  // 2: postaradmin.v1.UpdateTemplateRequest
	(*UpdateTemplateResponse)(nil), // 3: postaradmin.v1.UpdateTemplateResponse
	(*GetTemplateRequest)(nil),     // 4: postaradmin.v1.GetTemplateRequest
	(*GetTemplateResponse)(nil),    // 5: postaradmin.v1.GetTemplateResponse
	(*ListTemplatesFilter)(nil),    // 6: postaradmin.v1.ListTemplatesFilter
	(*ListTemplatesRequest)(nil),   // 7: postaradmin.v1.ListTemplatesRequest
	(*ListTemplatesResponse)(nil),  // 8: postaradmin.v1.ListTemplatesResponse
	(*DeleteTemplateRequest)(nil),  // 9: postaradmin.v1.DeleteTemplateRequest
	(*DeleteTemplateResponse)(nil), // 10: postaradmin.v1.DeleteTemplateResponse
	(*Template)(nil),               // 11: postaradmin.v1.Template
	(TemplateState)(0),             // 12: postaradmin.v1.TemplateState
}
var file_postaradmin_v1_template_service_proto_depIdxs = []int32{
	11, // 0: postaradmin.v1.CreateTemplateRequest.template:type_name -> postaradmin.v1.Template
	11, // 1: postaradmin.v1.CreateTemplateResponse.template:type_name -> postaradmin.v1.Template
	11, // 2: postaradmin.v1.UpdateTemplateRequest.template:type_name -> postaradmin.v1.Template
	11, // 3: postaradmin.v1.UpdateTemplateResponse.template:type_name -> postaradmin.v1.Template
	11, // 4: postaradmin.v1.GetTemplateResponse.template:type_name -> postaradmin.v1.Template
	12, // 5: postaradmin.v1.ListTemplatesFilter.template_state:type_name -> postaradmin.v1.TemplateState
	6,  // 6: postaradmin.v1.ListTemplatesRequest.filter:type_name -> postaradmin.v1.ListTemplatesFilter
	11, // 7: postaradmin.v1.ListTemplatesResponse.templates:type_name -> postaradmin.v1.Template
	0,  // 8: postaradmin.v1.TemplateService.CreateTemplate:input_type -> postaradmin.v1.CreateTemplateRequest
	2,  // 9: postaradmin.v1.TemplateService.UpdateTemplate:input_type -> postaradmin.v1.UpdateTemplateRequest
	4,  // 10: postaradmin.v1.TemplateService.GetTemplate:input_type -> postaradmin.v1.GetTemplateRequest
	7,  // 11: postaradmin.v1.TemplateService.ListTemplates:input_type -> postaradmin.v1.ListTemplatesRequest
	9,  // 12: postaradmin.v1.TemplateService.DeleteTemplate:input_type -> postaradmin.v1.DeleteTemplateRequest
	1,  // 13: postaradmin.v1.TemplateService.CreateTemplate:output_type -> postaradmin.v1.CreateTemplateResponse
	3,  // 14: postaradmin.v1.TemplateService.UpdateTemplate:output_type -> postaradmin.v1.UpdateTemplateResponse
	5,  // 15: postaradmin.v1.TemplateService.GetTemplate:output_type -> postaradmin.v1.GetTemplateResponse
	8,  // 16: postaradmin.v1.TemplateService.ListTemplates:output_type -> postaradmin.v1.ListTemplatesResponse
	10, // 17: postaradmin.v1.TemplateService.DeleteTemplate:output_type -> postaradmin.v1.DeleteTemplateResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_postaradmin_v1_template_service_proto_init() }
func file_postaradmin_v1_template_service_proto_init() {
	if File_postaradmin_v1_template_service_proto != nil {
		return
	}
	file_postaradmin_v1_template_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_postaradmin_v1_template_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateRequest); i {
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
		file_postaradmin_v1_template_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateResponse); i {
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
		file_postaradmin_v1_template_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTemplateRequest); i {
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
		file_postaradmin_v1_template_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTemplateResponse); i {
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
		file_postaradmin_v1_template_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateRequest); i {
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
		file_postaradmin_v1_template_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateResponse); i {
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
		file_postaradmin_v1_template_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplatesFilter); i {
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
		file_postaradmin_v1_template_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplatesRequest); i {
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
		file_postaradmin_v1_template_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplatesResponse); i {
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
		file_postaradmin_v1_template_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTemplateRequest); i {
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
		file_postaradmin_v1_template_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTemplateResponse); i {
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
			RawDescriptor: file_postaradmin_v1_template_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_postaradmin_v1_template_service_proto_goTypes,
		DependencyIndexes: file_postaradmin_v1_template_service_proto_depIdxs,
		MessageInfos:      file_postaradmin_v1_template_service_proto_msgTypes,
	}.Build()
	File_postaradmin_v1_template_service_proto = out.File
	file_postaradmin_v1_template_service_proto_rawDesc = nil
	file_postaradmin_v1_template_service_proto_goTypes = nil
	file_postaradmin_v1_template_service_proto_depIdxs = nil
}
