// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"testing"

	"google.golang.org/grpc"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNewServiceMethod$
func TestNewServiceMethod(t *testing.T) {
	info := &grpc.UnaryServerInfo{
		FullMethod: "/package.service/method",
	}

	method := newServiceMethod(info)
	if method != "method" {
		t.Errorf("method %s is wrong", method)
	}
}
