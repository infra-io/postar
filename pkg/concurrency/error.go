// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/17 01:35:47

package concurrency

import "errors"

var (
	ErrEnqueueTimeout = errors.New("pool: enqueue a task to pool timeout") // An error: enqueue a task to pool timeout.
)

// IsEnqueueTimeout returns if the err is ErrEnqueueTimeout.
func IsEnqueueTimeout(err error) bool {
	return err == ErrEnqueueTimeout
}
