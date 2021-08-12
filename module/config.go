// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/08/12 22:59:05

package module

import (
	"fmt"
	"os"
)

type GlobalConfig struct {
	SenderType string
	ServerType string
}

type LoggerConfig struct {
}

type SenderConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}
type ServerConfig struct {
	Address string
}

type Config struct {
	Global *GlobalConfig
	Logger *LoggerConfig
	Sender *SenderConfig
	Server *ServerConfig
}

func (c *Config) String() string {
	return fmt.Sprintf("- GlobalConfig is %+v\n- LoggerConfig is %+v\n- SenderConfig is %+v\n- ServerConfig is %+v\n", *c.Global, *c.Logger, *c.Sender, *c.Server)
}

// TODO 加入 os.LookupEnv 获取环境变量配置的功能
func DefaultConfig() *Config {
	return &Config{
		Global: &GlobalConfig{
			SenderType: "smtp",
			ServerType: "http",
		},
		Logger: &LoggerConfig{},
		Sender: &SenderConfig{
			Host:     os.Getenv("POSTAR_SENDER_HOST"),
			Port:     587,
			User:     os.Getenv("POSTAR_SENDER_USER"),
			Password: os.Getenv("POSTAR_SENDER_PASSWORD"),
		},
		Server: &ServerConfig{
			Address: "127.0.0.1:5897",
		},
	}
}

type Configurer interface {
	Configure(config *Config) error
}
