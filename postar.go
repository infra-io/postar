// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/14 15:00:07

package main

import (
	"os"
	"sync"
	"time"

	"github.com/avino-plan/postar/module"
	"github.com/avino-plan/postar/module/sender"
	"github.com/avino-plan/postar/module/server"
	"github.com/go-ini/ini"
)

type Postar struct {
	svr server.Server
	wg  *sync.WaitGroup
}

func newPostar() *Postar {
	return &Postar{
		wg: &sync.WaitGroup{},
	}
}

func (p *Postar) ReadConfig(configFile string) (*module.Config, error) {
	config := module.DefaultConfig()
	err := ini.MapTo(config, configFile)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	return config, nil
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

	err := p.svr.Serve()
	if err != nil {
		return err
	}

	// TODO 使用 signal 机制通知 Shutdown
	//p.wg.Add(1)
	//p.wg.Wait()
	time.Sleep(time.Hour)
	return nil
}

func (p *Postar) Shutdown() {
	p.wg.Done()
}
