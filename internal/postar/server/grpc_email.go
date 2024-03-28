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

func fromEmail(email *model.Email) *postarv1.Email {
	if email == nil {
		return new(postarv1.Email)
	}

	result := &postarv1.Email{
		TemplateId:    email.TemplateID,
		To:            email.To,
		Cc:            email.CC,
		Bcc:           email.BCC,
		SubjectParams: email.SubjectParams,
		ContentParams: email.ContentParams,
	}

	return result
}

func parseSendEmailRequest(request *postarv1.SendEmailRequest) *model.Email {
	return newEmail(request.Email)
}

func newSendEmailResponse() *postarv1.SendEmailResponse {
	return new(postarv1.SendEmailResponse)
}

func (gs *GrpcServer) SendEmail(ctx context.Context, request *postarv1.SendEmailRequest) (response *postarv1.SendEmailResponse, err error) {
	email := parseSendEmailRequest(request)

	if err = gs.emailService.SendEmail(ctx, email); err != nil {
		return nil, err
	}

	response = newSendEmailResponse()
	return response, nil
}
