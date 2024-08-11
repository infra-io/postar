// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"strings"
	"time"

	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/config"
	"github.com/infra-io/postar/internal/postar-admin/model"
)

type TemplateStore interface {
	CreateTemplate(ctx context.Context, spaceID int32, template *model.Template) error
	UpdateTemplate(ctx context.Context, spaceID int32, template *model.Template) error
	GetTemplate(ctx context.Context, spaceID int32, templateID int64) (*model.Template, error)
	ListTemplates(ctx context.Context, spaceID int32, skip int64, limit int64, filter *model.ListTemplatesFilter) ([]*model.Template, error)
	DeleteTemplate(ctx context.Context, spaceID int32, templateID int64) error
}

type defaultTemplateService struct {
	conf          *config.PostarAdminConfig
	templateStore TemplateStore
}

func NewTemplateService(conf *config.PostarAdminConfig, templateStore TemplateStore) TemplateService {
	service := &defaultTemplateService{
		conf:          conf,
		templateStore: templateStore,
	}

	return service
}

func (dts *defaultTemplateService) checkCreateTemplateParams(template *model.Template) error {
	if template.AccountID <= 0 {
		return errors.BadRequest("模板绑定的账号编号需要大于 0")
	}

	if strings.TrimSpace(template.Name) == "" {
		return errors.BadRequest("模板名称不能为空")
	}

	if strings.TrimSpace(template.Email.Subject) == "" {
		return errors.BadRequest("模板邮件主题不能为空")
	}

	if !template.Email.ContentType.Valid() {
		return errors.BadRequest("模板邮件内容类型 %d 无效", template.Email.ContentType)
	}

	return nil
}

func (dts *defaultTemplateService) CreateTemplate(ctx context.Context, spaceID int32, template *model.Template) (*model.Template, error) {
	logger := logit.FromContext(ctx)

	if err := dts.checkCreateTemplateParams(template); err != nil {
		logger.Error("check create template params failed", "err", err, "template", template)
		return nil, err
	}

	now := time.Now().Unix()
	template.State = model.TemplateStateEnabled
	template.CreateTime = now
	template.UpdateTime = now

	err := dts.templateStore.CreateTemplate(ctx, spaceID, template)
	return template, err
}

func (dts *defaultTemplateService) checkUpdateTemplateParams(template *model.Template) error {
	if template.ID <= 0 {
		return errors.BadRequest("模板编号需要大于 0")
	}

	if template.AccountID <= 0 {
		return errors.BadRequest("模板绑定的账号编号需要大于 0")
	}

	if strings.TrimSpace(template.Name) == "" {
		return errors.BadRequest("模板名称不能为空")
	}

	if strings.TrimSpace(template.Email.Subject) == "" {
		return errors.BadRequest("模板邮件主题不能为空")
	}

	if template.Email.ContentType > 0 && !template.Email.ContentType.Valid() {
		return errors.BadRequest("模板邮件内容类型 %d 无效", template.Email.ContentType)
	}

	if template.State > 0 && !template.State.Valid() {
		return errors.BadRequest("模板状态 %d 无效", template.State)
	}

	return nil
}

func (dts *defaultTemplateService) UpdateTemplate(ctx context.Context, spaceID int32, template *model.Template) (*model.Template, error) {
	logger := logit.FromContext(ctx)

	if err := dts.checkUpdateTemplateParams(template); err != nil {
		logger.Error("check update template params failed", "err", err, "template", template)
		return nil, err
	}

	now := time.Now().Unix()
	template.UpdateTime = now

	err := dts.templateStore.UpdateTemplate(ctx, spaceID, template)
	return template, err
}

func (dts *defaultTemplateService) checkGetTemplateParams(templateID int64) error {
	if templateID <= 0 {
		return errors.BadRequest("模板编号需要大于 0")
	}

	return nil
}

func (dts *defaultTemplateService) GetTemplate(ctx context.Context, spaceID int32, templateID int64) (*model.Template, error) {
	logger := logit.FromContext(ctx)

	if err := dts.checkGetTemplateParams(templateID); err != nil {
		logger.Error("check get template params failed", "err", err, "template_id", templateID)
		return nil, err
	}

	return dts.templateStore.GetTemplate(ctx, spaceID, templateID)
}

func (dts *defaultTemplateService) checkListTemplatesParams(pageSize int32, filter *model.ListTemplatesFilter) error {
	if pageSize < minPageSize || pageSize > maxPageSize {
		return errors.BadRequest("分页大小 %d 需要位于区间 [%d, %d] 内", pageSize, minPageSize, maxPageSize)
	}

	if filter.AccountID < 0 {
		return errors.BadRequest("账号编号不能为负数")
	}

	if filter.TemplateID < 0 {
		return errors.BadRequest("模板编号不能为负数")
	}

	if filter.TemplateState > 0 && !filter.TemplateState.Valid() {
		return errors.BadRequest("模板状态 %d 无效", filter.TemplateState)
	}

	return nil
}

func (dts *defaultTemplateService) ListTemplates(ctx context.Context, spaceID int32, pageSize int32, pageToken string, filter *model.ListTemplatesFilter) ([]*model.Template, string, error) {
	logger := logit.FromContext(ctx)

	if err := dts.checkListTemplatesParams(pageSize, filter); err != nil {
		logger.Error("check list accounts params failed", "err", err, "page_size", pageSize, "page_token", pageToken, "filter", filter)
		return nil, "", err
	}

	skip, err := parsePageToken(pageToken)
	if err != nil {
		logger.Error("parse page token failed", "err", err, "page_size", pageSize, "page_token", pageToken, "filter", filter)
		return nil, "", err
	}

	templates, err := dts.templateStore.ListTemplates(ctx, spaceID, skip, int64(pageSize), filter)
	if err != nil {
		logger.Error("list templates failed", "err", err, "skip", skip, "page_size", pageSize, "page_token", pageToken, "filter", filter)
		return nil, "", err
	}

	nextPageToken := newNextPageToken(skip, pageSize, len(templates))
	return templates, nextPageToken, nil
}

func (dts *defaultTemplateService) checkDeleteTemplateParams(templateID int64) error {
	if templateID <= 0 {
		return errors.BadRequest("模板编号需要大于 0")
	}

	return nil
}

func (dts *defaultTemplateService) DeleteTemplate(ctx context.Context, spaceID int32, templateID int64) error {
	logger := logit.FromContext(ctx)

	if err := dts.checkDeleteTemplateParams(templateID); err != nil {
		logger.Error("check delete template params failed", "err", err, "template_id", templateID)
		return err
	}

	return dts.templateStore.DeleteTemplate(ctx, spaceID, templateID)
}
