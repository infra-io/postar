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

	"github.com/FishGoddess/logit"
	"github.com/avinoplan/postar/api"
	"github.com/avinoplan/postar/internal/biz"
	"github.com/avinoplan/postar/pkg/errors"
	"github.com/avinoplan/postar/pkg/trace"
	"google.golang.org/grpc"
)

// GRPCServer is a grpc implement of PostardServer.
type GRPCServer struct {
	api.UnimplementedPostardServer
	server  *grpc.Server
	logger  *logit.Logger
	smtpBiz *biz.SmtpBiz
}

// NewGrpcServer returns a new GRPCServer.
func NewGrpcServer(logger *logit.Logger, smtpBiz *biz.SmtpBiz) *GRPCServer {
	return &GRPCServer{
		logger:  logger,
		smtpBiz: smtpBiz,
	}
}

// SendEmail sends emails.
func (gs *GRPCServer) SendEmail(ctx context.Context, request *api.SendEmailRequest) (*api.SendEmailResponse, error) {
	traceID := trace.NewTraceID()
	ctx = trace.NewContext(ctx, traceID)
	ctx = logit.NewContext(ctx, gs.logger)

	err := gs.smtpBiz.SendEmail(ctx, nil, nil)
	if errors.IsTimeout(err) {
		return &api.SendEmailResponse{
			Code:    api.ServerCode_TIMEOUT,
			Msg:     "send email timeout",
			TraceId: traceID,
		}, nil
	}

	if err != nil {
		return &api.SendEmailResponse{
			Code:    api.ServerCode_SEND_EMAIL_FAILED,
			Msg:     "send email failed",
			TraceId: traceID,
		}, nil
	}

	return &api.SendEmailResponse{
		Code:    api.ServerCode_OK,
		TraceId: traceID,
	}, nil
}

// Run runs GRPCServer with listener.
func (gs *GRPCServer) Run(listener net.Listener) error {
	gs.server = grpc.NewServer()
	api.RegisterPostardServer(gs.server, gs)
	return gs.server.Serve(listener)
}

// Shutdown shutdowns GRPCServer gracefully.
func (gs *GRPCServer) Shutdown() {
	gs.server.GracefulStop()
}
