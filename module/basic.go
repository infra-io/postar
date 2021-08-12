// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/12 23:00:43

package module

import (
	"github.com/FishGoddess/logit"
)

var (
	globalLogger *logit.Logger

	initializations = []func(config *Config) error{
		initLogger,
	}
)

func initLogger(config *Config) error {
	options := logit.Options()
	globalLogger = logit.NewLogger(
		options.WithCaller(), options.WithTimeFormat("2006-01-02 15:04:05.000"),
	)
	return nil
}

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

func Logger() *logit.Logger {
	return globalLogger
}
