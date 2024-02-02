// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"context"
	"database/sql"
	stdsql "database/sql"

	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/configs"
	"github.com/infra-io/postar/internal/postar/model"
)

type TemplateStore struct {
	conf *configs.PostarConfig
	db   *sql.DB
}

func NewTemplateStore(conf *configs.PostarConfig) *TemplateStore {
	store := &TemplateStore{
		conf: conf,
		db:   db,
	}

	return store
}

func (ts *TemplateStore) newGetTemplateSQL(spaceID int32, templateID int64) (sql string, args []any) {
	sql = "SELECT `id`, `account_id`, `name`, `email_subject`, `email_to`, `email_cc`, `email_bcc`, `email_content_type`, `email_content`, `state` FROM `templates` WHERE `space_id` = ? AND `id` = ? AND `state` != ?"
	args = []any{spaceID, templateID, model.TemplateStateDeleted}

	return sql, args
}

func (ts *TemplateStore) newTemplate(row *stdsql.Row) (*model.Template, error) {
	template := new(model.Template)
	to, cc, bcc := "", "", ""

	err := row.Scan(&template.ID, &template.AccountID, &template.Name, &template.Email.Subject, &to, &cc, &bcc, &template.Email.ContentType, &template.Email.Content, &template.State)
	if err == stdsql.ErrNoRows {
		return nil, errors.NotFound(err, errors.WithMsg("模板不存在"))
	}

	template.Email.To = decodeStrings(to)
	template.Email.CC = decodeStrings(cc)
	template.Email.BCC = decodeStrings(bcc)

	return template, err
}

func (ts *TemplateStore) GetTemplate(ctx context.Context, spaceID int32, templateID int64) (*model.Template, error) {
	logger := logit.FromContext(ctx)

	sql, args := ts.newGetTemplateSQL(spaceID, templateID)
	logger.Debug("new get template sql", "sql", sql, "args", args)

	row := ts.db.QueryRowContext(ctx, sql, args...)
	if err := row.Err(); err != nil {
		logger.Error("query sql failed", "err", err, "sql", sql, "args", args)
		return nil, err
	}

	return ts.newTemplate(row)
}
