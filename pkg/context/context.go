// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/27 22:40:58

package context

import (
	"context"

	"github.com/FishGoddess/logit"
	"github.com/avino-plan/postar/pkg/trace"
)

// WithLogger wraps ctx with logger.
func WithLogger(ctx context.Context, logger *logit.Logger) context.Context {
	return logit.NewContext(ctx, logger)
}

// WithTraceID wraps ctx with traceID.
func WithTraceID(ctx context.Context, traceID string) context.Context {
	if traceID == "" {
		traceID = trace.NewTraceID()
	}
	return trace.NewContext(ctx, traceID)
}
