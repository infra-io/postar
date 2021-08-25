// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/11 00:22:42

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/avino-plan/postar/module"
	"github.com/go-ini/ini"
)

func configFile() *module.Config {

	configFile := flag.String("conf", "/opt/postar/conf/postar.ini", "The file path of Postar configuration.")
	flag.Parse()
	fmt.Printf("Postar got config file: %s\n", *configFile)

	config := module.DefaultConfig()
	if *configFile == "" {
		return config
	}

	err := ini.MapTo(config, *configFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Postar's config:\n%+v\n", config)
	return config
}

func main() {

	beginTime := time.Now()

	postar := newPostar()
	err := postar.Initialize(configFile())
	if err != nil {
		panic(err)
	}

	endTime := time.Now()
	fmt.Printf("Postar initialized successfully! It took %dms.\n", endTime.Sub(beginTime).Milliseconds())

	err = postar.Run()
	if err != nil {
		panic(err)
	}
}
