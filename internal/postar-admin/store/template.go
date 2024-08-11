// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"context"
	"database/sql"
	stdsql "database/sql"
	"strings"
	"time"

	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/config"
	"github.com/infra-io/postar/internal/postar-admin/model"
)

type TemplateStore struct {
	conf *config.PostarAdminConfig
	db   *sql.DB
}

func NewTemplateStore(conf *config.PostarAdminConfig) *TemplateStore {
	store := &TemplateStore{
		conf: conf,
		db:   db,
	}

	return store
}

func (ts *TemplateStore) newCreateTemplateSQL(spaceID int32, template *model.Template) (sql string, args []any) {
	to := encodeStrings(template.Email.To)
	cc := encodeStrings(template.Email.CC)
	bcc := encodeStrings(template.Email.BCC)

	sql = "INSERT INTO `templates`(`space_id`, `account_id`, `name`, `description`, `email_subject`, `email_to`, `email_cc`, `email_bcc`, `email_content_type`, `email_content`, `state`, `create_time`, `update_time`) VALUE(?,?,?,?,?,?,?,?,?,?,?,?,?)"
	args = []any{spaceID, template.AccountID, template.Name, template.Description, template.Email.Subject, to, cc, bcc, template.Email.ContentType, template.Email.Content, template.State, template.CreateTime, template.UpdateTime}

	return sql, args
}

func (ts *TemplateStore) CreateTemplate(ctx context.Context, spaceID int32, template *model.Template) error {
	logger := logit.FromContext(ctx)

	sql, args := ts.newCreateTemplateSQL(spaceID, template)
	logger.Debug("new create template sql", "sql", sql, "args", args)

	stmt, err := ts.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.Error("prepare sql failed", "err", err, "sql", sql, "args", args)
		return err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.Error("exec sql failed", "err", err, "sql", sql, "args", args)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("get ltst insert id failed", "err", err, "template", template)
		err = nil // Ignore this error because template hts been created.
	}

	template.ID = id
	return nil
}

func (ts *TemplateStore) newUpdateTemplateSQL(spaceID int32, template *model.Template) (sql string, args []any) {
	var builder strings.Builder

	builder.WriteString("UPDATE `templates` SET `update_time` = ?, `email_to` = ?, `email_cc` = ?, `email_bcc` = ?")
	args = append(args, template.UpdateTime, encodeStrings(template.Email.To), encodeStrings(template.Email.CC), encodeStrings(template.Email.BCC))

	if template.AccountID > 0 {
		builder.WriteString(", `account_id` = ?")
		args = append(args, template.AccountID)
	}

	if template.Name != "" {
		builder.WriteString(", `name` = ?")
		args = append(args, template.Name)
	}

	if template.Description != "" {
		builder.WriteString(", `description` = ?")
		args = append(args, template.Description)
	}

	if template.Email.Subject != "" {
		builder.WriteString(", `email_subject` = ?")
		args = append(args, template.Email.Subject)
	}

	if template.Email.ContentType > 0 {
		builder.WriteString(", `email_content_type` = ?")
		args = append(args, template.Email.ContentType)
	}

	if template.Email.Content != "" {
		builder.WriteString(", `email_content` = ?")
		args = append(args, template.Email.Content)
	}

	if template.State > 0 {
		builder.WriteString(", `state` = ?")
		args = append(args, template.State)
	}

	builder.WriteString(" WHERE `space_id` = ? AND `id` = ? AND `state` != ?")
	args = append(args, spaceID, template.ID, model.TemplateStateDeleted)

	sql = builder.String()
	return sql, args
}

func (ts *TemplateStore) UpdateTemplate(ctx context.Context, spaceID int32, template *model.Template) error {
	logger := logit.FromContext(ctx)

	sql, args := ts.newUpdateTemplateSQL(spaceID, template)
	logger.Debug("new update template sql", "sql", sql, "args", args)

	stmt, err := ts.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.Error("prepare sql failed", "err", err, "sql", sql, "args", args)
		return err
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, args...); err != nil {
		logger.Error("exec sql failed", "err", err, "sql", sql, "args", args)
		return err
	}

	return nil
}

func (ts *TemplateStore) newGetTemplateSQL(spaceID int32, templateID int64) (sql string, args []any) {
	sql = "SELECT `id`, `account_id`, `name`, `description`, `email_subject`, `email_to`, `email_cc`, `email_bcc`, `email_content_type`, `email_content`, `state`, `create_time`, `update_time` FROM `templates` WHERE `space_id` = ? AND `id` = ? AND `state` != ?"
	args = []any{spaceID, templateID, model.TemplateStateDeleted}

	return sql, args
}

