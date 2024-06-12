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
	"github.com/infra-io/servicex/observe/logging"
	"github.com/infra-io/servicex/observe/tracing"
	"github.com/infra-io/servicex/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func newServiceMethod(info *grpc.UnaryServerInfo) string {
	if method := path.Base(info.FullMethod); method != "" {
		return method
	}

	return info.FullMethod
}

func Interceptor(serviceName string, timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				logit.FromContext(ctx).Error("recovery from panic", "r", r, "stack", runtime.Stack())
			}
		}()

		defer func() {
			if err != nil {
				err = wrapWithStatus(err)
			}
		}()

		beginTime := time.Now()
		serviceMethod := newServiceMethod(info)
		trace := tracing.New()
		ctx = tracing.NewContext(ctx, trace)

		resolvers := newResolvers(serviceName, serviceMethod, trace)
		logger := logging.NewLogger(ctx, req, resolvers...)
		ctx = logit.NewContext(ctx, logger)

		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()

		reqJson := jsonifyProto(req)
		logger.Debug("service method begin", "request", reqJson)

		defer func() {
			cost := time.Since(beginTime)
			respJson := jsonifyProto(resp)

			if err == nil {
				logger.Debug("service method end", "response", respJson, "cost", cost)
			} else {
				logger.Error("service method end", "err", err, "request", reqJson, "response", respJson, "cost", cost)
			}
		}()

		grpc.SetHeader(ctx, metadata.Pairs(contextutil.MetadataKeyTraceID, trace.ID()))
		return handler(ctx, req)
	}
}
