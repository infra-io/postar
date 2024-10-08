// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
)

const (
	aesKeyLength = 24
	aesIVLength  = 16
)

type ServerConfig struct {
	Type               string   `json:"type" toml:"type"`
	GrpcEndpoint       string   `json:"grpc_endpoint" toml:"grpc_endpoint"`
	HttpEndpoint       string   `json:"http_endpoint" toml:"http_endpoint"`
	UseTLS             bool     `json:"use_tls" toml:"use_tls"`
	CertFile           string   `json:"cert_file" toml:"cert_file"`
	KeyFile            string   `json:"key_file" toml:"key_file"`
	RequestTimeout     Duration `json:"request_timeout" toml:"request_timeout"`
	CloseServerTimeout Duration `json:"close_server_timeout" toml:"close_server_timeout"`
}

type DatabaseConfig struct {
	Address         string   `json:"address" toml:"address"`
	Username        string   `json:"username" toml:"username"`
	Password        string   `json:"-" toml:"password"`
	Database        string   `json:"database" toml:"database"`
	MaxIdleConns    int      `json:"max_idle_conns" toml:"max_idle_conns"`
	MaxOpenConns    int      `json:"max_open_conns" toml:"max_open_conns"`
	ConnMaxLifetime Duration `json:"conn_max_lifetime" toml:"conn_max_lifetime"`
	ConnMaxIdleTime Duration `json:"conn_max_idle_time" toml:"conn_max_idle_time"`
	ReportStatsTime Duration `json:"report_stats_time" toml:"report_stats_time"`
}

type CryptoConfig struct {
	AESKey string `json:"-" toml:"aes_key"`
	AESIV  string `json:"-" toml:"aes_iv"`
}

func (cc *CryptoConfig) Check() error {
	if len(cc.AESKey) != aesKeyLength {
		return fmt.Errorf("the length of aes key must be %d", aesKeyLength)
	}

	if len(cc.AESIV) != aesIVLength {
		return fmt.Errorf("the length of aes iv must be %d", aesIVLength)
	}

	return nil
}

type CacheConfig struct {
	UseSpaceCache    bool `json:"use_space_cache" toml:"use_space_cache"`
	UseAccountCache  bool `json:"use_account_cache" toml:"use_account_cache"`
	UseTemplateCache bool `json:"use_template_cache" toml:"use_template_cache"`
}
