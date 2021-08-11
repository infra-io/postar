// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/12 00:34:45

package main

import "github.com/FishGoddess/logit"

var (
	globalLogger *logit.Logger
)

func InitLogger() {
	options := logit.Options()
	globalLogger = logit.NewLogger(
		options.WithCaller(), options.WithTimeFormat("2006-01-02 15:04:05.000"),
	)
}

func Logger() *logit.Logger {
	return globalLogger
}
