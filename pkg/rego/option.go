// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rego

type config struct {
	maxAcquired uint64
	maxIdle     uint64
	fastFailed  bool
}

func newDefaultConfig() *config {
	conf := &config{
		maxAcquired: 128,
		maxIdle:     64,
		fastFailed:  false,
	}

	return conf
}

type Option func(conf *config)

func (o Option) ApplyTo(conf *config) {
	o(conf)
}

// WithMaxAcquired sets maxAcquired to config.
func WithMaxAcquired(maxAcquired uint64) Option {
	return func(conf *config) {
		conf.maxAcquired = maxAcquired
	}
}

// WithMaxIdle sets maxIdle to config.
func WithMaxIdle(maxIdle uint64) Option {
	return func(conf *config) {
		conf.maxIdle = maxIdle
	}
}

// WithFastFailed sets fastFailed to config.
func WithFastFailed() Option {
	return func(conf *config) {
		conf.fastFailed = true
	}
}
