// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/01/25 23:24:00

package model

import (
	"testing"
	"time"

	"github.com/avinoplan/postar/configs"
)

// go test -v -cover -run=^TestDefaultSendEmailOptions$
func TestDefaultSendEmailOptions(t *testing.T) {
	opts := DefaultSendEmailOptions(nil)
	if opts.Async {
		t.Errorf("opts.Async %+v is wrong", opts.Async)
	}

	if opts.Timeout != 10*time.Second {
		t.Errorf("opts.Timeout %d != 10 * time.Second", opts.Timeout)
	}

	c := configs.NewDefaultConfig()
	opts = DefaultSendEmailOptions(c)
	if opts.Async != c.TaskAsync() {
		t.Errorf("opts.Async %+v != c.TaskAsync() %+v", opts.Async, c.TaskAsync())
	}

	if opts.Timeout != c.TaskTimeout() {
		t.Errorf("opts.Timeout %d != c.TaskTimeout() %d", opts.Timeout, c.TaskTimeout())
	}
}
