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
// Created at 2020/07/13 22:38:33

package core

import (
	"flag"
	"sync"

	"github.com/FishGoddess/logit"
	"gopkg.in/ini.v1"
)

var (
	// globalConfig is the config holder for global usage.
	// This holder is singleton, so it uses initConfigOnce to do that.
	globalConfig   = &config{}
	initConfigOnce = &sync.Once{}
)

// config is the struct represents of all settings of this system.
type config struct {
	Smtp smtpConfig `ini:"smtp"`
}

// smtpConfig is the struct represents of all settings of smtp.
type smtpConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// getConfig returns the global config.
func getConfig() *config {

	// Only init config once.
	initConfigOnce.Do(func() {
		pathOfConfigFile := flag.String("c", "./postar.ini", "The path of config file.")
		flag.Parse()
		err := ini.MapTo(globalConfig, *pathOfConfigFile)
		if err != nil {
			logit.Warnf("Can't map globalConfig to path: %s. Using default config.", *pathOfConfigFile)
		}
	})
	return globalConfig
}
