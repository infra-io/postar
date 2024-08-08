// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package trace

import (
	"context"
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNewContext$
func TestNewContext(t *testing.T) {
	traceID := ID()
	ctx := NewContext(context.Background(), traceID)

	value := ctx.Value(contextKey)
	if value == nil {
		t.Error("ctx.Value returns nil")
	}

	if value.(string) != traceID {
		t.Errorf("value %s != trace %s", value, traceID)
	}

	t.Log("traceID:", traceID)
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestFromContext$
func TestFromContext(t *testing.T) {
	ctx := context.Background()

	traceID := FromContext(ctx)
	if traceID != "" {
		t.Errorf("traceID %s != ''", traceID)
	}

	newTraceID := ID()
	ctx = context.WithValue(ctx, contextKey, newTraceID)

	traceID = FromContext(ctx)
	if traceID != newTraceID {
		t.Errorf("traceID %s != newTraceID %s", traceID, newTraceID)
	}

	t.Log("traceID:", traceID)
}
