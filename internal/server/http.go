// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"io"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/FishGoddess/errors"
	baseapi "github.com/avino-plan/api/go-out/base"
	postarapi "github.com/avino-plan/api/go-out/postar"
	"github.com/avino-plan/postar/configs"
	"github.com/avino-plan/postar/internal/biz"
	"github.com/avino-plan/postar/pkg/trace"
	"github.com/go-logit/logit"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/protobuf/proto"
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

func (hs *HTTPServer) unmarshalSendEmailRequest(reader io.Reader) (*postarapi.SendEmailRequest, error) {
	marshaled, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	request := new(postarapi.SendEmailRequest)
	err = proto.Unmarshal(marshaled, request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (hs *HTTPServer) marshalSendEmailResponse(response *postarapi.SendEmailResponse) []byte {
	marshaled, err := proto.Marshal(response)
	if err != nil {
		logit.Error("proto.Marshal(response) failed").Error("err", err).Stringer("response", response).End()
		return nil // should never happen...
	}
	return marshaled
}

func (hs *HTTPServer) writeSendEmailResponse(writer http.ResponseWriter, statusCode int, response *postarapi.SendEmailResponse) {
	writer.WriteHeader(statusCode)
	_, err := writer.Write(hs.marshalSendEmailResponse(response))
	if err != nil {
		logit.Error("write send email response failed").Error("err", err).Int("statusCode", statusCode).Stringer("response", response).End()
	}
}

func (hs *HTTPServer) sendEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	traceID := trace.NewTraceID()
	ctx := trace.NewContext(request.Context(), traceID)

	req, err := hs.unmarshalSendEmailRequest(request.Body)
	if err != nil {
		hs.writeSendEmailResponse(writer, http.StatusBadRequest, &postarapi.SendEmailResponse{
			Code:    baseapi.ServerCode_BAD_REQUEST,
			Msg:     "unmarshal send email request failed",
			TraceId: traceID,
		})
		return
	}

	err = hs.smtpBiz.SendEmail(ctx, toModelEmail(req.Email), toModelSendEmailOptions(hs.c, req.Options))
	if errors.IsBadRequest(err) {
		hs.writeSendEmailResponse(writer, http.StatusBadRequest, &postarapi.SendEmailResponse{
			Code:    baseapi.ServerCode_BAD_REQUEST,
			Msg:     err.Error(),
			TraceId: traceID,
		})
		return
	}

	if errors.IsTimeout(err) {
		hs.writeSendEmailResponse(writer, http.StatusRequestTimeout, &postarapi.SendEmailResponse{
			Code:    baseapi.ServerCode_TIMEOUT,
			Msg:     "send email timeout",
			TraceId: traceID,
		})
		return
	}

	if err != nil {
		hs.writeSendEmailResponse(writer, http.StatusInternalServerError, &postarapi.SendEmailResponse{
			Code:    baseapi.ServerCode_SEND_EMAIL_FAILED,
			Msg:     "send email failed",
			TraceId: traceID,
		})
		return
	}

	hs.writeSendEmailResponse(writer, http.StatusOK, &postarapi.SendEmailResponse{
		Code:    baseapi.ServerCode_OK,
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
