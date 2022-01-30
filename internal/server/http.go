// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 02:04:54

package server

import (
	"context"
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/biz"
	"net"
	"net/http"
	"time"
)

type HTTPServer struct {
	c       *configs.Config
	smtpBiz *biz.SMTPBiz
	server  *http.Server
}

func NewHTTPServer(c *configs.Config, smtpBiz *biz.SMTPBiz) Server {
	hs := &HTTPServer{
		c:       c,
		smtpBiz: smtpBiz,
	}

	hs.server = &http.Server{
		Addr:    c.ServerAddress(),
		Handler: hs,
	}
	return hs
}

func (hs *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Hi, I am postar!"))
}

// Start starts HTTPServer.
func (hs *HTTPServer) Start() error {
	listener, err := net.Listen("tcp", hs.c.ServerAddress())
	if err != nil {
		return err
	}
	defer listener.Close()
	return hs.server.Serve(listener)
}

// Stop stops HTTPServer gracefully.
func (hs *HTTPServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return hs.server.Shutdown(ctx)
}
