// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 23:07:54

package trace

import (
	"context"
	"testing"
)

// go test -v -cover -run=^TestWithContext$
func TestWithContext(t *testing.T) {
	ctx := WithContext(context.Background())

	value := ctx.Value(traceIdKey)
	if value == nil {
		t.Error("ctx.Value returns nil")
	}

	traceId, ok := value.(string)
	if !ok {
		t.Errorf("value %+v isn't string", value)
	}
	t.Log("traceId:", traceId)
}

// go test -v -cover -run=^TestFromContext$
func TestFromContext(t *testing.T) {
	traceId := RandomString(4) + TimeHex() + PidHex() + RandomString(4)
	ctx := context.WithValue(context.Background(), traceIdKey, traceId)

	traceIdInCtx := FromContext(ctx)
	if traceIdInCtx != traceId {
		t.Errorf("traceIdInCtx %s != traceId %s", traceIdInCtx, traceId)
	}
}
