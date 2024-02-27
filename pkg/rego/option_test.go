// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rego

import (
	"testing"
)

// go test -v -cover -run=^TestWithMaxAcquired$
func TestWithMaxAcquired(t *testing.T) {
	conf := &config{maxAcquired: 0}
	WithMaxAcquired(64)(conf)

	if conf.maxAcquired != 64 {
		t.Errorf("conf.maxAcquired %d is wrong", conf.maxAcquired)
	}
}

// go test -v -cover -run=^TestWithMaxIdle$
func TestWithMaxIdle(t *testing.T) {
	conf := &config{maxIdle: 0}
	WithMaxIdle(16)(conf)

	if conf.maxIdle != 16 {
		t.Errorf("conf.maxIdle %d is wrong", conf.maxIdle)
	}
}

// go test -v -cover -run=^TestWithFastFailed$
func TestWithFastFailed(t *testing.T) {
	conf := &config{fastFailed: false}
	WithFastFailed()(conf)

	if conf.fastFailed {
		t.Errorf("conf.fastFailed %+v is wrong", conf.fastFailed)
	}
}
