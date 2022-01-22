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

	"github.com/FishGoddess/logit"
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/model"
	"github.com/panjf2000/ants/v2"
)

func newConfig() *configs.Config {
	port, err := strconv.ParseInt(os.Getenv("POSTAR_SMTP_PORT"), 10, 64)
	if err != nil {
		port = 587
	}

	c := configs.NewDefaultConfig()
	c.SMTP.Host = os.Getenv("POSTAR_SMTP_HOST")
	c.SMTP.Port = int(port)
	c.SMTP.User = os.Getenv("POSTAR_SMTP_USER")
	c.SMTP.Password = os.Getenv("POSTAR_SMTP_PASSWORD")
	return c
}

// go test -v -cover -run=^TestSMTPBiz$
func TestSMTPBiz(t *testing.T) {
	c := newConfig()
	receiver := os.Getenv("POSTAR_SMTP_RECEIVER")
	if c.SMTP.Host == "" || c.SMTP.User == "" || c.SMTP.Password == "" || receiver == "" {
		t.Skipf("smtp host %s or user %s or password %s or receiver %s is empty", c.SMTP.Host, c.SMTP.User, c.SMTP.Password, receiver)
	}

	pool, _ := ants.NewPool(64)
	defer pool.Release()
	
	smtpService := NewSMTPBiz(c, logit.NewLogger(), pool)
	err := smtpService.SendEmail(context.Background(), &model.Email{
		Subject:   t.Name(),
		Receivers: []string{receiver},
		BodyType:  "text/html;charset=utf-8",
		Body:      t.Name() + time.Now().Format("20060102150405.000"),
	}, model.DefaultSendEmailOptions())

	if err != nil {
		t.Error(err)
	}
}
