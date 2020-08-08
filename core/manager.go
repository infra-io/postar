// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/13 23:00:26

package core

const (
	Version = "v0.1.2-alpha"
)

var (
	// globalSender is the sender holder for global usage.
	globalSender *sender
)

func init() {
	// Create an email sender with config.
	config := getConfig()
	globalSender = newSender(config.Smtp.Host, config.Smtp.Port, config.Smtp.Username, config.Smtp.Password)
}

// SendSync sends the email and returns an error if failed.
func SendSync(email *Email) error {
	return <-SendAsync(email)
}

// SendAsync sends the email and returns an error if failed.
func SendAsync(email *Email) <-chan error {
	done := make(chan error)
	go func() {
		done <- globalSender.Send(email)
	}()
	return done
}

// ===================================== for fetching settings =====================================

// SystemCommand returns the system command postar will execute.
func SystemCommand() string {
	getConfig() // For initializing config.
	return systemCommand
}

// ServerType returns the type of server in the config.
func ServerType() string {
	return getConfig().Server.Type
}

// ServerPort returns the port of server in the config.
func ServerPort() string {
	return getConfig().Server.Port
}

// ServerClosedPort returns the closedPort of server in the config.
func ServerClosedPort() string {
	return getConfig().Server.ClosedPort
}
