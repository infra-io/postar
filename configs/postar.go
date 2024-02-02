// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package configs

import (
	"time"

	logitconf "github.com/FishGoddess/logit/extension/config"
	timex "github.com/infra-io/servicex/time"
)

// TaskConfig is the config of sending task.
type TaskConfig struct {
	MaxWorkers int  `json:"max_workers" toml:"max_workers"`
	MaxTasks   int  `json:"max_tasks" toml:"max_tasks"`
	Async      bool `json:"async" toml:"async"`
}

// PostarConfig stores all configurations of postar.
type PostarConfig struct {
	Logger   logitconf.Config `json:"logger" toml:"logger"`
	Server   ServerConfig     `json:"server" toml:"server"`
	Database DatabaseConfig   `json:"database" toml:"database"`
	Crypto   CryptoConfig     `json:"crypto" toml:"crypto"`
	Task     TaskConfig       `json:"task" toml:"task"`
}

// NewPostarConfig returns a new config for postar.
func NewPostarConfig() *PostarConfig {
	conf := &PostarConfig{
		Logger: logitconf.Config{
			Level: "debug",
			Writer: logitconf.WriterConfig{
				Target:         "./postar.log",
				FileRotate:     true,
				FileMaxSize:    "128MB",
				FileMaxAge:     "30d",
				FileMaxBackups: 60,
			},
			WithSource: false,
			WithPID:    false,
		},
		Server: ServerConfig{
			Type:             "grpc",
			GrpcEndpoint:     ":5897",
			HttpEndpoint:     ":6897",
			RequestTimeout:   timex.NewDuration(10 * time.Second),
			MaxCloseWaitTime: timex.NewDuration(time.Minute),
		},
		Database: DatabaseConfig{
			Address:         "127.0.0.1:6033",
			Username:        "postar",
			Password:        "123456",
			MaxOpenConns:    64,
			MaxIdleConns:    16,
			ConnMaxLifetime: timex.NewDuration(5 * time.Minute),
			ConnMaxIdleTime: timex.NewDuration(3 * time.Second),
			ReportStatsTime: timex.NewDuration(time.Minute),
		},
		Task: TaskConfig{
			MaxWorkers: 64,
			MaxTasks:   100000,
			Async:      false,
		},
	}

	return conf
}

func (pc *PostarConfig) Check() error {
	return pc.Crypto.Check()
}
