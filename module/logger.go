// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/16 00:57:03

package module

import (
	"io"
	"os"

	"github.com/FishGoddess/logit"
)

var (
	globalLogger *logit.Logger
)

func loggerLevelOptionFrom(config *Config) logit.Option {

	options := logit.Options()
	if config.Logger.Level == "off" {
		// TODO 替换为 OffLevel
		return options.WithErrorLevel()
	}

	if config.Logger.Level == "error" {
		return options.WithErrorLevel()
	}

	if config.Logger.Level == "warn" {
		return options.WithWarnLevel()
	}

	if config.Logger.Level == "info" {
		return options.WithInfoLevel()
	}
	return options.WithDebugLevel()
}

func outputFilesFrom(config *Config) (io.Writer, io.Writer, error) {

	outputFile, err := os.OpenFile(config.Logger.OutputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, err
	}

	errorOutputFile, err := os.OpenFile(config.Logger.ErrorOutputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, err
	}
	return outputFile, errorOutputFile, nil
}

func initLogger(config *Config) error {

	options := logit.Options()
	outputFile, errorOutputFile, err := outputFilesFrom(config)
	if err != nil {
		return err
	}
	globalLogger = logit.NewLogger(
		loggerLevelOptionFrom(config),
		options.WithTimeFormat(config.Logger.TimeFormat),
		options.WithWriter(outputFile, true),
		options.WithErrorWriter(errorOutputFile, false),
	)
	return nil
}

func Logger() *logit.Logger {
	return globalLogger
}
