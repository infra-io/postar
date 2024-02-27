// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rego

import (
	"context"
	"errors"
	"sync"
)

var (
	errPoolIsFull   = errors.New("rego: pool is full")
	errPoolIsClosed = errors.New("rego: pool is closed")
)

type Status struct {
	// Acquired is the count of acquired resources.
	Acquired uint64 `json:"acquired"`

	// Idle is the count of idle resources.
	Idle uint64 `json:"idle"`

	// Waiting is the count of waiting for a resource.
	Waiting uint64 `json:"waiting"`
}

// AcquireFunc is a function acquires a new resource and returns error if failed.
type AcquireFunc[T any] func() (T, error)

// ReleaseFunc is a function releases a resource and returns error if failed.
type ReleaseFunc[T any] func(resource T) error

// DefaultReleaseFunc is a default func to release a resource.
// It does nothing to the resource.
func DefaultReleaseFunc[T any](resource T) error {
	return nil
}

type Pool[T any] struct {
	config

	acquire AcquireFunc[T]
	release ReleaseFunc[T]

	resources chan T
	status    Status
	closed    bool

	lock sync.RWMutex
}

func New[T any](acquire AcquireFunc[T], release ReleaseFunc[T], opts ...Option) *Pool[T] {
	if acquire == nil || release == nil {
		panic("rego: acquire or release func can't be nil")
	}

	conf := newDefaultConfig()
	for _, opt := range opts {
		opt.ApplyTo(conf)
	}

	pool := &Pool[T]{
		config:    *conf,
		acquire:   acquire,
		release:   release,
		resources: make(chan T, conf.maxAcquired),
		closed:    false,
	}

	return pool
}

func (p *Pool[T]) releaseResource(resource T) error {
	if p.release == nil {
		return nil
	}

	return p.release(resource)
}

func (p *Pool[T]) Put(resource T) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.closed {
		return p.releaseResource(resource)
	}

	// Only waiting count < idle count will close the client immediately.
	if p.status.Waiting < p.status.Idle && p.status.Idle >= p.maxIdle {
		p.status.Acquired--
		return p.releaseResource(resource)
	}

	select {
	case p.resources <- resource:
		p.status.Idle++
		return nil
	default:
		return p.releaseResource(resource)
	}
}

func (p *Pool[T]) tryToTake() (resource T, ok bool) {
	select {
	case resource = <-p.resources:
		return resource, true
	default:
		return resource, false
	}
}

// waitToTake waits to take an idle resource from pool.
// Record: Add ctx.Done() to select will cause a performance problem...
// The select will call runtime.selectgo if there are more than one case in it, and runtime.selectgo has two steps which is slow:
//
//	sellock(scases, lockorder)
//	sg := acquireSudog()
//
// We don't know what to do yet, but we think timeout mechanism should be supported even we haven't solved it.
func (p *Pool[T]) waitToTake(ctx context.Context) (resource T, err error) {
	select {
	case resource = <-p.resources:
		return resource, nil
	case <-ctx.Done():
		return resource, ctx.Err()
	}
}

// Take takes a resource from pool and returns an error if failed.
// You should call pool.Put to put a taken resource back to the pool.
// We recommend you to use a defer for putting a resource safely.
func (p *Pool[T]) Take(ctx context.Context) (resource T, err error) {
	p.lock.Lock()
	if p.closed {
		p.lock.Unlock()

		return resource, errPoolIsClosed
	}

	var ok bool
	if resource, ok = p.tryToTake(); ok {
		p.status.Idle--
		p.lock.Unlock()

		return resource, nil
	}

	if p.status.Acquired < p.maxAcquired {
		p.status.Acquired++
		p.lock.Unlock()

		// Increase the acquired and unlock before new client may cause the pool becomes full in advance.
		// So we should decrease the acquired if acquired failed.
		defer func() {
			if err != nil {
				p.lock.Lock()
				p.status.Acquired--
				p.lock.Unlock()
			}
		}()

		return p.acquire()
	}

	if p.fastFailed {
		p.lock.Unlock()
		return resource, errPoolIsFull
	}

	p.status.Waiting++
	p.lock.Unlock()

	resource, err = p.waitToTake(ctx)

	p.lock.Lock()
	defer p.lock.Unlock()

	if err == nil {
		p.status.Idle--
	}

	p.status.Waiting--
	return resource, err
}

// Status returns the status of the pool.
func (p *Pool[T]) Status() Status {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.status
}

func (p *Pool[T]) releaseResources() error {
	for i := uint64(0); i < p.status.Acquired; i++ {
		resource := <-p.resources
		if err := p.release(resource); err != nil {
			return err
		}
	}

	return nil
}

// Close closes pool and releases all resources.
func (p *Pool[T]) Close() error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.closed {
		return nil
	}

	if err := p.releaseResources(); err != nil {
		return err
	}

	p.closed = true
	p.status = Status{}

	close(p.resources)
	return nil
}
