// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gateway

import (
	grpcx "github.com/infra-io/servicex/net/grpc"
)

const (
	HeaderKeySpaceID    = "X-Postar-Space-Id"
	HeaderKeySpaceToken = "X-Postar-Space-Token"
	HeaderKeyTraceID    = "X-Postar-Trace-Id"
)

const (
	MetadataKeySpaceID    = "service.space_id"
	MetadataKeySpaceToken = "service.space_token"
)

func matchRequestHeader(key string) (string, bool) {
	if key == HeaderKeySpaceID {
		return MetadataKeySpaceID, true
	}

	if key == HeaderKeySpaceToken {
		return MetadataKeySpaceToken, true
	}

	return key, false
}

func matchResponseHeader(key string) (string, bool) {
	if key == grpcx.ServiceKeyTraceID {
		return HeaderKeyTraceID, true
	}

	return key, false
}
