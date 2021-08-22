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
)

func printTakenTime(fn func()) {
	beginTime := time.Now()
	fn()
	endTime := time.Now()
	fmt.Printf("Postar initialized successfully! It took %dms.\n", endTime.Sub(beginTime).Milliseconds())
}

func getConfigFile() string {
	configFile := flag.String("conf", "/opt/postar/conf/postar.ini", "The file path of Postar configuration.")
	flag.Parse()
	return *configFile
}

func main() {

	postar := newPostar()
	printTakenTime(func() {

		config, err := postar.ReadConfig(getConfigFile())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Postar's config:\n%+v\n", config)

		err = postar.Initialize(config)
		if err != nil {
			panic(err)
		}
	})

	err := postar.Run()
	if err != nil {
		panic(err)
	}
}
