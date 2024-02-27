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

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithFastFailed$
func TestPool(t *testing.T) {
	acquire := func() (int, error) {
		return 0, nil
	}

	release := func(resource int) error {
		return nil
	}

	pool := New[int](acquire, release, WithLimit(16))
	defer pool.Close()

	go func() {
		for {
			status := pool.Status()
			t.Logf("%+v", status)

			if status.Acquired > pool.limit {
				t.Errorf("status.Acquired %d is wrong", status.Acquired)
				return
			}

			if status.Idle > pool.limit {
				t.Errorf("status.Idle %d is wrong", status.Idle)
				return
			}

			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 1024; i++ {
		resource, err := pool.Take(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		time.Sleep(5 * time.Millisecond)
		pool.Put(resource)

		status := pool.Status()
		if status.Acquired != 1 {
			t.Fatalf("status.Acquired %d is wrong", status.Acquired)
		}

		if status.Idle != 1 {
			t.Fatalf("status.Idle %d is wrong", status.Idle)
		}
	}

	t.Logf("%+v", pool.Status())

	var wg sync.WaitGroup
	for i := 0; i < 4096; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			resource, err := pool.Take(context.Background())
			if err != nil {
				t.Error(err)
				return
			}

			time.Sleep(20 * time.Millisecond)
			pool.Put(resource)

			status := pool.Status()
			if status.Acquired > pool.limit {
				t.Errorf("status.Acquired %d is wrong", status.Acquired)
				return
			}

			if status.Idle > pool.limit {
				t.Errorf("status.Idle %d is wrong", status.Idle)
				return
			}
		}(i)
	}

	wg.Wait()
	t.Logf("%+v", pool.Status())
}
