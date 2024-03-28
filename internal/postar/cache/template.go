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
	templateCacheTTL = 3 * time.Second
)

type TemplateStore struct {
	cache         cachego.Cache
	templateStore service.TemplateStore
}

func WrapTemplateStore(templateStore service.TemplateStore) service.TemplateStore {
	store := &TemplateStore{
		cache:         cache.New("template"),
		templateStore: templateStore,
	}

	return store
}

func (ss *TemplateStore) templateCacheKey(spaceID int32, templateID int64) string {
	return "template:" + strconv.FormatInt(int64(spaceID), 10) + ":" + strconv.FormatInt(templateID, 10)
}

func (ss *TemplateStore) GetTemplate(ctx context.Context, spaceID int32, templateID int64) (template *model.Template, err error) {
	cacheKey := ss.templateCacheKey(spaceID, templateID)

	value, ok := ss.cache.Get(cacheKey)
	if !ok {
		load := func() (value interface{}, err error) {
			return ss.templateStore.GetTemplate(ctx, spaceID, templateID)
		}

		value, err = ss.cache.Load(cacheKey, templateCacheTTL, load)
		if err != nil {
			return nil, err
		}
	}

	return value.(*model.Template), nil
}
