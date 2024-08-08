// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"
	"path"
	"time"

	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/pkg/grpc/contextutil"
	"github.com/infra-io/servicex/runtime"
	"github.com/infra-io/servicex/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func parseServiceInfo(info *grpc.UnaryServerInfo) (serviceName string, serviceMethod string) {
	if method := path.Base(info.FullMethod); method != "" {
		return "", method
	}

	return "", info.FullMethod
}

func newLogger(ctx context.Context, info *grpc.UnaryServerInfo, trace *tracing.Trace) *logit.Logger {
	serviceName, serviceMethod := parseServiceInfo(info)
	args := []any{"service_name", serviceName, "service_method", serviceMethod, "trace_id", trace.ID()}

	spaceID := contextutil.GetSpaceID(ctx)
	if spaceID > 0 {
		args = append(args, "space_id", spaceID)
	}

	return logit.FromContext(ctx).With(args...)
}

func Interceptor(serviceName string, timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, serverInfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				logit.FromContext(ctx).Error("recovery from panic", "r", r, "stack", runtime.Stack())
			}
		}()

		beginTime := time.Now()

		trace := tracing.New()
		ctx = tracing.NewContext(ctx, trace)

		logger := newLogger(ctx, serverInfo, trace)
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

		grpc.SetHeader(ctx, metadata.Pairs(contextutil.MetadataKeyTraceID, trace.ID()))
		return handler(ctx, req)
	}
}
