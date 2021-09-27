// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/27 22:40:58

package service

import (
	"context"

	"github.com/FishGoddess/logit"
	"github.com/avino-plan/postar/internal/pkg/trace"
)

// contextServiceImpl is the service of context.
type contextServiceImpl struct {
	logger *logit.Logger
}

// NewContextService returns a new ContextService.
func NewContextService(logger *logit.Logger) ContextService {
	return &contextServiceImpl{logger: logger}
}

// WrapContext wraps context with something and returns a new context.
func (csi *contextServiceImpl) WrapContext(ctx context.Context) context.Context {
	ctx = logit.NewContext(ctx, csi.logger)
	return trace.WithContext(ctx)
}
