// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	postarv1 "github.com/infra-io/postar/api/genproto/postar/v1"
	"github.com/infra-io/postar/internal/postar/model"
)

func newEmail(email *postarv1.Email) *model.Email {
	if email == nil {
		return new(model.Email)
	}

	result := &model.Email{
		TemplateID:    email.TemplateId,
		To:            email.To,
		CC:            email.Cc,
		BCC:           email.Bcc,
		SubjectParams: email.SubjectParams,
		ContentParams: email.ContentParams,
	}

	return result
}

func newSendEmailOptions(conf *postarv1.SendEmailOptions) *model.SendEmailOptions {
	if conf == nil {
		return new(model.SendEmailOptions)
	}

	result := &model.SendEmailOptions{
		Async: conf.Async,
	}

	return result
}

func parseSendEmailRequest(request *postarv1.SendEmailRequest) (email *model.Email, options *model.SendEmailOptions) {
	email = newEmail(request.Email)
	options = newSendEmailOptions(request.Options)

	return email, options
}

func newSendEmailResponse() *postarv1.SendEmailResponse {
	resp := &postarv1.SendEmailResponse{}

	return resp
}

func (gs *GrpcServer) SendEmail(ctx context.Context, request *postarv1.SendEmailRequest) (response *postarv1.SendEmailResponse, err error) {
	email, options := parseSendEmailRequest(request)

	if err = gs.emailService.SendEmail(ctx, email, options); err != nil {
		return nil, err
	}

	response = newSendEmailResponse()
	return response, nil
}
