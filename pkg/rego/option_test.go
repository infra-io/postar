// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rego

import (
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithLimit$
func TestWithLimit(t *testing.T) {
	conf := &config{limit: 0}
	WithLimit(64)(conf)

	if conf.limit != 64 {
		t.Fatalf("conf.limit %d is wrong", conf.limit)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithFastFailed$
func TestWithFastFailed(t *testing.T) {
	conf := &config{fastFailed: false}
	WithFastFailed()(conf)

	if !conf.fastFailed {
		t.Fatalf("conf.fastFailed %+v is wrong", conf.fastFailed)
	}
}
