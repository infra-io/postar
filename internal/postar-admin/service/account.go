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
	"github.com/infra-io/postar/pkg/aes"
)

type AccountStore interface {
	CreateAccount(ctx context.Context, spaceID int32, account *model.Account) error
	UpdateAccount(ctx context.Context, spaceID int32, account *model.Account) error
	GetAccount(ctx context.Context, spaceID int32, accountID int32) (*model.Account, error)
	ListAccounts(ctx context.Context, spaceID int32, skip int64, limit int64, filter *model.ListAccountsFilter) ([]*model.Account, error)
}

type defaultAccountService struct {
	conf         *config.PostarAdminConfig
	accountStore AccountStore
}

func NewAccountService(conf *config.PostarAdminConfig, accountStore AccountStore) AccountService {
	service := &defaultAccountService{
		conf:         conf,
		accountStore: accountStore,
	}

	return service
}

func (das *defaultAccountService) checkCreateAccountParams(account *model.Account) error {
	if strings.TrimSpace(account.Host) == "" {
		return errors.BadRequest("账号主机不能为空")
	}

	if account.Port <= 0 {
		return errors.BadRequest("账号端口需要大于 0")
	}

	if strings.TrimSpace(account.Username) == "" {
		return errors.BadRequest("账号用户名不能为空")
	}

	if account.Password == "" {
		return errors.BadRequest("账号密码不能为空")
	}

	if account.SMTPAuth <= 0 {
		return errors.BadRequest("未指定 SMTP 认证方式")
	}

	return nil
}

func (das *defaultAccountService) CreateAccount(ctx context.Context, spaceID int32, account *model.Account) (*model.Account, error) {
	logger := logit.FromContext(ctx)

	if err := das.checkCreateAccountParams(account); err != nil {
		logger.Error("check create account params failed", "err", err, "account", account)
		return nil, err
	}

	password := account.Password
	encrypted, err := aes.Encrypt(das.conf.Crypto.AESKey, das.conf.Crypto.AESIV, password)
	if err != nil {
		logger.Error("encrypt account password failed", "err", err)
		return nil, err
	}

	now := time.Now().Unix()
	account.Password = encrypted
	account.State = model.AccountStateEnabled
	account.CreateTime = now
	account.UpdateTime = now

	if err = das.accountStore.CreateAccount(ctx, spaceID, account); err != nil {
		logger.Error("create account failed", "err", err, "account", account)
		return nil, err
	}

	account.Password = password
	return account, nil
}

func (das *defaultAccountService) checkUpdateAccountParams(account *model.Account) error {
	if account.ID <= 0 {
		return errors.BadRequest("账号编号需要大于 0")
	}

	if account.State > 0 && !account.State.Valid() {
		return errors.BadRequest("账号状态 %d 无效", account.State)
	}

	return nil
}

func (das *defaultAccountService) UpdateAccount(ctx context.Context, spaceID int32, account *model.Account) (*model.Account, error) {
	logger := logit.FromContext(ctx)

	if err := das.checkUpdateAccountParams(account); err != nil {
		logger.Error("check update account params failed", "account", account)
		return nil, err
	}

	password := account.Password
	if password != "" {
		encrypted, err := aes.Encrypt(das.conf.Crypto.AESKey, das.conf.Crypto.AESIV, password)
		if err != nil {
			logger.Error("encrypt account password failed", "err", err)
			return nil, err
		}

		account.Password = encrypted
	}

	now := time.Now().Unix()
	account.UpdateTime = now

	if err := das.accountStore.UpdateAccount(ctx, spaceID, account); err != nil {
		logger.Error("update account failed", "err", err, "account", account)
		return nil, err
	}

	account.Password = password
	return account, nil
}

func (das *defaultAccountService) checkGetAccountParams(accountID int32) error {
	if accountID <= 0 {
		return errors.BadRequest("账号编号需要大于 0")
	}

	return nil
}

func (das *defaultAccountService) GetAccount(ctx context.Context, spaceID int32, accountID int32, withPassword bool) (*model.Account, error) {
	logger := logit.FromContext(ctx)

	if err := das.checkGetAccountParams(accountID); err != nil {
		logger.Error("check get account params failed", "err", err, "account_id", accountID)
		return nil, err
	}

	account, err := das.accountStore.GetAccount(ctx, spaceID, accountID)
	if err != nil {
		logger.Error("get account failed", "err", err, "account_id", accountID)
		return nil, err
	}

	if withPassword {
		decrypted, err := aes.Decrypt(das.conf.Crypto.AESKey, das.conf.Crypto.AESIV, account.Password)
		if err != nil {
			logger.Error("decrypt account password failed", "err", err, "password", account.Password)
			return nil, err
		}

		account.Password = decrypted
	} else {
		account.Password = ""
	}

	return account, nil
}

func (das *defaultAccountService) checkListAccountsParams(pageSize int32, filter *model.ListAccountsFilter) error {
	if pageSize < minPageSize || pageSize > maxPageSize {
		return errors.BadRequest("分页大小 %d 需要位于区间 [%d, %d] 内", pageSize, minPageSize, maxPageSize)
	}

	if filter.AccountID < 0 {
		return errors.BadRequest("账号编号不能为负数")
	}

	if filter.AccountState > 0 && !filter.AccountState.Valid() {
		return errors.BadRequest("账号状态 %d 无效", filter.AccountState)
	}

	return nil
}

func (das *defaultAccountService) removePasswordFromAccounts(accounts []*model.Account) []*model.Account {
	for i, account := range accounts {
		account.Password = ""
		accounts[i] = account
	}

	return accounts
}

func (das *defaultAccountService) ListAccounts(ctx context.Context, spaceID int32, pageSize int32, pageToken string, filter *model.ListAccountsFilter) ([]*model.Account, string, error) {
	logger := logit.FromContext(ctx)

	if err := das.checkListAccountsParams(pageSize, filter); err != nil {
		logger.Error("check list accounts params failed", "err", err, "page_size", pageSize, "page_token", pageToken, "filter", filter)
		return nil, "", err
	}

	skip, err := parsePageToken(pageToken)
	if err != nil {
		logger.Error("parse page token failed", "err", err, "page_size", pageSize, "page_token", pageToken, "filter", filter)
		return nil, "", err
	}

	accounts, err := das.accountStore.ListAccounts(ctx, spaceID, skip, int64(pageSize), filter)
	if err != nil {
		logger.Error("list accounts failed", "err", err, "skip", skip, "page_size", pageSize, "page_token", pageToken, "filter", filter)
		return nil, "", err
	}

	accounts = das.removePasswordFromAccounts(accounts)
	nextPageToken := newNextPageToken(skip, pageSize, len(accounts))
	return accounts, nextPageToken, nil
}
