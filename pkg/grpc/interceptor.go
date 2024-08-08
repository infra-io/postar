// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"
	"time"

	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/pkg/grpc/contextutil"
	"github.com/infra-io/postar/pkg/runtime"
	"github.com/infra-io/postar/pkg/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Interceptor(serviceName string, timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, serverInfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				logit.FromContext(ctx).Error("recovery from panic", "r", r, "callers", runtime.Callers())
			}
		}()

		beginTime := time.Now()

		traceID := trace.ID()
		ctx = trace.NewContext(ctx, traceID)

		logger := newLogger(ctx, serverInfo, traceID)
		ctx = logit.NewContext(ctx, logger)

		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()

		reqJson := jsonifyProto(req)
		logger.Debug("service method begin", "request", reqJson)

		defer func() {
			respJson := jsonifyProto(resp)
			cost := time.Since(beginTime)

			if err == nil {
				logger.Debug("service method end", "response", respJson, "cost", cost)
			} else {
				logger.Error("service method end", "err", err, "request", reqJson, "response", respJson, "cost", cost)

				err = wrapStatus(err)
			}
		}()

		grpc.SetHeader(ctx, metadata.Pairs(contextutil.MetadataKeyTraceID, traceID))
		return handler(ctx, req)
	}
}
