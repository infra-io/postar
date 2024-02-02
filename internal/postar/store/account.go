// Copyright 2024 FishGoddeas. All rights reserved.
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

type AccountStore struct {
	conf *configs.PostarConfig
	db   *sql.DB
}

func NewAccountStore(conf *configs.PostarConfig) *AccountStore {
	store := &AccountStore{
		conf: conf,
		db:   db,
	}

	return store
}

func (as *AccountStore) newGetAccountSQL(spaceID int32, accountID int32) (sql string, args []any) {
	sql = "SELECT `id`, `host`, `port`, `username`, `password`, `smtp_auth`, `state` FROM `accounts` WHERE `space_id` = ? AND `id` = ?"
	args = []any{spaceID, accountID}

	return sql, args
}

func (as *AccountStore) newAccount(row *stdsql.Row) (*model.Account, error) {
	account := new(model.Account)

	err := row.Scan(&account.ID, &account.Host, &account.Port, &account.Username, &account.Password, &account.SMTPAuth, &account.State)
	if err == stdsql.ErrNoRows {
		return nil, errors.NotFound(err, errors.WithMsg("账号不存在"))
	}

	return account, err
}

func (as *AccountStore) GetAccount(ctx context.Context, spaceID int32, accountID int32) (*model.Account, error) {
	logger := logit.FromContext(ctx)

	sql, args := as.newGetAccountSQL(spaceID, accountID)
	logger.Debug("new get account sql", "sql", sql, "args", args)

	row := as.db.QueryRowContext(ctx, sql, args...)
	if err := row.Err(); err != nil {
		logger.Error("query sql failed", "err", err, "sql", sql, "args", args)
		return nil, err
	}

	return as.newAccount(row)
}
