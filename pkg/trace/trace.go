// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package trace

import (
	"context"

	"github.com/avinoplan/postar/pkg/encode"
)

var (
	traceIDKey struct{} // The context key of trace id.
)

// NewTraceID returns a new trace id.
func NewTraceID() string {
	salt := encode.StringHex(6)
	return encode.NowHex() + salt[:3] + encode.PIDHex() + salt[3:]
}

// NewContext wraps ctx with a trace id.
func NewContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

// FromContext gets the trace id from context.
func FromContext(ctx context.Context) string {
	traceID, ok := ctx.Value(traceIDKey).(string)
	if !ok && traceID == "" {
		return ""
	}
	return traceID
}
