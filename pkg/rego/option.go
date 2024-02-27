// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rego

type config struct {
	limit      uint64
	fastFailed bool
}

func newDefaultConfig() *config {
	conf := &config{
		limit:      64,
		fastFailed: false,
	}

	return conf
}

type Option func(conf *config)

func (o Option) ApplyTo(conf *config) {
	o(conf)
}

// WithLimit sets limit to config.
func WithLimit(limit uint64) Option {
	return func(conf *config) {
		conf.limit = limit
	}
}

// WithFastFailed sets fastFailed to config.
func WithFastFailed() Option {
	return func(conf *config) {
		conf.fastFailed = true
	}
}
