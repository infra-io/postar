// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/27 22:50:56

package service

import (
	"context"
	"testing"

	"github.com/FishGoddess/logit"
	"github.com/avino-plan/postar/internal/pkg/trace"
)

// go test -v -cover -run=^TestNewContextService$
func TestNewContextService(t *testing.T) {
	logger := logit.NewLogger()

	service := NewContextService(logger)
	ctx := service.WrapContext(context.Background())

	contextLogger := logit.FromContext(ctx)
	if contextLogger != logger {
		t.Errorf("contextLogger %+v != logger %+v", contextLogger, logger)
	}

	traceId := trace.FromContext(ctx)
	if traceId == "" {
		t.Error("traceId == ''")
	}
}
