// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package contextutil

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
)

const (
	MetadataKeySpaceID    = "postar.space_id"
	MetadataKeySpaceToken = "postar.space_token"
	MetadataKeyTraceID    = "postar.trace_id"
)

func getIncomingMetadata(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	strs := md.Get(key)
	if len(strs) != 1 {
		return ""
	}

	return strs[0]
}

func GetSpaceID(ctx context.Context) int32 {
	id := getIncomingMetadata(ctx, MetadataKeySpaceID)
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
	return getIncomingMetadata(ctx, MetadataKeySpaceToken)
}
