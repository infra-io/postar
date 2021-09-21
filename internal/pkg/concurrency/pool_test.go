// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 23:07:51

package concurrency

import (
	"context"
	"testing"
	"time"
)

// go test -v -cover -run=^TestNewPool$
func TestNewPool(t *testing.T) {
	ctx := context.Background()
	pool := NewPool(WithMaxWorkers(4), WithMaxWorkerTasks(16)).Start()

	numbers := [1000]int{}
	for i := 0; i < 1000; i++ {
		no := i
		errorCh := pool.Go(ctx, func(ctx context.Context) error {
			numbers[no] = no
			return nil
		})
		if err := <-errorCh; err != nil {
			t.Error(err)
		}
	}

	time.Sleep(time.Second)
	pool.Stop()

	for i, num := range numbers {
		if i != num {
			t.Errorf("i %d != num %d", i, num)
		}
	}
}
