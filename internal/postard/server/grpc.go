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

	"github.com/avino-plan/postar/api/postard"
	"github.com/avino-plan/postar/internal/pkg/trace"
	"github.com/avino-plan/postar/internal/postard/service"
)

type PostardServerGrpcImpl struct {
	postard.UnimplementedPostardServer
	service service.SmtpService
}

func NewPostardServerGrpcImpl(service service.SmtpService) *PostardServerGrpcImpl {
	return &PostardServerGrpcImpl{
		service: service,
	}
}

func (psgi *PostardServerGrpcImpl) SendEmail(pCtx context.Context, request *postard.SendEmailRequest) (*postard.SendEmailResponse, error) {
	ctx := trace.WithContext(pCtx)

	err := psgi.service.SendEmail(ctx, nil, nil)
	if service.IsSendTimeout(err) {

	}

	if err != nil {

	}
	return &postard.SendEmailResponse{TraceId: trace.FromContext(ctx)}, nil
}
