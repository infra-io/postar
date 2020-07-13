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

package config

import (
	"flag"

	"github.com/avino-plan/postar/src/models"
	"gopkg.in/ini.v1"
)

var (
	pathOfConfigFile string
)

func init() {
	flag.StringVar(&pathOfConfigFile, "c", "./postar.ini", "The path of config file.")
	flag.Parse()
}

func UseConfig(useFunc func(config *models.Config)) {
	config := &models.Config{}
	ini.MapTo(config, pathOfConfigFile)
	useFunc(config)
}
