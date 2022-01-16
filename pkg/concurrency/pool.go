// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 22:05:23

package concurrency

import (
	"context"
	"sync"
	"sync/atomic"
)

// PoolOptions is the options of Pool.
type PoolOptions func(pool *Pool)

// WithMaxWorkers sets maxWorkers of pool.
func WithMaxWorkers(maxWorkers int) PoolOptions {
	return func(pool *Pool) {
		pool.maxWorkers = maxWorkers
		pool.taskQueues = make([]chan func(), maxWorkers)
	}
}

// WithMaxWorkerTasks sets maxWorkerTasks of pool.
func WithMaxWorkerTasks(maxWorkerTasks int) PoolOptions {
	return func(pool *Pool) {
		pool.maxWorkerTasks = maxWorkerTasks
	}
}

// Pool is a set of goroutines.
// This is for controlling the number of goroutines and managing their lifecycles.
type Pool struct {
	// maxWorkers is the max number of workers.
	maxWorkers int

	// maxWorkerTasks is the max number of one task queue.
	// The Go() method will block until the task queue has enough capacity.
	maxWorkerTasks int

	// currentQueue records the current sequence of task queues.
	currentQueue int64

	// taskQueues stores all task queues.
	// Every worker has its own task queue.
	taskQueues []chan func()

	// onRecover is a function which will call in defer after finishing task.
	onRecover func(cause interface{})

	// wg is for managing all goroutines.
	wg *sync.WaitGroup
}

// NewPool returns a new Pool holder.
func NewPool(options ...PoolOptions) *Pool {
	pool := &Pool{
		maxWorkers:     64,
		maxWorkerTasks: 1024,
		taskQueues:     make([]chan func(), 64),
		currentQueue:   -1,
		wg:             &sync.WaitGroup{},
	}

	for _, option := range options {
		option(pool)
	}
	return pool
}

// OnRecover sets onRecover to p.
func (p *Pool) OnRecover(onRecover func(cause interface{})) {
	p.onRecover = onRecover
}

// Start starts all workers in pool and starts receiving tasks.
func (p *Pool) Start() *Pool {
	for i := 0; i < p.maxWorkers; i++ {
		taskQueue := make(chan func(), p.maxWorkerTasks)
		p.taskQueues[i] = taskQueue

		p.wg.Add(1)
		go func() {
			defer p.wg.Done()

			for task := range taskQueue {
				task()
			}
		}()
	}
	return p
}

// nextTaskQueue returns next task queue.
func (p *Pool) nextTaskQueue() chan<- func() {
	currentQueue := atomic.AddInt64(&p.currentQueue, 1)
	if currentQueue >= int64(p.maxWorkers)-1 {
		atomic.StoreInt64(&p.currentQueue, -1)
	}
	return p.taskQueues[currentQueue]
}

// wrapTask wraps task to a complete task of pool.
func (p *Pool) wrapTask(ctx context.Context, fn func(ctx context.Context) error, errorCh chan<- error) func() {
	return func() {
		var err error
		defer func() {
			errorCh <- err
			if cause := recover(); p.onRecover != nil {
				p.onRecover(cause) // Notice: Just do some simple records which don't panic...
			}
		}()
		err = fn(ctx)
	}
}

// Go sends the task to task queue and waits for executing.
func (p *Pool) Go(ctx context.Context, fn func(ctx context.Context) error) <-chan error {
	errorCh := make(chan error, 1)

	taskQueue := p.nextTaskQueue()
	select {
	case taskQueue <- p.wrapTask(ctx, fn, errorCh):
	case <-ctx.Done():
		errorCh <- ctx.Err()
	}
	return errorCh
}

// Stop closes all task queues and waits for all workers to be shutdown.
func (p *Pool) Stop() {
	for _, taskQueue := range p.taskQueues {
		close(taskQueue)
	}
	p.wg.Wait()
}
