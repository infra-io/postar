// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"fmt"
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
		err := errors.New("template.AccountID <= 0")
		return errors.BadRequest(err, errors.WithMsg("模板绑定的账号编号非法"))
	}

	if strings.TrimSpace(template.Name) == "" {
		err := errors.New("trim template.Name == ''")
		return errors.BadRequest(err, errors.WithMsg("模板名称不能为空"))
	}

	if strings.TrimSpace(template.Email.Subject) == "" {
		err := errors.New("trim template.Email.Subject == ''")
		return errors.BadRequest(err, errors.WithMsg("模板邮件主题不能为空"))
	}

	if !template.Email.ContentType.Valid() {
		err := fmt.Errorf("template.Email.ContentType %d not valid", template.Email.ContentType)
		return errors.BadRequest(err, errors.WithMsg("模板邮件内容类型无效"))
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
		err := errors.New("template.ID <= 0")
		return errors.BadRequest(err, errors.WithMsg("模板编号非法"))
	}

	if template.AccountID <= 0 {
		err := errors.New("template.AccountID <= 0")
		return errors.BadRequest(err, errors.WithMsg("模板绑定的账号编号非法"))
	}

	if strings.TrimSpace(template.Name) == "" {
		err := errors.New("trim template.Name == ''")
		return errors.BadRequest(err, errors.WithMsg("模板名称不能为空"))
	}

	if strings.TrimSpace(template.Email.Subject) == "" {
		err := errors.New("trim template.Email.Subject == ''")
		return errors.BadRequest(err, errors.WithMsg("模板邮件主题不能为空"))
	}

	if template.Email.ContentType > 0 && !template.Email.ContentType.Valid() {
		err := fmt.Errorf("template.Email.Content.Type %d not valid", template.Email.ContentType)
		return errors.BadRequest(err, errors.WithMsg("模板邮件内容类型无效"))
	}

	if template.State > 0 && !template.State.Valid() {
		err := fmt.Errorf("template.State %d not valid", template.State)
		return errors.BadRequest(err, errors.WithMsg("模板状态无效"))
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
		err := fmt.Errorf("templateID %d <= 0", templateID)
		return errors.BadRequest(err, errors.WithMsg("模板编号非法"))
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
		err := fmt.Errorf("pageSize %d not in [%d, %d]", pageSize, minPageSize, maxPageSize)
		return errors.BadRequest(err, errors.WithMsg("分页大小需要位于区间 [%d, %d] 内", minPageSize, maxPageSize))
	}

	if filter.AccountID < 0 {
		err := fmt.Errorf("filter.AccountID %d < 0", filter.AccountID)
		return errors.BadRequest(err, errors.WithMsg("过滤的账号编号非法"))
	}

	if filter.TemplateID < 0 {
		err := fmt.Errorf("filter.TemplateID %d < 0", filter.TemplateID)
		return errors.BadRequest(err, errors.WithMsg("过滤的模板编号非法"))
	}

	if filter.TemplateState > 0 && !filter.TemplateState.Valid() {
		err := fmt.Errorf("filter.TemplateState %d not valid", filter.TemplateState)
		return errors.BadRequest(err, errors.WithMsg("过滤的模板状态非法"))
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
		err := fmt.Errorf("templateID %d <= 0", templateID)
		return errors.BadRequest(err, errors.WithMsg("模板编号非法"))
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
