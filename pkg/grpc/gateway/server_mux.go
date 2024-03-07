// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gateway

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewServeMux() *runtime.ServeMux {
	jpb := &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:     true,
			UseEnumNumbers:    true,
			EmitUnpopulated:   true,
			EmitDefaultValues: true,
		},
	}

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(matchRequestHeader),
		runtime.WithOutgoingHeaderMatcher(matchResponseHeader),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jpb),
	)

	return mux
}
