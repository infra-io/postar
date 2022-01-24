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
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/biz"
	"github.com/avinoplan/postar/pkg/errors"
	"github.com/avinoplan/postar/pkg/trace"
	"google.golang.org/grpc"
)

// GRPCServer is a grpc implement of PostardServer.
type GRPCServer struct {
	api.UnimplementedPostarServiceServer
	server  *grpc.Server
	c       *configs.Config
	logger  *logit.Logger
	smtpBiz *biz.SMTPBiz
}

// NewGRPCServer returns a new GRPCServer.
func NewGRPCServer(c *configs.Config, logger *logit.Logger, smtpBiz *biz.SMTPBiz) Server {
	return &GRPCServer{
		c:       c,
		logger:  logger,
		smtpBiz: smtpBiz,
	}
}

// SendEmail sends emails.
func (gs *GRPCServer) SendEmail(ctx context.Context, request *api.SendEmailRequest) (*api.SendEmailResponse, error) {
	traceID := trace.NewTraceID()
	ctx = trace.NewContext(ctx, traceID)
	ctx = logit.NewContext(ctx, gs.logger)

	err := gs.smtpBiz.SendEmail(ctx, toModelEmail(request.Email), toModelSendEmailOptions(gs.c, request.Options))
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

// Start starts GRPCServer.
func (gs *GRPCServer) Start() error {
	listener, err := net.Listen("tcp", gs.c.ServerAddress())
	if err != nil {
		return err
	}

	gs.server = grpc.NewServer()
	api.RegisterPostarServiceServer(gs.server, gs)
	return gs.server.Serve(listener)
}

// Stop stops GRPCServer gracefully.
func (gs *GRPCServer) Stop() error {
	gs.server.GracefulStop()
	return nil
}
