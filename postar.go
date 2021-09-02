// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/14 15:00:07

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/avino-plan/postar/module"
	"github.com/avino-plan/postar/module/sender"
	"github.com/avino-plan/postar/module/server"
)

type Postar struct {
	svr server.Server
}

func newPostar() *Postar {
	return &Postar{}
}

func (p *Postar) Initialize(config *module.Config) error {

	err := module.Initialize(config)
	if err != nil {
		return err
	}

	sdr, err := sender.Initialize(config)
	if err != nil {
		return err
	}

	p.svr, err = server.Initialize(config)
	if err != nil {
		return err
	}

	p.svr.ConfigureSender(sdr)
	return nil
}

func (p *Postar) Run() error {
	return p.svr.Serve()
}

func (p *Postar) WaitForShutdown() error {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-signalCh
	fmt.Println("Postar is shutdown gracefully...")
	return p.svr.Close()
}
