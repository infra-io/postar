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
	SenderType string `ini:"sender_type"`
	ServerType string `ini:"server_type"`
}

type LoggerConfig struct {
	Level      string `ini:"level"`
	TimeFormat string `ini:"time_format"`
}

type SenderConfig struct {
	Host               string `ini:"host"`
	Port               int    `ini:"port"`
	User               string `ini:"user"`
	Password           string `ini:"password"`
	WorkerNumber       int    `ini:"work_number"`
	RequestChannelSize int    `ini:"request_channel_size"`
}
type ServerConfig struct {
	Address string `ini:"address"`
}

type Config struct {
	Global *GlobalConfig `ini:"global"`
	Logger *LoggerConfig `ini:"logger"`
	Sender *SenderConfig `ini:"sender"`
	Server *ServerConfig `ini:"server"`
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
		Logger: &LoggerConfig{
			Level:      "info",
			TimeFormat: "2006-01-02 15:04:05.000",
		},
		Sender: &SenderConfig{
			Host:               os.Getenv("POSTAR_SENDER_HOST"),
			Port:               587,
			User:               os.Getenv("POSTAR_SENDER_USER"),
			Password:           os.Getenv("POSTAR_SENDER_PASSWORD"),
			WorkerNumber:       64,
			RequestChannelSize: 65536,
		},
		Server: &ServerConfig{
			Address: "127.0.0.1:5897",
		},
	}
}

type Configurer interface {
	Configure(config *Config) error
}
