// Copyright 2021 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar"
	"github.com/infra-io/postar/internal/postar/server"
	"github.com/infra-io/postar/internal/postar/service"
	"github.com/infra-io/postar/internal/postar/store"
	"github.com/infra-io/servicex/runtime/maxprocs"

	"github.com/infra-io/postar/configs"
)

func parseConfigFile() (string, error) {
	configFile := flag.String("conf", "postar.toml", "The config file of postar.")
	printVersion := flag.Bool("version", false, "Print the version information of postar.")
	flag.Parse()

	if *printVersion {
		fmt.Printf("postar-%s\nos: %s\narch: %s\ngo: %s\n", postar.Version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		os.Exit(0)
	}

	return *configFile, nil
}

func setupConfig() (*configs.PostarConfig, error) {
	configFile, err := parseConfigFile()
	if err != nil {
		return nil, err
	}

	conf := configs.NewPostarConfig()
	if _, err = toml.DecodeFile(configFile, conf); err != nil {
		return nil, err
	}

	if err = conf.Check(); err != nil {
		return nil, err
	}

	return conf, nil
}

func setupLogger(conf *configs.PostarConfig) error {
	opts, err := conf.Logger.Options()
	if err != nil {
		return err
	}

	logger, err := logit.NewLoggerGracefully(opts...)
	if err != nil {
		return err
	}

	logit.SetDefault(logger)
	return nil
}

func newServer(conf *configs.PostarConfig) (server.Server, io.Closer, error) {
	if err := store.Connect(&conf.Database); err != nil {
		return nil, nil, err
	}

	spaceStore := store.NewSpaceStore(conf)
	accountStore := store.NewAccountStore(conf)
	templateStore := store.NewTemplateStore(conf)
	emailService := service.NewEmailService(conf, spaceStore, accountStore, templateStore)

	server, err := server.New(conf, emailService)
	if err != nil {
		return nil, nil, err
	}

	return server, emailService, nil
}

func main() {
	conf, err := setupConfig()
	if err != nil {
		panic(err)
	}

	if err = setupLogger(conf); err != nil {
		panic(err)
	}

	logit.Info("using config", "conf", conf)
	defer logit.Close()

	// Setup process information automatically.
	maxprocs.Setup()

	svr, closer, err := newServer(conf)
	if err != nil {
		panic(err)
	}

	defer closer.Close()

	if err = svr.Serve(); err != nil {
		panic(err)
	}
}
