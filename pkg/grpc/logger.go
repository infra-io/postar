// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"
	"strings"

	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/pkg/grpc/contextutil"
	"google.golang.org/grpc"
)

func parseServiceInfo(info *grpc.UnaryServerInfo) (serviceName string, serviceMethod string) {
	strs := strings.Split(info.FullMethod, "/")
	if len(strs) < 3 {
		return info.FullMethod, info.FullMethod
	}

	last := len(strs) - 1
	serviceMethod = strs[last]

	strs = strings.Split(strs[1], ".")
	if len(strs) < 2 {
		return info.FullMethod, serviceMethod
	}

	last = len(strs) - 1
	serviceName = strs[last]

	return serviceName, serviceMethod
}

func newLogger(ctx context.Context, info *grpc.UnaryServerInfo, traceID string) *logit.Logger {
	serviceName, serviceMethod := parseServiceInfo(info)
	args := []any{"service_name", serviceName, "service_method", serviceMethod, "trace_id", traceID}

	spaceID := contextutil.GetSpaceID(ctx)
	if spaceID > 0 {
		args = append(args, "space_id", spaceID)
	}

	return logit.FromContext(ctx).With(args...)
}
