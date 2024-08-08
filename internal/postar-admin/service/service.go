// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"strconv"

	"github.com/infra-io/postar/internal/postar-admin/model"
)

const (
	minPageSize = 1
	maxPageSize = 200
)

type SpaceService interface {
	CreateSpace(ctx context.Context, space *model.Space) (*model.Space, error)
	UpdateSpace(ctx context.Context, space *model.Space) (*model.Space, error)
	GetSpace(ctx context.Context, spaceID int32, withToken bool) (*model.Space, error)
	ListSpaces(ctx context.Context, pageSize int32, pageToken string, filter *model.ListSpacesFilter) ([]*model.Space, string, error)
}

type AccountService interface {
	CreateAccount(ctx context.Context, spaceID int32, account *model.Account) (*model.Account, error)
	UpdateAccount(ctx context.Context, spaceID int32, account *model.Account) (*model.Account, error)
	GetAccount(ctx context.Context, spaceID int32, accountID int32, withPassword bool) (*model.Account, error)
	ListAccounts(ctx context.Context, spaceID int32, pageSize int32, pageToken string, filter *model.ListAccountsFilter) ([]*model.Account, string, error)
}

type TemplateService interface {
	CreateTemplate(ctx context.Context, spaceID int32, template *model.Template) (*model.Template, error)
	UpdateTemplate(ctx context.Context, spaceID int32, template *model.Template) (*model.Template, error)
	GetTemplate(ctx context.Context, spaceID int32, templateID int64) (*model.Template, error)
	ListTemplates(ctx context.Context, spaceID int32, pageSize int32, pageToken string, filter *model.ListTemplatesFilter) ([]*model.Template, string, error)
	DeleteTemplate(ctx context.Context, spaceID int32, templateID int64) error
}

func parsePageToken(pageToken string) (skip int64, err error) {
	if pageToken == "" {
		return 0, nil
	}

	return strconv.ParseInt(pageToken, 10, 64)
}

func newNextPageToken(skip int64, pageSize int32, currentSize int) string {
	if currentSize < int(pageSize) {
		return ""
	}

	skip = skip + int64(pageSize)
	return strconv.FormatInt(skip, 10)
}
