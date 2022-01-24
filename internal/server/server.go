// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/01/23 02:18:24

package server

import (
	"fmt"
	"time"

	"github.com/FishGoddess/logit"
	"github.com/avinoplan/postar/api"
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/biz"
	"github.com/avinoplan/postar/internal/model"
)

var (
	servers = map[string]func(c *configs.Config, logger *logit.Logger, smtpBiz *biz.SMTPBiz) Server{
		"grpc": NewGRPCServer,
	}
)

type Server interface {
	Start() error
	Stop() error
}

func NewServer(c *configs.Config, logger *logit.Logger, smtpBiz *biz.SMTPBiz) Server {
	newServer, ok := servers[c.ServerType()]
	if !ok {
		panic(fmt.Errorf("server: type %s not found", c.ServerType()))
	}
	return newServer(c, logger, smtpBiz)
}

func toModelEmail(email *api.Email) *model.Email {
	if email == nil {
		return nil
	}

	result := model.NewEmail()
	result.Subject = email.Subject
	result.Receivers = email.Receivers
	result.BodyType = email.BodyType
	result.Body = email.Body
	return result
}

func toModelSendEmailOptions(c *configs.Config, opts *api.SendEmailOptions) *model.SendEmailOptions {
	if opts == nil {
		return nil
	}

	result := model.DefaultSendEmailOptions(c)
	result.Async = opts.Async
	if opts.Timeout > 0 {
		result.Timeout = time.Duration(opts.Timeout) * time.Millisecond
	}
	return result
}
