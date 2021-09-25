// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 22:44:59

package trace

import "context"

var (
	traceIdKey struct{} // The key of trace id in context.
)

// WithContext returns a context with a trace id inside.
func WithContext(ctx context.Context) context.Context {
	traceId := RandomString(4) + TimeHex() + PidHex() + RandomString(4)
	return context.WithValue(ctx, traceIdKey, traceId)
}

// FromContext gets the trace id from context.
func FromContext(ctx context.Context) string {
	value := ctx.Value(traceIdKey)
	if value == nil {
		return ""
	}

	traceId, ok := value.(string)
	if !ok {
		return ""
	}
	return traceId
}
