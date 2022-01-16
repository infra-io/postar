// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/22 00:24:19

package biz

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/avino-plan/postar/internal/postard/model"
	"github.com/avino-plan/postar/pkg/concurrency"
)

// go test -v -cover -run=^TestSmtpBiz$
func TestSmtpBiz(t *testing.T) {
	host := os.Getenv("POSTAR_SMTP_HOST")
	user := os.Getenv("POSTAR_SMTP_USER")
	password := os.Getenv("POSTAR_SMTP_PASSWORD")
	to := os.Getenv("POSTAR_SMTP_TO")
	if host == "" || user == "" || password == "" || to == "" {
		t.Skipf("smtp host %s or user %s or password %s or to %s is empty", host, user, password, to)
	}

	port, err := strconv.ParseInt(os.Getenv("POSTAR_SMTP_PORT"), 10, 64)
	if err != nil {
		port = 587
	}

	pool := concurrency.NewPool().Start()
	defer pool.Stop()

	smtpService := NewSmtpService(pool, host, int(port), user, password)
	err = smtpService.SendEmail(context.Background(), &model.Email{
		To:       []string{to},
		Subject:  t.Name(),
		BodyType: "text/html;charset=utf-8",
		Body:     t.Name() + time.Now().Format("20060102150405.000"),
	}, model.DefaultSendEmailOptions())
	if err != nil {
		t.Error(err)
	}
}