func (ts *TemplateStore) newTemplate(row *stdsql.Row) (*model.Template, error) {
	template := new(model.Template)
	to, cc, bcc := "", "", ""

	err := row.Scan(&template.ID, &template.AccountID, &template.Name, &template.Description, &template.Email.Subject, &to, &cc, &bcc, &template.Email.ContentType, &template.Email.Content, &template.State, &template.CreateTime, &template.UpdateTime)
	if err == stdsql.ErrNoRows {
		return nil, errors.NotFound("模板不存在").With(err)
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

func (ts *TemplateStore) newTemplates(rows *stdsql.Rows) ([]*model.Template, error) {
	templates := make([]*model.Template, 0, 16)

	for rows.Next() {
		template := new(model.Template)
		to, cc, bcc := "", "", ""

		err := rows.Scan(&template.ID, &template.AccountID, &template.Name, &template.Description, &template.Email.Subject, &to, &cc, &bcc, &template.Email.ContentType, &template.Email.Content, &template.State, &template.CreateTime, &template.UpdateTime)
		if err != nil {
			return nil, err
		}

		template.Email.To = decodeStrings(to)
		template.Email.CC = decodeStrings(cc)
		template.Email.BCC = decodeStrings(bcc)

		templates = append(templates, template)
	}

	return templates, nil
}

func (ts *TemplateStore) newListTemplatesSQL(spaceID int32, skip int64, limit int64, filter *model.ListTemplatesFilter) (sql string, args []any) {
	var builder strings.Builder

	builder.WriteString("SELECT `id`, `account_id`, `name`, `description`, `email_subject`, `email_to`, `email_cc`, `email_bcc`, `email_content_type`, `email_content`, `state`, `create_time`, `update_time` FROM `templates` WHERE `space_id` = ? AND `state` != ?")
	args = append(args, spaceID, model.TemplateStateDeleted)

	if filter.AccountID > 0 {
		builder.WriteString(" AND `account_id` = ?")
		args = append(args, filter.AccountID)
	}

	if filter.TemplateID > 0 {
		builder.WriteString(" AND `id` = ?")
		args = append(args, filter.TemplateID)
	}

	if filter.TemplateName != "" {
		builder.WriteString(" AND `name` LIKE '%?%'")
		args = append(args, filter.TemplateName)
	}

	if filter.TemplateState > 0 {
		builder.WriteString(" AND `state` = ?")
		args = append(args, filter.TemplateState)
	}

	if filter.EmailSubject != "" {
		builder.WriteString(" AND `email_subject` LIKE '%?%'")
		args = append(args, filter.EmailSubject)
	}

	builder.WriteString(" ORDER BY `id` DESC LIMIT ?, ?")
	args = append(args, skip, limit)

	sql = builder.String()
	return sql, args
}

func (ts *TemplateStore) ListTemplates(ctx context.Context, spaceID int32, skip int64, limit int64, filter *model.ListTemplatesFilter) ([]*model.Template, error) {
	logger := logit.FromContext(ctx)

	sql, args := ts.newListTemplatesSQL(spaceID, skip, limit, filter)
	logger.Debug("new list templates sql", "sql", sql, "args", args)

	rows, err := ts.db.QueryContext(ctx, sql, args...)
	if err != nil {
		logger.Error("query sql failed", "err", err, "sql", sql, "args", args)
		return nil, err
	}

	defer rows.Close()

	return ts.newTemplates(rows)
}

func (ts *TemplateStore) newDeleteTemplateSQL(spaceID int32, templateID int64) (sql string, args []any) {
	sql = "UPDATE `templates` SET `state` = ? AND `update_time` = ? WHERE `space_id` = ? AND `id` = ?"
	args = []any{model.TemplateStateDeleted, time.Now().Unix(), spaceID, templateID}

	return sql, args
}

func (ts *TemplateStore) DeleteTemplate(ctx context.Context, spaceID int32, templateID int64) error {
	logger := logit.FromContext(ctx)

	sql, args := ts.newDeleteTemplateSQL(spaceID, templateID)
	logger.Debug("new delete template sql", "sql", sql, "args", args)

	stmt, err := ts.db.PrepareContext(ctx, sql)
	if err != nil {
		logger.Error("prepare sql failed", "err", err, "sql", sql, "args", args)
		return err
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, args...); err != nil {
		logger.Error("exec sql failed", "err", err, "sql", sql, "args", args)
		return err
	}

	return nil
}
