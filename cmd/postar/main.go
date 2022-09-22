// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/avino-plan/postar/pkg/trace"
	"github.com/go-logit/logit"
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/avino-plan/postar/configs"
	"github.com/avino-plan/postar/internal/biz"
	"github.com/avino-plan/postar/internal/server"
	"github.com/go-ini/ini"
	"github.com/panjf2000/ants/v2"
)

const (
	version = "postar-v0.3.2-alpha"
)

func funnyFunnyChickenHomie() {
	// You know, for funny.
	fmt.Print(`.______     ______        _______.___________.    ___      .______      
|   _  \   /  __  \      /       |           |   /   \     |   _  \     
|  |_)  | |  |  |  |    |   (--- '---|  |----'  /  ^  \    |  |_)  |    
|   ___/  |  |  |  |     \   \       |  |      /  /_\  \   |      /     
|  |      |  '--'  | .----)   |      |  |     /  _____  \  |  |\  \----.
| _|       \______/  |_______/       |__|    /__/     \__\ | _| '._____|

`)
}

func initConfig() (*configs.Config, error) {
	configFile := flag.String("config.file", "postar.ini", "The configuration file of postar.")
	showVersion := flag.Bool("version", false, "Check version of postar.")
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s (%s, %s, %s)\n", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		os.Exit(0)
	}

	c := configs.NewDefaultConfig()
	err := ini.MapTo(c, *configFile)
	return c, err
}

func initLogger(c *configs.Config) *logit.Logger {
	traceInterceptor := func(ctx context.Context, log *logit.Log) {
		traceID := trace.FromContext(ctx)
		if traceID != "" {
			log.String("traceID", traceID)
		}
	}

	logger := logit.NewLogger(logit.Options().WithInterceptors(traceInterceptor))
	logger.SetToGlobal()
	return logger
}

func initPool(c *configs.Config, logger *logit.Logger) *ants.Pool {
	pool, err := ants.NewPool(c.TaskWorkerNumber(), ants.WithMaxBlockingTasks(c.TaskQueueSize()), ants.WithLogger(logger))
	if err != nil {
		panic(err)
	}

	return pool
}

func runServer(c *configs.Config, smtpBiz *biz.SMTPBiz) error {
	cc := *c
	cc.SMTP.Password = "*" // Password is sensitive
	logit.Info("running server").Any("config", cc).Log()

	svr := server.NewServer(c, smtpBiz)
	go func() {
		err := svr.Start()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Postar is running...")

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-signalCh

	fmt.Println("Postar is shutdown gracefully...")
	return svr.Stop()
}

func main() {
	c, err := initConfig()
	if err != nil {
		panic(err)
	}

	logger := initLogger(c)
	defer logger.Close()

	pool := initPool(c, logger)
	defer pool.Release()

	undo, err := maxprocs.Set(maxprocs.Logger(logit.Printf))
	if err != nil {
		undo()
		logit.Error("set maxprocs failed").Error("err", err).Log()
	}

	funnyFunnyChickenHomie()

	err = runServer(c, biz.NewSMTPBiz(c, pool))
	if err != nil {
		logit.Error("stop server failed").Error("err", err).Log()
	}
}
