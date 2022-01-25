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
	WorkerNumber int  `ini:"worker_number"` // The number of task worker.
	QueueSize    int  `ini:"queue_size"`    // The max size of task queue.
	Async        bool `ini:"async"`         // The sending mode of task.
	Timeout      int  `ini:"timeout"`       // The sending timeout in millisecond of task.
}

type ServerConfig struct {
	Network string `ini:"network"` // The network of server, see net.Listen.
	Type    string `ini:"type"`    // The type of server.
	Address string `ini:"address"` // The address(including ip and port) of server.
}

type SMTPConfig struct {
	Host     string `ini:"host"`     // The host of smtp server.
	Port     int    `ini:"port"`     // The port of smtp server.
	User     string `ini:"user"`     // The user of smtp server.
	Password string `ini:"password"` // The password of smtp server.
}

// Config stores all configurations of postar.
type Config struct {
	Task   TaskConfig   `ini:"task"`
	Server ServerConfig `ini:"server"`
	SMTP   SMTPConfig   `ini:"smtp"`
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
