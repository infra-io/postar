// Copyright 2020 Ye Zi Jie. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/13 23:00:26

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

	// globalSender is the sender holder for global usage.
	globalSender *sender
)

func init() {
	// Create an email sender with config.
	config := getConfig()
	globalSender = newSender(config.Smtp.Host, config.Smtp.Port, config.Smtp.Username, config.Smtp.Password)
}

// Log returns the global logger.
func Log() *logit.Logger {
	initLoggerOnce.Do(func() {
		globalLogger = logit.Me()
	})
	return globalLogger
}

// Send sends the email and returns an error if failed.
func Send(email *Email) error {
	return globalSender.Send(email)
}
