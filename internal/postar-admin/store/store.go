// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/FishGoddess/logit"
	_ "github.com/go-sql-driver/mysql"
	"github.com/infra-io/postar/configs"
)

const (
	strConnector = "|"
)

var db *sql.DB

func newDSN(conf *configs.DatabaseConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&time_zone=%s",
		conf.Username, conf.Password, conf.Address, conf.Database, url.QueryEscape("'Asia/Shanghai'"),
	)

	return dsn
}

func reportStats(conf *configs.DatabaseConfig) {
	logger := logit.Default().WithGroup("db").With("address", conf.Address, "username", conf.Username, "database", conf.Database)

	ticker := time.NewTicker(conf.ReportStatsTime.Standard())
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			stats := db.Stats()
			logger.Info("report db stats", "stats", stats)
		}
	}
}

func Connect(conf *configs.DatabaseConfig) (err error) {
	db, err = sql.Open("mysql", newDSN(conf))
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetConnMaxLifetime(conf.ConnMaxLifetime.Standard())
	db.SetConnMaxIdleTime(conf.ConnMaxIdleTime.Standard())

	go reportStats(conf)
	return nil
}

func encodeStrings(strs []string) string {
	if marshaled, err := json.Marshal(strs); err == nil {
		return string(marshaled)
	}

	return strings.Join(strs, strConnector)
}

func decodeStrings(str string) []string {
	strs := make([]string, 0, 4)
	if err := json.Unmarshal([]byte(str), &strs); err == nil {
		return strs
	}

	return strings.Split(str, strConnector)
}
