// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rego

import (
	"context"
	"sync"
	"testing"
	"time"
)

// go test -v -cover -run=^TestPool$
func TestPool(t *testing.T) {
	acquire := func() (int, error) {
		return 0, nil
	}

	release := func(resource int) error {
		return nil
	}

	pool := New[int](acquire, release, WithMaxAcquired(16), WithMaxIdle(8))
	defer pool.Close()

	test := func(i int) {
		resource, err := pool.Take(context.Background())
		if err != nil {
			t.Error(err)
		}

		defer pool.Put(resource)
		time.Sleep(time.Millisecond)
	}

	for i := 0; i < 4096; i++ {
		test(i)

		status := pool.Status()
		if status.Acquired != 1 {
			t.Errorf("status.Acquired %d is wrong", status.Acquired)
		}

		if status.Idle != 1 {
			t.Errorf("status.Idle %d is wrong", status.Idle)
		}
	}

	t.Logf("%+v", pool.Status())

	var wg sync.WaitGroup
	for i := 0; i < 4096; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			test(i)

			status := pool.Status()
			if status.Acquired > pool.maxAcquired {
				t.Errorf("status.Acquired %d is wrong", status.Acquired)
			}

			if status.Idle > pool.maxIdle {
				t.Errorf("status.Idle %d is wrong", status.Idle)
			}
		}(i)
	}

	wg.Wait()
	t.Logf("%+v", pool.Status())
}
