// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: postaradmin/v1/space_service.proto

package postaradminv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SpaceService_CreateSpace_FullMethodName = "/postaradmin.v1.SpaceService/CreateSpace"
	SpaceService_UpdateSpace_FullMethodName = "/postaradmin.v1.SpaceService/UpdateSpace"
	SpaceService_GetSpace_FullMethodName    = "/postaradmin.v1.SpaceService/GetSpace"
	SpaceService_ListSpaces_FullMethodName  = "/postaradmin.v1.SpaceService/ListSpaces"
)

// SpaceServiceClient is the client API for SpaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpaceServiceClient interface {
	CreateSpace(ctx context.Context, in *CreateSpaceRequest, opts ...grpc.CallOption) (*CreateSpaceResponse, error)
	UpdateSpace(ctx context.Context, in *UpdateSpaceRequest, opts ...grpc.CallOption) (*UpdateSpaceResponse, error)
	GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*GetSpaceResponse, error)
	ListSpaces(ctx context.Context, in *ListSpacesRequest, opts ...grpc.CallOption) (*ListSpacesResponse, error)
}

type spaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceServiceClient(cc grpc.ClientConnInterface) SpaceServiceClient {
	return &spaceServiceClient{cc}
}

func (c *spaceServiceClient) CreateSpace(ctx context.Context, in *CreateSpaceRequest, opts ...grpc.CallOption) (*CreateSpaceResponse, error) {
	out := new(CreateSpaceResponse)
	err := c.cc.Invoke(ctx, SpaceService_CreateSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) UpdateSpace(ctx context.Context, in *UpdateSpaceRequest, opts ...grpc.CallOption) (*UpdateSpaceResponse, error) {
	out := new(UpdateSpaceResponse)
	err := c.cc.Invoke(ctx, SpaceService_UpdateSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*GetSpaceResponse, error) {
	out := new(GetSpaceResponse)
	err := c.cc.Invoke(ctx, SpaceService_GetSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) ListSpaces(ctx context.Context, in *ListSpacesRequest, opts ...grpc.CallOption) (*ListSpacesResponse, error) {
	out := new(ListSpacesResponse)
	err := c.cc.Invoke(ctx, SpaceService_ListSpaces_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceServiceServer is the server API for SpaceService service.
// All implementations must embed UnimplementedSpaceServiceServer
// for forward compatibility
type SpaceServiceServer interface {
	CreateSpace(context.Context, *CreateSpaceRequest) (*CreateSpaceResponse, error)
	UpdateSpace(context.Context, *UpdateSpaceRequest) (*UpdateSpaceResponse, error)
	GetSpace(context.Context, *GetSpaceRequest) (*GetSpaceResponse, error)
	ListSpaces(context.Context, *ListSpacesRequest) (*ListSpacesResponse, error)
	mustEmbedUnimplementedSpaceServiceServer()
}

// UnimplementedSpaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpaceServiceServer struct {
}

func (UnimplementedSpaceServiceServer) CreateSpace(context.Context, *CreateSpaceRequest) (*CreateSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpace not implemented")
}
func (UnimplementedSpaceServiceServer) UpdateSpace(context.Context, *UpdateSpaceRequest) (*UpdateSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSpace not implemented")
}
func (UnimplementedSpaceServiceServer) GetSpace(context.Context, *GetSpaceRequest) (*GetSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpace not implemented")
}
func (UnimplementedSpaceServiceServer) ListSpaces(context.Context, *ListSpacesRequest) (*ListSpacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpaces not implemented")
}
func (UnimplementedSpaceServiceServer) mustEmbedUnimplementedSpaceServiceServer() {}

// UnsafeSpaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpaceServiceServer will
// result in compilation errors.
type UnsafeSpaceServiceServer interface {
	mustEmbedUnimplementedSpaceServiceServer()
}

func RegisterSpaceServiceServer(s grpc.ServiceRegistrar, srv SpaceServiceServer) {
	s.RegisterService(&SpaceService_ServiceDesc, srv)
}

func _SpaceService_CreateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).CreateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_CreateSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).CreateSpace(ctx, req.(*CreateSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_UpdateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).UpdateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_UpdateSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).UpdateSpace(ctx, req.(*UpdateSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_GetSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).GetSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_GetSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).GetSpace(ctx, req.(*GetSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_ListSpaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSpacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).ListSpaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_ListSpaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).ListSpaces(ctx, req.(*ListSpacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpaceService_ServiceDesc is the grpc.ServiceDesc for SpaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "postaradmin.v1.SpaceService",
	HandlerType: (*SpaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSpace",
			Handler:    _SpaceService_CreateSpace_Handler,
		},
		{
			MethodName: "UpdateSpace",
			Handler:    _SpaceService_UpdateSpace_Handler,
		},
		{
			MethodName: "GetSpace",
			Handler:    _SpaceService_GetSpace_Handler,
		},
		{
			MethodName: "ListSpaces",
			Handler:    _SpaceService_ListSpaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "postaradmin/v1/space_service.proto",
}
