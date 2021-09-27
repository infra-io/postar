// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 02:05:02

package server

import (
	"context"
	"net"

	"github.com/avino-plan/postar/api/postard"
	"github.com/avino-plan/postar/internal/pkg/trace"
	"github.com/avino-plan/postar/internal/postard/service"
	"google.golang.org/grpc"
)

// PostardGrpcServer is a grpc implement of PostardServer.
type PostardGrpcServer struct {
	postard.UnimplementedPostardServer
	server         *grpc.Server
	contextService service.ContextService
	smtpService    service.SmtpService
}

// NewPostardGrpcServer returns a new PostardGrpcServer.
func NewPostardGrpcServer(contextService service.ContextService, smtpService service.SmtpService) *PostardGrpcServer {
	return &PostardGrpcServer{
		contextService: contextService,
		smtpService:    smtpService,
	}
}

// SendEmail sends emails.
func (pgs *PostardGrpcServer) SendEmail(ctx context.Context, request *postard.SendEmailRequest) (*postard.PostardResponse, error) {
	ctx = pgs.contextService.WrapContext(ctx)
	traceId := trace.FromContext(ctx)

	err := pgs.smtpService.SendEmail(ctx, nil, nil)
	if service.IsSendTimeout(err) {
		return &postard.PostardResponse{
			Code:    postard.ResponseCodes_TimeoutError,
			Msg:     "send email timeout",
			TraceId: traceId,
		}, nil
	}

	if err != nil {
		return &postard.PostardResponse{
			Code:    postard.ResponseCodes_InternalServerError,
			Msg:     "send email failed",
			TraceId: traceId,
		}, nil
	}

	return &postard.PostardResponse{
		Code:    postard.ResponseCodes_OK,
		TraceId: traceId,
	}, nil
}

// Run runs PostardGrpcServer with listener.
func (pgs *PostardGrpcServer) Run(listener net.Listener) error {
	pgs.server = grpc.NewServer()
	postard.RegisterPostardServer(pgs.server, pgs)
	return pgs.server.Serve(listener)
}

// Shutdown shutdowns PostardGrpcServer gracefully.
func (pgs *PostardGrpcServer) Shutdown() {
	pgs.server.GracefulStop()
}
