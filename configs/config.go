// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/01/23 01:28:36

package configs

import "time"

type WorkerConfig struct {
	Number  int  `int:"number"`  // The number of worker.
	Async   bool `int:"async"`   // The sending mode of worker.
	Timeout int  `int:"timeout"` // The sending timeout of worker.
}

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
	Worker WorkerConfig `ini:"worker"`
	Server ServerConfig `int:"server"`
	SMTP   SMTPConfig   `int:"smtp"`
}

// NewDefaultConfig returns a new config.
func NewDefaultConfig() *Config {
	return &Config{
		Worker: WorkerConfig{
			Number:  64,
			Async:   false,
			Timeout: 10000,
		},
		Server: ServerConfig{
			Type:    "http",
			Address: ":5897",
		},
		SMTP: SMTPConfig{
			Port: 587,
		},
	}
}

func (c *Config) WorkerNumber() int {
	return c.Worker.Number
}

func (c *Config) WorkerAsync() bool {
	return c.Worker.Async
}

func (c *Config) WorkerTimeout() time.Duration {
	return time.Duration(c.Worker.Timeout) * time.Millisecond
}

func (c *Config) ServerType() string {
	return c.Server.Type
}

func (c *Config) ServerAddress() string {
	return c.Server.Address
}

func (c *Config) SMTPHost() string {
	return c.SMTP.Host
}

func (c *Config) SMTPPort() int {
	return c.SMTP.Port
}

func (c *Config) SMTPUser() string {
	return c.SMTP.User
}

func (c *Config) SMTPPassword() string {
	return c.SMTP.Password
}
