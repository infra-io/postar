// Copyright 2023 FishGoddeas. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"context"
	"database/sql"
	stdsql "database/sql"
	"strings"

	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/config"
	"github.com/infra-io/postar/internal/postar-admin/model"
)

type AccountStore struct {
	conf *config.PostarAdminConfig
	db   *sql.DB
}

func NewAccountStore(conf *config.PostarAdminConfig) *AccountStore {
	store := &AccountStore{
		conf: conf,
		db:   db,
	}

	return store
}

func (as *AccountStore) newCreateAccountSQL(spaceID int32, account *model.Account) (sql string, args []any) {
	sql = "INSERT INTO `accounts`(`space_id`, `host`, `port`, `username`, `password`, `smtp_auth`, `state`, `create_time`, `update_time`) VALUE(?,?,?,?,?,?,?,?,?)"
	args = []any{spaceID, account.Host, account.Port, account.Username, account.Password, account.SMTPAuth, account.State, account.CreateTime, account.UpdateTime}

	return sql, args
}

func (as *AccountStore) CreateAccount(ctx context.Context, spaceID int32, account *model.Account) error {
	logger := logit.FromContext(ctx)

	sql, args := as.newCreateAccountSQL(spaceID, account)
	logger.Debug("new create account sql", "sql", sql, "args", args)

	stmt, err := as.db.PrepareContext(ctx, sql)
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
		logger.Error("get last insert id failed", "err", err, "account", account)
		err = nil // Ignore this error because account has been created.
	}

	account.ID = int32(id)
	return nil
}

func (as *AccountStore) newUpdateAccountSQL(spaceID int32, account *model.Account) (sql string, args []any) {
	var builder strings.Builder

	builder.WriteString("UPDATE `accounts` SET `update_time` = ?")
	args = append(args, account.UpdateTime)

	if account.Host != "" {
		builder.WriteString(", `host` = ?")
		args = append(args, account.Host)
	}

	if account.Port > 0 {
		builder.WriteString(", `port` = ?")
		args = append(args, account.Port)
	}

	if account.Username != "" {
		builder.WriteString(", `username` = ?")
		args = append(args, account.Username)
	}

	if account.Password != "" {
		builder.WriteString(", `password` = ?")
		args = append(args, account.Password)
	}

	if account.SMTPAuth > 0 {
		builder.WriteString(", `smtp_auth` = ?")
		args = append(args, account.SMTPAuth)
	}

	if account.State > 0 {
		builder.WriteString(", `state` = ?")
		args = append(args, account.State)
	}

	builder.WriteString(" WHERE `space_id` = ? AND `id` = ?")
	args = append(args, spaceID, account.ID)

	sql = builder.String()
	return sql, args
}

func (as *AccountStore) UpdateAccount(ctx context.Context, spaceID int32, account *model.Account) error {
	logger := logit.FromContext(ctx)

	sql, args := as.newUpdateAccountSQL(spaceID, account)
	logger.Debug("new update account sql", "sql", sql, "args", args)

	stmt, err := as.db.PrepareContext(ctx, sql)
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

func (as *AccountStore) newGetAccountSQL(spaceID int32, accountID int32) (sql string, args []any) {
	sql = "SELECT `id`, `host`, `port`,`username`, `password`, `smtp_auth`, `state`, `create_time`, `update_time` FROM `accounts` WHERE `space_id` = ? AND `id` = ?"
	args = []any{spaceID, accountID}

	return sql, args
}

func (as *AccountStore) newAccount(row *stdsql.Row) (*model.Account, error) {
	account := new(model.Account)

	err := row.Scan(&account.ID, &account.Host, &account.Port, &account.Username, &account.Password, &account.SMTPAuth, &account.State, &account.CreateTime, &account.UpdateTime)
	if err == stdsql.ErrNoRows {
		return nil, errors.NotFound("账号不存在").With(err)
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

func (as *AccountStore) newListAccountsSQL(spaceID int32, skip int64, limit int64, filter *model.ListAccountsFilter) (sql string, args []any) {
	var builder strings.Builder

	builder.WriteString("SELECT `id`, `host`, `port`, `username`, `smtp_auth`, `state`, `create_time`, `update_time` FROM `accounts` WHERE `space_id` = ?")
	args = append(args, spaceID)

	if filter.AccountID > 0 {
		builder.WriteString(" AND `id` = ?")
		args = append(args, filter.AccountID)
	}

	if filter.AccountUsername != "" {
		builder.WriteString(" AND `username` = ?")
		args = append(args, filter.AccountUsername)
	}

	if filter.AccountHost != "" {
		builder.WriteString(" AND `host` = ?")
		args = append(args, filter.AccountHost)
	}

	if filter.AccountState > 0 {
		builder.WriteString(" AND `state` = ?")
		args = append(args, filter.AccountState)
	}

	builder.WriteString(" ORDER BY `id` DESC LIMIT ?, ?")
	args = append(args, skip, limit)

	sql = builder.String()
	return sql, args
}

func (as *AccountStore) newAccounts(rows *stdsql.Rows) ([]*model.Account, error) {
	accounts := make([]*model.Account, 0, 4)

	for rows.Next() {
		account := new(model.Account)

		err := rows.Scan(&account.ID, &account.Host, &account.Port, &account.Username, &account.SMTPAuth, &account.State, &account.CreateTime, &account.UpdateTime)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (as *AccountStore) ListAccounts(ctx context.Context, spaceID int32, skip int64, limit int64, filter *model.ListAccountsFilter) ([]*model.Account, error) {
	logger := logit.FromContext(ctx)

	sql, args := as.newListAccountsSQL(spaceID, skip, limit, filter)
	logger.Debug("new list accounts sql", "sql", sql, "args", args)

	rows, err := as.db.QueryContext(ctx, sql, args...)
	if err != nil {
		logger.Error("query sql failed", "err", err, "sql", sql, "args", args)
		return nil, err
	}

	defer rows.Close()

	return as.newAccounts(rows)
}
