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
	spaceCacheTTL = 30 * time.Second
)

type SpaceStore struct {
	cache      cachego.Cache
	spaceStore service.SpaceStore
}

func NewSpaceStore(spaceStore service.SpaceStore) service.SpaceStore {
	store := &SpaceStore{
		cache:      cache.New("space"),
		spaceStore: spaceStore,
	}

	return store
}

func (ss *SpaceStore) spaceCacheKey(spaceID int32) string {
	return "space:" + strconv.FormatInt(int64(spaceID), 10)
}

func (ss *SpaceStore) GetSpace(ctx context.Context, spaceID int32) (space *model.Space, err error) {
	cacheKey := ss.spaceCacheKey(spaceID)

	value, ok := ss.cache.Get(cacheKey)
	if !ok {
		load := func() (value interface{}, err error) {
			return ss.spaceStore.GetSpace(ctx, spaceID)
		}

		value, err = ss.cache.Load(cacheKey, spaceCacheTTL, load)
		if err != nil {
			return nil, err
		}
	}

	return value.(*model.Space), nil
}
