// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package trace

import (
	"context"
	"testing"
)

// go test -v -cover -run=^TestNewTraceID$
func TestNewTraceID(t *testing.T) {
	traceID := NewTraceID()
	if traceID == "" {
		t.Error("traceID == ''")
	}

	t.Log("traceID:", traceID)
}

// go test -v -cover -run=^TestWithContext$
func TestWithContext(t *testing.T) {
	ctx := NewContext(context.Background(), NewTraceID())

	value := ctx.Value(traceIDKey)
	if value == nil {
		t.Error("ctx.Value returns nil")
	}

	traceID, ok := value.(string)
	if !ok {
		t.Errorf("value %+v isn't string", value)
	}
	t.Log("traceID:", traceID)
}

// go test -v -cover -run=^TestFromContext$
func TestFromContext(t *testing.T) {
	ctx := context.Background()
	traceIDInCtx := FromContext(ctx)
	if traceIDInCtx != "" {
		t.Errorf("traceIDInCtx %s != ''", traceIDInCtx)
	}

	traceID := NewTraceID()
	ctx = context.WithValue(ctx, traceIDKey, traceID)

	traceIDInCtx = FromContext(ctx)
	if traceIDInCtx != traceID {
		t.Errorf("traceIDInCtx %s != traceId %s", traceIDInCtx, traceID)
	}
}
