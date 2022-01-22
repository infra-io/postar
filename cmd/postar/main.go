// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 01:33:43

package main

import (
	"github.com/FishGoddess/logit"
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/biz"
	"github.com/avinoplan/postar/internal/server"
	"github.com/panjf2000/ants/v2"
)

func loadConfig() (*configs.Config, error) {
	// TODO 加载配置文件初始化配置
	return configs.NewDefaultConfig(), nil
}

func initLogger(c *configs.Config) *logit.Logger {
	return logit.NewLogger()
}

func initPool(c *configs.Config) *ants.Pool {
	pool, err := ants.NewPool(64)
	if err != nil {
		panic(err)
	}
	return pool
}

func main() {
	c, err := loadConfig()
	if err != nil {
		panic(err)
	}

	logger := initLogger(c)
	pool := initPool(c)
	defer pool.Release()

	smtpBiz := biz.NewSMTPBiz(c, logger, pool)
	err = server.NewGRPCServer(c, logger, smtpBiz).Start()
	if err != nil {
		panic(err)
	}
}
