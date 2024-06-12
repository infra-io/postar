// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"

	"github.com/infra-io/postar/pkg/grpc/contextutil"
	"github.com/infra-io/servicex/observe/logging"
	"github.com/infra-io/servicex/observe/tracing"
)

func spaceResolver(ctx context.Context, _ any) []any {
	spaceID := contextutil.GetSpaceID(ctx)
	if spaceID > 0 {
		return []any{"space_id", spaceID}
	}

	return nil
}

func newResolvers(serviceName string, serviceMethod string, trace *tracing.Trace) []logging.ArgsResolver {
	return []logging.ArgsResolver{
		logging.ServiceResolver(serviceName, serviceMethod), logging.TraceResolver(trace), spaceResolver,
	}
}
