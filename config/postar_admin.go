// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"time"

	logitconf "github.com/FishGoddess/logit/extension/config"
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
			Level:   "debug",
			Handler: "tape",
			Writer: logitconf.WriterConfig{
				Target:         "./log/postar_admin.log",
				FileRotate:     true,
				FileMaxSize:    "128MB",
				FileMaxAge:     "60d",
				FileMaxBackups: 90,
			},
			WithSource: false,
			WithPID:    false,
		},
		Server: ServerConfig{
			Type:               "gateway",
			GrpcEndpoint:       ":7985",
			HttpEndpoint:       ":7986",
			UseTLS:             false,
			CertFile:           "./cert/localhost.crt",
			KeyFile:            "./cert/localhost.key",
			RequestTimeout:     Duration(time.Second),
			CloseServerTimeout: Duration(time.Minute),
		},
		Database: DatabaseConfig{
			Address:         "127.0.0.1:6033",
			Username:        "postar",
			Password:        "123456",
			Database:        "postar",
			MaxOpenConns:    16,
			MaxIdleConns:    4,
			ConnMaxIdleTime: Duration(3 * time.Minute),
			ConnMaxLifetime: Duration(10 * time.Minute),
			ReportStatsTime: Duration(time.Minute),
		},
		Crypto: CryptoConfig{
			AESKey: "",
			AESIV:  "",
		},
	}

	return conf
}

func (pac *PostarAdminConfig) Check() error {
	return pac.Crypto.Check()
}
