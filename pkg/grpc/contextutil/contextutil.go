// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package contextutil

import (
	"context"
	"strconv"

	"github.com/infra-io/postar/pkg/grpc/gateway"
	grpcx "github.com/infra-io/servicex/net/grpc"
)

func GetSpaceID(ctx context.Context) int32 {
	id := grpcx.GetMetadata(ctx, gateway.MetadataKeySpaceID)
	if id == "" {
		return 0
	}

	spaceID, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}

	return int32(spaceID)
}

func GetSpaceToken(ctx context.Context) string {
	return grpcx.GetMetadata(ctx, gateway.MetadataKeySpaceToken)
}
