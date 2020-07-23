// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/16 23:09:31

package core

import (
	"sync"

	"github.com/FishGoddess/logit"
)

var (
	// globalLogger is the logger holder for global usage.
	// This holder is singleton, so it uses initLoggerOnce to do that.
	globalLogger   *logit.Logger
	initLoggerOnce = &sync.Once{}
)

// Logger returns the global logger.
func Logger() *logit.Logger {
	initLoggerOnce.Do(func() {
		globalLogger = logit.Me()
	})
	return globalLogger
}
