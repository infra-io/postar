// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/27 22:50:56

package context

import (
	"context"
	"testing"

	"github.com/FishGoddess/logit"
	"github.com/avino-plan/postar/pkg/trace"
)

// go test -v -cover -run=^TestWithLogger$
func TestWithLogger(t *testing.T) {
	logger := logit.NewLogger()
	ctx := WithLogger(context.Background(), logger)
	contextLogger := logit.FromContext(ctx)
	if contextLogger != logger {
		t.Errorf("contextLogger %+v != logger %+v", contextLogger, logger)
	}
}

// go test -v -cover -run=^TestWithTraceID$
func TestWithTraceID(t *testing.T) {
	ctx := WithTraceID(context.Background())
	traceId := trace.FromContext(ctx)
	if traceId == "" {
		t.Error("traceId == ''")
	}
}
