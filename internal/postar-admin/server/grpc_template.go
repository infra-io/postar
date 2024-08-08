// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	postaradminv1 "github.com/infra-io/postar/api/genproto/postaradmin/v1"
	"github.com/infra-io/postar/internal/postar-admin/model"
	"github.com/infra-io/postar/pkg/grpc/contextutil"
)

func newTemplateEmail(email *postaradminv1.TemplateEmail) *model.TemplateEmail {
	if email == nil {
		return new(model.TemplateEmail)
	}

	result := &model.TemplateEmail{
		Subject:     email.Subject,
		To:          email.To,
		CC:          email.Cc,
		BCC:         email.Bcc,
		ContentType: model.EmailContentType(email.ContentType),
		Content:     email.Content,
	}

	return result
}

func newTemplate(template *postaradminv1.Template) *model.Template {
	if template == nil {
		return new(model.Template)
	}

	result := &model.Template{
		ID:          template.Id,
		AccountID:   template.AccountId,
		Name:        template.Name,
		Email:       *newTemplateEmail(template.Email),
		Description: template.Description,
		State:       model.TemplateState(template.State),
		CreateTime:  template.CreateTime,
		UpdateTime:  template.UpdateTime,
	}

	return result
}

func fromTemplateEmail(email *model.TemplateEmail) *postaradminv1.TemplateEmail {
	if email == nil {
		return new(postaradminv1.TemplateEmail)
	}

	result := &postaradminv1.TemplateEmail{
		Subject:     email.Subject,
		To:          email.To,
		Cc:          email.CC,
		Bcc:         email.BCC,
		ContentType: postaradminv1.EmailContentType(email.ContentType),
		Content:     email.Content,
	}

	return result
}

func fromTemplate(template *model.Template) *postaradminv1.Template {
	if template == nil {
		return new(postaradminv1.Template)
	}

	result := &postaradminv1.Template{
		Id:          template.ID,
		AccountId:   template.AccountID,
		Name:        template.Name,
		Email:       fromTemplateEmail(&template.Email),
		Description: template.Description,
		State:       postaradminv1.TemplateState(template.State),
		CreateTime:  template.CreateTime,
		UpdateTime:  template.UpdateTime,
	}

	return result
}

func fromTemplates(templates []*model.Template) []*postaradminv1.Template {
	result := make([]*postaradminv1.Template, 0, len(templates))
	for _, template := range templates {
		result = append(result, fromTemplate(template))
	}

	return result
}

func newListTemplatesFilter(filter *postaradminv1.ListTemplatesFilter) *model.ListTemplatesFilter {
	if filter == nil {
		return new(model.ListTemplatesFilter)
	}

	result := &model.ListTemplatesFilter{
		AccountID:     filter.AccountId,
		TemplateID:    filter.TemplateId,
		TemplateName:  filter.TemplateName,
		TemplateState: model.TemplateState(filter.TemplateState),
		EmailSubject:  filter.EmailSubject,
	}

	return result
}

func newCreateTemplateResponse(Template *model.Template) *postaradminv1.CreateTemplateResponse {
	result := &postaradminv1.CreateTemplateResponse{
		Template: fromTemplate(Template),
	}

	return result
}

func newUpdateTemplateResponse(template *model.Template) *postaradminv1.UpdateTemplateResponse {
	result := &postaradminv1.UpdateTemplateResponse{
		Template: fromTemplate(template),
	}

	return result
}

func newGetTemplateResponse(template *model.Template) *postaradminv1.GetTemplateResponse {
	result := &postaradminv1.GetTemplateResponse{
		Template: fromTemplate(template),
	}

	return result
}

func newListTemplatesResponse(templates []*model.Template, nextPageToken string) *postaradminv1.ListTemplatesResponse {
	result := &postaradminv1.ListTemplatesResponse{
		Templates:     fromTemplates(templates),
		NextPageToken: nextPageToken,
	}

	return result
}

func newDeleteTemplateResponse() *postaradminv1.DeleteTemplateResponse {
	result := &postaradminv1.DeleteTemplateResponse{}

	return result
}

func (gs *GrpcServer) CreateTemplate(ctx context.Context, request *postaradminv1.CreateTemplateRequest) (response *postaradminv1.CreateTemplateResponse, err error) {
	spaceID := contextutil.GetSpaceID(ctx)
	template := newTemplate(request.Template)

	createdTemplate, err := gs.templateService.CreateTemplate(ctx, spaceID, template)
	if err != nil {
		return nil, err
	}

	response = newCreateTemplateResponse(createdTemplate)
	return response, nil
}

func (gs *GrpcServer) UpdateTemplate(ctx context.Context, request *postaradminv1.UpdateTemplateRequest) (response *postaradminv1.UpdateTemplateResponse, err error) {
	spaceID := contextutil.GetSpaceID(ctx)
	template := newTemplate(request.Template)

	updatedTemplate, err := gs.templateService.UpdateTemplate(ctx, spaceID, template)
	if err != nil {
		return nil, err
	}

	response = newUpdateTemplateResponse(updatedTemplate)
	return response, nil
}

func (gs *GrpcServer) GetTemplate(ctx context.Context, request *postaradminv1.GetTemplateRequest) (response *postaradminv1.GetTemplateResponse, err error) {
	spaceID := contextutil.GetSpaceID(ctx)

	template, err := gs.templateService.GetTemplate(ctx, spaceID, request.TemplateId)
	if err != nil {
		return nil, err
	}

	response = newGetTemplateResponse(template)
	return response, nil
}

func (gs *GrpcServer) ListTemplates(ctx context.Context, request *postaradminv1.ListTemplatesRequest) (response *postaradminv1.ListTemplatesResponse, err error) {
	spaceID := contextutil.GetSpaceID(ctx)
	filter := newListTemplatesFilter(request.Filter)

	templates, nextPageToken, err := gs.templateService.ListTemplates(ctx, spaceID, request.PageSize, request.PageToken, filter)
	if err != nil {
		return nil, err
	}

	response = newListTemplatesResponse(templates, nextPageToken)
	return response, nil
}

func (gs *GrpcServer) DeleteTemplate(ctx context.Context, request *postaradminv1.DeleteTemplateRequest) (response *postaradminv1.DeleteTemplateResponse, err error) {
	spaceID := contextutil.GetSpaceID(ctx)

	err = gs.templateService.DeleteTemplate(ctx, spaceID, request.TemplateId)
	if err != nil {
		return nil, err
	}

	response = newDeleteTemplateResponse()
	return response, nil
}
