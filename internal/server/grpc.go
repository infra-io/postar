// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"net"

	"github.com/FishGoddess/errors"
	baseapi "github.com/avino-plan/api/go-out/base"
	postarapi "github.com/avino-plan/api/go-out/postar"
	"github.com/avino-plan/postar/configs"
	"github.com/avino-plan/postar/internal/biz"
	"github.com/avino-plan/postar/pkg/trace"
	"google.golang.org/grpc"
)

// GRPCServer is a grpc implement of PostardServer.
type GRPCServer struct {
	postarapi.UnimplementedPostarServiceServer
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

	postarapi.RegisterPostarServiceServer(gs.server, gs)
	return gs
}

// SendEmail sends emails.
func (gs *GRPCServer) SendEmail(ctx context.Context, request *postarapi.SendEmailRequest) (*postarapi.SendEmailResponse, error) {
	traceID := trace.NewTraceID()
	ctx = trace.NewContext(ctx, traceID)

	err := gs.smtpBiz.SendEmail(ctx, toModelEmail(request.Email), toModelSendEmailOptions(gs.c, request.Options))
	if errors.IsBadRequest(err) {
		return &postarapi.SendEmailResponse{
			Code:    baseapi.ServerCode_BAD_REQUEST,
			Msg:     err.Error(),
			TraceId: traceID,
		}, nil
	}

	if errors.IsTimeout(err) {
		return &postarapi.SendEmailResponse{
			Code:    baseapi.ServerCode_TIMEOUT,
			Msg:     "send email timeout",
			TraceId: traceID,
		}, nil
	}

	if err != nil {
		return &postarapi.SendEmailResponse{
			Code:    baseapi.ServerCode_SEND_EMAIL_FAILED,
			Msg:     "send email failed",
			TraceId: traceID,
		}, nil
	}

	return &postarapi.SendEmailResponse{
		Code:    baseapi.ServerCode_OK,
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
