// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/11 00:22:42

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {

	beginTime := time.Now()
	InitLogger()

	sender := NewSmtpSender("smtp.office365.com", 587, os.Args[1], os.Args[2])
	defer sender.Close()

	server := NewHttpServer(sender)
	defer server.Close()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err := server.Serve("127.0.0.1:5897")
		if err != nil {
			Logger().Error("new server failed").Error("err", err).End()
			panic(err)
		}
	}()

	endTime := time.Now()
	fmt.Printf("Postar started successfully! It took %dms.\n", endTime.Sub(beginTime).Milliseconds())
	wg.Wait()
}
