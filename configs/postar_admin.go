// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package configs

import (
	"time"

	logitconf "github.com/FishGoddess/logit/extension/config"
	timex "github.com/infra-io/servicex/time"
)

// PostarAdminConfig stores all configurations of postar-admin.
type PostarAdminConfig struct {
	Logger   logitconf.Config `json:"logger" toml:"logger"`
	Server   ServerConfig     `json:"server" toml:"server"`
	Database DatabaseConfig   `json:"database" toml:"database"`
	Crypto   CryptoConfig     `json:"crypto" toml:"crypto"`
}

// NewPostarAdminConfig returns a new config for postar-admin.
func NewPostarAdminConfig() *PostarAdminConfig {
	conf := &PostarAdminConfig{
		Logger: logitconf.Config{
			Level: "debug",
			Writer: logitconf.WriterConfig{
				Target:         "./postar_admin.log",
				FileRotate:     true,
				FileMaxSize:    "128MB",
				FileMaxAge:     "60d",
				FileMaxBackups: 90,
			},
			WithSource: false,
			WithPID:    false,
		},
		Server: ServerConfig{
			Type:             "grpc",
			GrpcEndpoint:     ":7985",
			HttpEndpoint:     ":7986",
			RequestTimeout:   timex.NewDuration(10 * time.Second),
			MaxCloseWaitTime: timex.NewDuration(time.Minute),
		},
		Database: DatabaseConfig{
			Address:         "127.0.0.1:6033",
			Username:        "postar",
			Password:        "123456",
			MaxOpenConns:    16,
			MaxIdleConns:    4,
			ConnMaxLifetime: timex.NewDuration(3 * time.Minute),
			ConnMaxIdleTime: timex.NewDuration(time.Second),
			ReportStatsTime: timex.NewDuration(time.Minute),
		},
	}

	return conf
}

func (pac *PostarAdminConfig) Check() error {
	return pac.Crypto.Check()
}
