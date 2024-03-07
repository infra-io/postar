// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logging

import (
	"context"

	"github.com/infra-io/postar/pkg/grpc/contextutil"
)

func ResolveRequest(ctx context.Context, request any) []any {
	spaceID := contextutil.GetSpaceID(ctx)
	if spaceID > 0 {
		return []any{"space_id", spaceID}
	}

	return nil
}
