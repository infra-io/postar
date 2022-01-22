// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 22:44:59

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
	salt := encode.RandomString(6)
	return encode.NowHex() + salt[:3] + encode.PID() + salt[3:]
}

// NewContext wraps ctx with a trace id.
func NewContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

// FromContext gets the trace id from context.
func FromContext(ctx context.Context) string {
	traceID, ok := ctx.Value(traceIDKey).(string)
	if !ok && traceID == "" {
		return NewTraceID()
	}
	return traceID
}
