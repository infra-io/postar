// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package trace

import "context"

var contextKey = struct{}{}

func NewContext(ctx context.Context, traceID string) context.Context {
	ctx = context.WithValue(ctx, contextKey, traceID)
	return ctx
}

func FromContext(ctx context.Context) string {
	trace, ok := ctx.Value(contextKey).(string)
	if !ok {
		return ""
	}

	return trace
}
