// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/12 23:00:43

package module

const (
	Version = "v0.2.2-alpha"
)

var (
	initializations = []func(config *Config) error{
		initLogger,
	}
)

func Initialize(config *Config) error {

	if config == nil {
		config = DefaultConfig()
	}

	for _, initialize := range initializations {
		err := initialize(config)
		if err != nil {
			return err
		}
	}
	return nil
}
