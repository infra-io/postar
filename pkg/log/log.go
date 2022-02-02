// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 02:05:02

package log

import (
	"github.com/go-logit/logit"
	"github.com/avinoplan/postar/configs"
)

var (
	// globalLogger is the logger using in global.
	globalLogger *logit.Logger
)

// Initialize initializes log package with config.
func Initialize(c *configs.Config) error {
	options := logit.Options()
	globalLogger = logit.NewLogger(
		options.WithCallerDepth(4),
	)
	return nil
}

// Logger returns the global logger.
func Logger() *logit.Logger {
	return globalLogger
}

// Debug returns a Log with debug level if debug level is enabled.
func Debug(msg string, args ...interface{}) *logit.Log {
	return globalLogger.Debug(msg, args...)
}

// Info returns a Log with info level if info level is enabled.
func Info(msg string, args ...interface{}) *logit.Log {
	return globalLogger.Info(msg, args...)
}

// Warn returns a Log with warn level if warn level is enabled.
func Warn(msg string, args ...interface{}) *logit.Log {
	return globalLogger.Warn(msg, args...)
}

// Error adds an entry which key is string and value is error type to l.
func Error(err error, msg string, args ...interface{}) *logit.Log {
	return globalLogger.Error(msg, args...).Error("err", err)
}

// Printf prints a log as info level.
func Printf(msg string, args ...interface{}) {
	globalLogger.Info(msg, args...).End()
}

// Close closes logger in log.
func Close() error {
	return globalLogger.Close()
}
