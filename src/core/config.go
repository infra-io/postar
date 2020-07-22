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

	"gopkg.in/ini.v1"
)

var (
	// globalConfig is the config holder for global usage.
	// This holder is singleton, so it uses initConfigOnce to do that.
	globalConfig   = defaultConfig()
	initConfigOnce = &sync.Once{}

	// systemCommand is the system command postar will execute.
	systemCommand string
)

// config is the struct represents of all settings of this system.
type config struct {
	Smtp   *smtpConfig   `ini:"smtp"`
	Server *serverConfig `ini:"server"`
}

// smtpConfig is the struct represents of all settings of smtp.
type smtpConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// serverConfig is the struct represents of all settings of server.
type serverConfig struct {
	Type       string `ini:"type"`
	Port       string `ini:"port"`
	ClosedPort string `ini:"closedPort"`
}

// defaultConfig returns a default config for use.
func defaultConfig() *config {
	return &config{
		Smtp: &smtpConfig{
			Port: 587,
		},
		Server: &serverConfig{
			Type:       "http",
			Port:       "5779",
			ClosedPort: "5780",
		},
	}
}

// ensureGlobalConfigIsValid will check the global config and do all the prepared jobs
// if it's not ready. You should know that the global config is Only initialized once.
func ensureGlobalConfigIsValid() {
	initConfigOnce.Do(func() {

		// Parse flags and get the path of config file.
		flag.StringVar(&systemCommand, "system", "boot", "Execute a system's command.")
		pathOfConfigFile := flag.String("c", "./postar.ini", "The path of config file.")
		flag.Parse()

		// Map config file to global config.
		err := ini.MapTo(globalConfig, *pathOfConfigFile)
		if err != nil {
			Logger().Warnf("Can't map globalConfig to path: %s. Using default config.", *pathOfConfigFile)
		}
	})
}

// getConfig returns the global config.
func getConfig() *config {
	ensureGlobalConfigIsValid()
	return globalConfig
}
