// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"context"
	stdsql "database/sql"
	"strings"

	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/config"
	"github.com/infra-io/postar/internal/postar-admin/model"
)

type SpaceStore struct {
	conf *config.PostarAdminConfig
	db   *stdsql.DB
}

func NewSpaceStore(conf *config.PostarAdminConfig) *SpaceStore {
	store := &SpaceStore{
		conf: conf,
		db:   db,
	}

	return store
}

func (ss *SpaceStore) newCreateSpaceSQL(space *model.Space) (sql string, args []any) {
	sql = "INSERT INTO `spaces`(`name`, `token`, `state`, `create_time`, `update_time`) VALUE(?,?,?,?,?)"
	args = []any{space.Name, space.Token, space.State, space.CreateTime, space.UpdateTime}

	return sql, args
}

func (ss *SpaceStore) CreateSpace(ctx context.Context, space *model.Space) error {
	logger := logit.FromContext(ctx)

	sql, args := ss.newCreateSpaceSQL(space)
	logger.Debug("new create space sql", "sql", sql, "args", args)

	stmt, err := ss.db.PrepareContext(ctx, sql)
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
		logger.Error("get last insert id failed", "err", err)
		err = nil // Ignore this error because space has been created.
	}

	space.ID = int32(id)
	return nil
}

func (ss *SpaceStore) newUpdateSpaceSQL(space *model.Space) (sql string, args []any) {
	var builder strings.Builder

	builder.WriteString("UPDATE `spaces` SET `update_time` = ?")
	args = append(args, space.UpdateTime)

	if space.Name != "" {
		builder.WriteString(", `name` = ?")
		args = append(args, space.Name)
	}

	if space.State > 0 {
		builder.WriteString(", `state` = ?")
		args = append(args, space.State)
	}

	builder.WriteString(" WHERE `id` = ?")

	sql = builder.String()
	args = append(args, space.ID)
	return sql, args
}

func (ss *SpaceStore) UpdateSpace(ctx context.Context, space *model.Space) error {
	logger := logit.FromContext(ctx)

	sql, args := ss.newUpdateSpaceSQL(space)
	logger.Debug("new update space sql", "sql", sql, "args", args)

	stmt, err := ss.db.PrepareContext(ctx, sql)
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

func (ss *SpaceStore) newGetSpaceSQL(spaceID int32) (sql string, args []any) {
	sql = "SELECT `id`, `name`, `token`, `state`, `create_time`, `update_time` FROM `spaces` WHERE `id` = ?"
	args = []any{spaceID}

	return sql, args
}

func (ss *SpaceStore) newSpace(row *stdsql.Row) (*model.Space, error) {
	space := new(model.Space)

	err := row.Scan(&space.ID, &space.Name, &space.Token, &space.State, &space.CreateTime, &space.UpdateTime)
	if err == stdsql.ErrNoRows {
		return nil, errors.NotFound(err, errors.WithMsg("业务空间不存在"))
	}

	return space, err
}

func (ss *SpaceStore) GetSpace(ctx context.Context, spaceID int32) (*model.Space, error) {
	logger := logit.FromContext(ctx)

	sql, args := ss.newGetSpaceSQL(spaceID)
	logger.Debug("new get space sql", "sql", sql, "args", args)

	row := ss.db.QueryRowContext(ctx, sql, args...)
	if err := row.Err(); err != nil {
		logger.Error("query sql failed", "err", err, "sql", sql, "args", args)
		return nil, err
	}

	return ss.newSpace(row)
}

func (ss *SpaceStore) newSpaces(rows *stdsql.Rows) ([]*model.Space, error) {
	spaces := make([]*model.Space, 0, 4)

	for rows.Next() {
		space := new(model.Space)

		err := rows.Scan(&space.ID, &space.Name, &space.State, &space.CreateTime, &space.UpdateTime)
		if err != nil {
			return nil, err
		}

		spaces = append(spaces, space)
	}

	return spaces, nil
}

func (ss *SpaceStore) newListSpacesSQL(skip int64, limit int64, filter *model.ListSpacesFilter) (sql string, args []any) {
	var builder strings.Builder
	builder.WriteString("SELECT `id`, `name`, `state`, `create_time`, `update_time` FROM `spaces` WHERE `id` > 0")

	if filter.SpaceID > 0 {
		builder.WriteString(" AND `id` = ?")
		args = append(args, filter.SpaceID)
	}

	if filter.SpaceName != "" {
		builder.WriteString(" AND `name` LIKE '%?%'")
		args = append(args, filter.SpaceName)
	}

	if filter.SpaceState > 0 {
		builder.WriteString(" AND `state` = ?")
		args = append(args, filter.SpaceState)
	}

	builder.WriteString(" ORDER BY `id` DESC LIMIT ?, ?")

	sql = builder.String()
	args = append(args, skip, limit)
	return sql, args
}

func (ss *SpaceStore) ListSpaces(ctx context.Context, skip int64, limit int64, filter *model.ListSpacesFilter) ([]*model.Space, error) {
	logger := logit.FromContext(ctx)

	sql, args := ss.newListSpacesSQL(skip, limit, filter)
	logger.Debug("new list spaces sql", "sql", sql, "args", args)

	rows, err := ss.db.QueryContext(ctx, sql, args...)
	if err != nil {
		logger.Error("query sql failed", "err", err, "sql", sql, "args", args)
		return nil, err
	}

	defer rows.Close()

	return ss.newSpaces(rows)
}
