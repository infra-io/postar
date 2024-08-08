// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar"
	"github.com/infra-io/postar/config"
	"github.com/infra-io/postar/internal/postar-admin/server"
	"github.com/infra-io/postar/internal/postar-admin/service"
	"github.com/infra-io/postar/internal/postar-admin/store"
	_ "github.com/infra-io/servicex"
)

func parseConfigFile() (string, error) {
	configFile := flag.String("conf", "postar_admin.toml", "The config file of postar-admin.")
	printVersion := flag.Bool("version", false, "Print the version information of postar-admin.")
	flag.Parse()

	if *printVersion {
		fmt.Printf("postar-admin-%s\nos: %s\narch: %s\ngo: %s\n", postar.Version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		os.Exit(0)
	}

	return *configFile, nil
}

func setupConfig() (*config.PostarAdminConfig, error) {
	configFile, err := parseConfigFile()
	if err != nil {
		return nil, err
	}

	conf := config.NewPostarAdminConfig()
	if _, err = toml.DecodeFile(configFile, conf); err != nil {
		return nil, err
	}

	if err = conf.Check(); err != nil {
		return nil, err
	}

	return conf, nil
}

func setupLogger(conf *config.PostarAdminConfig) error {
	opts, err := conf.Logger.Options()
	if err != nil {
		return err
	}

	logger := logit.NewLogger(opts...)
	logit.SetDefault(logger)

	return nil
}

func newServer(conf *config.PostarAdminConfig) (server.Server, error) {
	if err := store.Connect(&conf.Database); err != nil {
		return nil, err
	}

	spaceStore := store.NewSpaceStore(conf)
	accountStore := store.NewAccountStore(conf)
	templateStore := store.NewTemplateStore(conf)

	spaceService := service.NewSpaceService(conf, spaceStore)
	accountService := service.NewAccountService(conf, accountStore)
	templateService := service.NewTemplateService(conf, templateStore)

	return server.New(conf, spaceService, accountService, templateService)
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

	svr, err := newServer(conf)
	if err != nil {
		panic(err)
	}

	if err = svr.Serve(); err != nil {
		panic(err)
	}
}
