// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"time"

	logitconf "github.com/FishGoddess/logit/extension/config"
)

// PostarConfig stores all configurations of postar.
type PostarConfig struct {
	Logger   logitconf.Config `json:"logger" toml:"logger"`
	Server   ServerConfig     `json:"server" toml:"server"`
	Database DatabaseConfig   `json:"database" toml:"database"`
	Crypto   CryptoConfig     `json:"crypto" toml:"crypto"`
	Cache    CacheConfig      `json:"cache" toml:"cache"`
}

// NewPostarConfig returns a new config for postar.
func NewPostarConfig() *PostarConfig {
	conf := &PostarConfig{
		Logger: logitconf.Config{
			Level:   "debug",
			Handler: "tape",
			Writer: logitconf.WriterConfig{
				Target:         "./log/postar.log",
				FileRotate:     true,
				FileMaxSize:    "128MB",
				FileMaxAge:     "30d",
				FileMaxBackups: 60,
			},
			WithSource: false,
			WithPID:    false,
		},
		Server: ServerConfig{
			Type:               "gateway",
			GrpcEndpoint:       ":5897",
			HttpEndpoint:       ":6897",
			UseTLS:             false,
			CertFile:           "./cert/localhost.crt",
			KeyFile:            "./cert/localhost.key",
			RequestTimeout:     Duration(10 * time.Second),
			CloseServerTimeout: Duration(time.Minute),
		},
		Database: DatabaseConfig{
			Address:         "127.0.0.1:6033",
			Username:        "postar",
			Password:        "123456",
			Database:        "postar",
			MaxOpenConns:    64,
			MaxIdleConns:    16,
			ConnMaxIdleTime: Duration(5 * time.Minute),
			ConnMaxLifetime: Duration(15 * time.Minute),
			ReportStatsTime: Duration(time.Minute),
		},
		Crypto: CryptoConfig{
			AESKey: "",
			AESIV:  "",
		},
		Cache: CacheConfig{
			UseSpaceCache:    true,
			UseAccountCache:  true,
			UseTemplateCache: true,
		},
	}

	return conf
}

func (pc *PostarConfig) Check() error {
	return pc.Crypto.Check()
}
