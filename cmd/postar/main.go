// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 01:33:43

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/FishGoddess/logit"
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/biz"
	"github.com/avinoplan/postar/internal/server"
	"github.com/go-ini/ini"
	"github.com/panjf2000/ants/v2"
)

const (
	version = "postar-v0.3.0-alpha"
)

func loadConfig() (*configs.Config, error) {
	configFile := flag.String("config.file", "postar.ini", "The configuration file of postar.")
	showVersion := flag.Bool("version", false, "Check version of postar.")
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s (%s, %s, %s)\n", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		os.Exit(0)
	}

	c := configs.NewDefaultConfig()
	err := ini.MapTo(c, *configFile)
	if err != nil {
		return nil, err
	}

	log.Printf("Load config %+v\n", *c)
	return c, nil
}

func initLogger(c *configs.Config) *logit.Logger {
	return logit.NewLogger()
}

func initPool(c *configs.Config) *ants.Pool {
	pool, err := ants.NewPool(c.TaskWorkerNumber(), ants.WithMaxBlockingTasks(c.TaskQueueSize()))
	if err != nil {
		panic(err)
	}
	return pool
}

func runServer(c *configs.Config, logger *logit.Logger, smtpBiz *biz.SMTPBiz) {
	svr := server.NewServer(c, logger, smtpBiz)

	go func() {
		err := svr.Start()
		if err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-signalCh
	fmt.Println("Postar is shutdown gracefully...")
	svr.Stop()
}

func main() {
	c, err := loadConfig()
	if err != nil {
		panic(err)
	}

	logger := initLogger(c)
	defer logger.Close()

	pool := initPool(c)
	defer pool.Release()

	smtpBiz := biz.NewSMTPBiz(c, logger, pool)
	runServer(c, logger, smtpBiz)
}
