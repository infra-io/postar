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
	"io/ioutil"
	"runtime"
	"strings"
	"time"

	"github.com/avino-plan/postar/module"
	"github.com/go-ini/ini"
)

func configFile() *module.Config {

	defaultConfigFile := "/opt/postar/conf/postar.ini"
	if strings.Contains(runtime.GOOS, "windows") {
		defaultConfigFile = "../conf/postar.ini"
	}

	configFile := flag.String("conf", defaultConfigFile, "The file path of Postar configuration.")
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

func recordStartError(msg string, err error) {
	msg = fmt.Sprintf("[%s] [%s] %+v\n", time.Now().Format("2006-01-02 15:04:05.000"), msg, err)
	ioutil.WriteFile("./start_error.log", []byte(msg), 0644)
}

func main() {

	beginTime := time.Now()

	postar := newPostar()
	err := postar.Initialize(configFile())
	if err != nil {
		recordStartError("Initialize error", err)
		panic(err)
	}

	endTime := time.Now()
	fmt.Printf("Postar initialized successfully! It took %dms.\n", endTime.Sub(beginTime).Milliseconds())

	go func() {
		err = postar.Run()
		if err != nil {
			recordStartError("Run error", err)
			panic(err)
		}
	}()

	err = postar.WaitForShutdown()
	if err != nil {
		recordStartError("Shutdown error", err)
		panic(err)
	}
}
