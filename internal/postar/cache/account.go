// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/FishGoddess/cachego"
	"github.com/infra-io/postar/internal/postar/model"
	"github.com/infra-io/postar/internal/postar/service"
	"github.com/infra-io/servicex/cache"
)

const (
	accountCacheTTL = 3 * time.Second
)

type AccountStore struct {
	cache        cachego.Cache
	accountStore service.AccountStore
}

func NewAccountStore(accountStore service.AccountStore) service.AccountStore {
	store := &AccountStore{
		cache:        cache.New("account"),
		accountStore: accountStore,
	}

	return store
}

func (ss *AccountStore) accountCacheKey(spaceID int32, accountID int32) string {
	return "account:" + strconv.FormatInt(int64(spaceID), 10) + ":" + strconv.FormatInt(int64(accountID), 10)
}

func (ss *AccountStore) GetAccount(ctx context.Context, spaceID int32, accountID int32) (account *model.Account, err error) {
	cacheKey := ss.accountCacheKey(spaceID, accountID)

	value, ok := ss.cache.Get(cacheKey)
	if !ok {
		load := func() (value interface{}, err error) {
			return ss.accountStore.GetAccount(ctx, spaceID, accountID)
		}

		value, err = ss.cache.Load(cacheKey, accountCacheTTL, load)
		if err != nil {
			return nil, err
		}
	}

	return value.(*model.Account), nil
}
