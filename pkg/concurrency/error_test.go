// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/17 01:36:24

package concurrency

import "testing"

// go test -v -cover -run=^TestIsEnqueueTimeout$
func TestIsEnqueueTimeout(t *testing.T) {

	err := ErrEnqueueTimeout
	if !IsEnqueueTimeout(err) {
		t.Error("err should be ErrEnqueueTimeout")
	}

	err = nil
	if IsEnqueueTimeout(err) {
		t.Error("err should not be ErrEnqueueTimeout")
	}
}
