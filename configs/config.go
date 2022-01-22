// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/01/23 01:28:36

package configs

type ServerConfig struct {
	Type    string `int:"type"`    // The type of server.
	Address string `ini:"address"` // The address(including ip and port) of server.
}
type SMTPConfig struct {
	Host     string `int:"host"`     // The host of smtp server.
	Port     int    `int:"port"`     // The port of smtp server.
	User     string `int:"user"`     // The user of smtp server.
	Password string `int:"password"` // The password of smtp server.
}

// Config stores all configurations of postar.
type Config struct {
	Server ServerConfig `int:"server"`
	SMTP   SMTPConfig   `int:"smtp"`
}

// NewDefaultConfig returns a new config.
func NewDefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Type:    "http",
			Address: ":5897",
		},
		SMTP: SMTPConfig{
			Port: 587,
		},
	}
}

// TODO 封装配置查询方法
