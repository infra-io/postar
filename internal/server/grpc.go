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
	"github.com/avinoplan/postar/api"
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/biz"
	"github.com/avinoplan/postar/pkg/errors"
	"github.com/avinoplan/postar/pkg/trace"
	"google.golang.org/grpc"
	"net"
	"time"
)

// GRPCServer is a grpc implement of PostardServer.
type GRPCServer struct {
	api.UnimplementedPostarServiceServer
	c       *configs.Config
	smtpBiz *biz.SMTPBiz
	server  *grpc.Server
}

// NewGRPCServer returns a new GRPCServer.
func NewGRPCServer(c *configs.Config, smtpBiz *biz.SMTPBiz) Server {
	gs := &GRPCServer{
		c:       c,
		smtpBiz: smtpBiz,
		server:  grpc.NewServer(),
	}

	api.RegisterPostarServiceServer(gs.server, gs)
	return gs
}

// SendEmail sends emails.
func (gs *GRPCServer) SendEmail(ctx context.Context, request *api.SendEmailRequest) (*api.SendEmailResponse, error) {
	traceID := trace.NewTraceID()
	ctx = trace.NewContext(ctx, traceID)

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
	defer listener.Close()
	return gs.server.Serve(listener)
}

// Stop stops GRPCServer gracefully.
func (gs *GRPCServer) Stop() error {
	stopCh := make(chan struct{}, 1)

	go func() {
		gs.server.GracefulStop()
		time.Sleep(time.Minute)
		stopCh <- struct{}{}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), gs.c.ServerStopTimeout())
	defer cancel()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-stopCh:
		return nil
	}
}
