// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/01/23 01:28:36

package configs

import "time"

type TaskConfig struct {
	WorkerNumber int  `int:"worker_number"` // The number of task worker.
	QueueSize    int  `int:"queue_size"`    // The max size of task queue.
	Async        bool `int:"async"`         // The sending mode of task.
	Timeout      int  `int:"timeout"`       // The sending timeout in millisecond of task.
}

type ServerConfig struct {
	Network string `int:"network"` // The network of server, see net.Listen.
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
	Task   TaskConfig   `ini:"task"`
	Server ServerConfig `int:"server"`
	SMTP   SMTPConfig   `int:"smtp"`
}

// NewDefaultConfig returns a new config.
func NewDefaultConfig() *Config {
	return &Config{
		Task: TaskConfig{
			WorkerNumber: 64,
			QueueSize:    0,
			Async:        false,
			Timeout:      10000, // 10s
		},
		Server: ServerConfig{
			Network: "tcp",
			Type:    "http",
			Address: ":5897",
		},
		SMTP: SMTPConfig{
			Port: 587,
		},
	}
}

func (c *Config) TaskWorkerNumber() int {
	return c.Task.WorkerNumber
}

func (c *Config) TaskQueueSize() int {
	return c.Task.QueueSize
}

func (c *Config) TaskAsync() bool {
	return c.Task.Async
}

func (c *Config) TaskTimeout() time.Duration {
	return time.Duration(c.Task.Timeout) * time.Millisecond
}

func (c *Config) ServerNetwork() string {
	return c.Server.Network
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
