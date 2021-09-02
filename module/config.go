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
	"runtime"
	"strconv"
	"strings"
)

type GlobalConfig struct {
	SenderType string `ini:"sender_type"`
	ServerType string `ini:"server_type"`
}

type LoggerConfig struct {
	Level           string `ini:"level"`
	TimeFormat      string `ini:"time_format"`
	OutputFile      string `ini:"output_file"`
	ErrorOutputFile string `ini:"error_output_file"`
}

type SenderConfig struct {
	SmtpHost           string `ini:"smtp_host"`
	SmtpPort           int    `ini:"smtp_port"`
	SmtpUser           string `ini:"smtp_user"`
	SmtpPassword       string `ini:"smtp_password"`
	WorkerNumber       int    `ini:"worker_number"`
	RequestChannelSize int    `ini:"request_channel_size"`
}

type ServerConfig struct {
	Address        string `ini:"address"`
	UseHttp2       bool   `ini:"use_http2"`
	TLSCert        string `ini:"tls_cert"`
	TLSCertKey     string `ini:"tls_cert_key"`
	WaitForClosing int    `int:"wait_for_closing"`
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

func DefaultConfig() *Config {

	logFile := "/opt/postar/log/postar.log"
	errorLogFile := "/opt/postar/log/postar.error.log"
	if strings.Contains(runtime.GOOS, "windows") {
		logFile = "../log/postar.log"
		errorLogFile = "../log/postar.error.log"
	}

	tlsCert := "/opt/postar/conf/postar_cert.pem"
	tlsCertKey := "/opt/postar/conf/postar_cert.key"
	if strings.Contains(runtime.GOOS, "windows") {
		tlsCert = "../conf/postar_cert.pem"
		tlsCertKey = "../conf/postar_cert.key"
	}

	port, err := strconv.Atoi(os.Getenv("POSTAR_SMTP_PORT"))
	if err != nil {
		port = 587
	}
	return &Config{
		Global: &GlobalConfig{
			SenderType: "smtp",
			ServerType: "http",
		},
		Logger: &LoggerConfig{
			Level:           "info",
			TimeFormat:      "2006-01-02 15:04:05.000",
			OutputFile:      logFile,
			ErrorOutputFile: errorLogFile,
		},
		Sender: &SenderConfig{
			SmtpHost:           os.Getenv("POSTAR_SMTP_HOST"),
			SmtpPort:           port,
			SmtpUser:           os.Getenv("POSTAR_SMTP_USER"),
			SmtpPassword:       os.Getenv("POSTAR_SMTP_PASSWORD"),
			WorkerNumber:       64,
			RequestChannelSize: 65536,
		},
		Server: &ServerConfig{
			Address:        ":5897",
			UseHttp2:       false,
			TLSCert:        tlsCert,
			TLSCertKey:     tlsCertKey,
			WaitForClosing: 30, // 30s
		},
	}
}

type Configurer interface {
	Configure(config *Config) error
}
