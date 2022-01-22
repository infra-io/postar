// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/01/23 02:18:24

package server

import (
	"github.com/FishGoddess/logit"
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/biz"
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
