// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package configs

import (
	"fmt"

	timex "github.com/infra-io/servicex/time"
)

const (
	aesKeyLength = 24
	aesIVLength  = 16
)

type ServerConfig struct {
	Type             string         `json:"type" toml:"type"`
	GrpcEndpoint     string         `json:"grpc_endpoint" toml:"grpc_endpoint"`
	HttpEndpoint     string         `json:"http_endpoint" toml:"http_endpoint"`
	RequestTimeout   timex.Duration `json:"request_timeout" toml:"request_timeout"`
	MaxCloseWaitTime timex.Duration `json:"max_close_wait_time" toml:"max_close_wait_time"`
}

type DatabaseConfig struct {
	Address         string         `json:"address" toml:"address"`
	Username        string         `json:"username" toml:"username"`
	Password        string         `json:"-" toml:"password"`
	MaxIdleConns    int            `json:"max_idle_conns" toml:"max_idle_conns"`
	MaxOpenConns    int            `json:"max_open_conns" toml:"max_open_conns"`
	ConnMaxLifetime timex.Duration `json:"conn_max_lifetime" toml:"conn_max_lifetime"`
	ConnMaxIdleTime timex.Duration `json:"conn_max_idle_time" toml:"conn_max_idle_time"`
	ReportStatsTime timex.Duration `json:"report_stats_time" toml:"report_stats_time"`
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
