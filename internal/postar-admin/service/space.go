// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/FishGoddess/cryptox"
	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/config"
	"github.com/infra-io/postar/internal/postar-admin/model"
	"github.com/infra-io/postar/pkg/aes"
)

type SpaceStore interface {
	CreateSpace(ctx context.Context, space *model.Space) error
	UpdateSpace(ctx context.Context, space *model.Space) error
	GetSpace(ctx context.Context, spaceID int32) (*model.Space, error)
	ListSpaces(ctx context.Context, skip int64, limit int64, filter *model.ListSpacesFilter) ([]*model.Space, error)
}

type defaultSpaceService struct {
	conf       *config.PostarAdminConfig
	spaceStore SpaceStore
}

func NewSpaceService(conf *config.PostarAdminConfig, spaceStore SpaceStore) SpaceService {
	service := &defaultSpaceService{
		conf:       conf,
		spaceStore: spaceStore,
	}

	return service
}

func (dss *defaultSpaceService) checkCreateSpaceParams(space *model.Space) error {
	if strings.TrimSpace(space.Name) == "" {
		err := errors.New("trim space.Name == ''")
		return errors.BadRequest(err, errors.WithMsg("空间名称不能为空"))
	}

	return nil
}

func (dss *defaultSpaceService) CreateSpace(ctx context.Context, space *model.Space) (*model.Space, error) {
	logger := logit.FromContext(ctx)

	if err := dss.checkCreateSpaceParams(space); err != nil {
		logger.Error("check create space params failed", "err", err, "space", space)
		return nil, err
	}

	token := cryptox.GenerateString(64)
	encrypted, err := aes.Encrypt(dss.conf.Crypto.AESKey, dss.conf.Crypto.AESIV, token)
	if err != nil {
		logger.Error("encrypt token failed", "err", err)
		return nil, err
	}

	now := time.Now().Unix()
	space.Token = encrypted
	space.State = model.SpaceStateEnabled
	space.CreateTime = now
	space.UpdateTime = now

	if err = dss.spaceStore.CreateSpace(ctx, space); err != nil {
		logger.Error("create space failed", "err", err, "space", space)
		return nil, err
	}

	space.Token = token
	return space, err
}

func (dss *defaultSpaceService) checkUpdateSpaceParams(space *model.Space) error {
	if space.ID <= 0 {
		err := fmt.Errorf("space.ID %d <= 0", space.ID)
		return errors.BadRequest(err, errors.WithMsg("空间编号需要大于 0"))
	}

	if strings.TrimSpace(space.Name) == "" {
		err := errors.New("trim space.Name == ''")
		return errors.BadRequest(err, errors.WithMsg("空间名称不能为空"))
	}

	if space.State > 0 && !space.State.Valid() {
		err := fmt.Errorf("space.State %d not valid", space.State)
		return errors.BadRequest(err, errors.WithMsg("账号状态无效"))
	}

	return nil
}

func (dss *defaultSpaceService) UpdateSpace(ctx context.Context, space *model.Space) (*model.Space, error) {
	logger := logit.FromContext(ctx)

	if err := dss.checkUpdateSpaceParams(space); err != nil {
		logger.Error("check update space params failed", "err", err, "space", space)
		return nil, err
	}

	now := time.Now().Unix()
	space.UpdateTime = now

	if err := dss.spaceStore.UpdateSpace(ctx, space); err != nil {
		logger.Error("update space failed", "err", err, "space", space)
		return nil, err
	}

	return space, nil
}

func (dss *defaultSpaceService) checkGetSpaceParams(spaceID int32) error {
	if spaceID <= 0 {
		err := fmt.Errorf("spaceID %d <= 0", spaceID)
		return errors.BadRequest(err, errors.WithMsg("空间编号非法"))
	}

	return nil
}

func (dss *defaultSpaceService) GetSpace(ctx context.Context, spaceID int32, withToken bool) (*model.Space, error) {
	logger := logit.FromContext(ctx)

	if err := dss.checkGetSpaceParams(spaceID); err != nil {
		logger.Error("check get space params failed", "err", err)
		return nil, err
	}

	space, err := dss.spaceStore.GetSpace(ctx, spaceID)
	if err != nil {
		logger.Error("get space failed", "err", err, "space_id", spaceID)
		return nil, err
	}

	if withToken {
		decrypted, err := aes.Decrypt(dss.conf.Crypto.AESKey, dss.conf.Crypto.AESIV, space.Token)
		if err != nil {
			logger.Error("decrypt space token failed", "err", err, "token", space.Token)
			return nil, err
		}

		space.Token = decrypted
	} else {
		space.Token = ""
	}

	return space, nil
}

func (dss *defaultSpaceService) checkListSpacesParams(pageSize int32, filter *model.ListSpacesFilter) error {
	if pageSize < minPageSize || pageSize > maxPageSize {
		err := fmt.Errorf("pageSize %d not in [%d, %d]", pageSize, minPageSize, maxPageSize)
		return errors.BadRequest(err, errors.WithMsg("分页大小需要位于区间 [%d, %d] 内", minPageSize, maxPageSize))
	}

	if filter.SpaceID < 0 {
		err := fmt.Errorf("filter.SpaceID %d < 0", filter.SpaceID)
		return errors.BadRequest(err, errors.WithMsg("过滤的空间编号非法"))
	}

	if filter.SpaceState > 0 && !filter.SpaceState.Valid() {
		err := fmt.Errorf("filter.SpaceState %d not valid", filter.SpaceState)
		return errors.BadRequest(err, errors.WithMsg("过滤的空间状态非法"))
	}

	return nil
}

func (dss *defaultSpaceService) ListSpaces(ctx context.Context, pageSize int32, pageToken string, filter *model.ListSpacesFilter) ([]*model.Space, string, error) {
	logger := logit.FromContext(ctx)

	if err := dss.checkListSpacesParams(pageSize, filter); err != nil {
		logger.Error("check list spaces params failed", "err", err, "page_size", pageSize, "page_token", pageToken, "filter", filter)
		return nil, "", err
	}

	skip, err := parsePageToken(pageToken)
	if err != nil {
		logger.Error("parse page token failed", "err", err, "page_size", pageSize, "page_token", pageToken, "filter", filter)
		return nil, "", err
	}

	spaces, err := dss.spaceStore.ListSpaces(ctx, skip, int64(pageSize), filter)
	if err != nil {
		logger.Error("list spaces failed", "err", err, "skip", skip, "limit", pageSize, "filter", filter)
		return nil, "", err
	}

	nextPageToken := newNextPageToken(skip, pageSize, len(spaces))
	return spaces, nextPageToken, nil
}
