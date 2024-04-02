// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"context"
	stdsql "database/sql"

	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/config"
	"github.com/infra-io/postar/internal/postar/model"
)

type SpaceStore struct {
	conf *config.PostarConfig
	db   *stdsql.DB
}

func NewSpaceStore(conf *config.PostarConfig) *SpaceStore {
	store := &SpaceStore{
		conf: conf,
		db:   db,
	}

	return store
}

func (ss *SpaceStore) newGetSpaceSQL(spaceID int32) (sql string, args []any) {
	sql = "SELECT `id`, `name`, `token`, `state` FROM `spaces` WHERE `id` = ?"
	args = []any{spaceID}

	return sql, args
}

func (ss *SpaceStore) newSpace(row *stdsql.Row) (*model.Space, error) {
	space := new(model.Space)

	err := row.Scan(&space.ID, &space.Name, &space.Token, &space.State)
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
