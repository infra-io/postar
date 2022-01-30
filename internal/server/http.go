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
	liberrors "github.com/FishGoddess/errors"
	"github.com/avinoplan/postar/api"
	"github.com/avinoplan/postar/configs"
	"github.com/avinoplan/postar/internal/biz"
	"github.com/avinoplan/postar/pkg/errors"
	"github.com/avinoplan/postar/pkg/log"
	"github.com/avinoplan/postar/pkg/trace"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/protobuf/proto"
	"io"
	"io/ioutil"
	"net"
	"net/http"
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
		server: &http.Server{
			Addr: c.ServerAddress(),
		},
	}

	router := httprouter.New()
	router.POST("/sendEmail", hs.sendEmail)
	hs.server.Handler = router
	return hs
}

func (hs *HTTPServer) unmarshalSendEmailRequest(reader io.Reader) (*api.SendEmailRequest, error) {
	marshaled, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	request := new(api.SendEmailRequest)
	err = proto.Unmarshal(marshaled, request)
	if err != nil {
		return nil, err
	}

	if request.Email == nil {
		return nil, errors.BadRequestErr(liberrors.New("request.Email == nil"))
	}
	return request, nil
}

func (hs *HTTPServer) marshalSendEmailResponse(response *api.SendEmailResponse) []byte {
	marshaled, err := proto.Marshal(response)
	if err != nil {
		log.Error(err, "proto.Marshal(response) failed").Stringer("response", response).End()
		return nil // should never happen...
	}
	return marshaled
}

func (hs *HTTPServer) writeSendEmailResponse(writer http.ResponseWriter, statusCode int, response *api.SendEmailResponse) {
	writer.WriteHeader(statusCode)
	_, err := writer.Write(hs.marshalSendEmailResponse(response))
	if err != nil {
		log.Error(err, "write send email response failed").Int("statusCode", statusCode).Stringer("response", response).End()
	}
}

func (hs *HTTPServer) sendEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	traceID := trace.NewTraceID()
	ctx := trace.NewContext(request.Context(), traceID)

	req, err := hs.unmarshalSendEmailRequest(request.Body)
	if errors.IsBadRequest(err) {
		hs.writeSendEmailResponse(writer, http.StatusBadRequest, &api.SendEmailResponse{
			Code:    api.ServerCode_BAD_REQUEST,
			Msg:     err.Error(),
			TraceId: traceID,
		})
		return
	}

	if err != nil {
		hs.writeSendEmailResponse(writer, http.StatusBadRequest, &api.SendEmailResponse{
			Code:    api.ServerCode_BAD_REQUEST,
			Msg:     "unmarshal send email request failed",
			TraceId: traceID,
		})
		return
	}

	err = hs.smtpBiz.SendEmail(ctx, toModelEmail(req.Email), toModelSendEmailOptions(hs.c, req.Options))
	if errors.IsTimeout(err) {
		hs.writeSendEmailResponse(writer, http.StatusRequestTimeout, &api.SendEmailResponse{
			Code:    api.ServerCode_TIMEOUT,
			Msg:     "send email timeout",
			TraceId: traceID,
		})
		return
	}

	if err != nil {
		hs.writeSendEmailResponse(writer, http.StatusInternalServerError, &api.SendEmailResponse{
			Code:    api.ServerCode_SEND_EMAIL_FAILED,
			Msg:     "send email failed",
			TraceId: traceID,
		})
		return
	}

	hs.writeSendEmailResponse(writer, http.StatusOK, &api.SendEmailResponse{
		Code:    api.ServerCode_OK,
		TraceId: traceID,
	})
}

// Start starts HTTPServer.
func (hs *HTTPServer) Start() error {
	listener, err := net.Listen("tcp", hs.c.ServerAddress())
	if err != nil {
		return err
	}
	defer listener.Close()

	err = hs.server.Serve(listener)
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Stop stops HTTPServer gracefully.
func (hs *HTTPServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), hs.c.ServerStopTimeout())
	defer cancel()
	return hs.server.Shutdown(ctx)
}
