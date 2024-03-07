// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	postaradminv1 "github.com/infra-io/postar/api/genproto/postaradmin/v1"
	"github.com/infra-io/postar/internal/postar-admin/model"
)

func newSpace(space *postaradminv1.Space) *model.Space {
	if space == nil {
		return new(model.Space)
	}

	result := &model.Space{
		ID:         space.Id,
		Name:       space.Name,
		Token:      space.Token,
		State:      model.SpaceState(space.State),
		CreateTime: space.CreateTime,
		UpdateTime: space.UpdateTime,
	}

	return result
}

func fromSpace(space *model.Space) *postaradminv1.Space {
	if space == nil {
		return new(postaradminv1.Space)
	}

	result := &postaradminv1.Space{
		Id:         space.ID,
		Name:       space.Name,
		Token:      space.Token,
		State:      postaradminv1.SpaceState(space.State),
		CreateTime: space.CreateTime,
		UpdateTime: space.UpdateTime,
	}

	return result
}

func fromSpaces(spaces []*model.Space) []*postaradminv1.Space {
	result := make([]*postaradminv1.Space, 0, len(spaces))
	for _, space := range spaces {
		result = append(result, fromSpace(space))
	}

	return result
}

func newListSpacesFilter(filter *postaradminv1.ListSpacesFilter) *model.ListSpacesFilter {
	if filter == nil {
		return new(model.ListSpacesFilter)
	}

	result := &model.ListSpacesFilter{
		SpaceName:  filter.SpaceName,
		SpaceState: model.SpaceState(filter.SpaceState),
	}

	return result
}

func newCreateSpaceResponse(space *model.Space) *postaradminv1.CreateSpaceResponse {
	result := &postaradminv1.CreateSpaceResponse{
		Space: fromSpace(space),
	}

	return result
}

func newUpdateSpaceResponse(space *model.Space) *postaradminv1.UpdateSpaceResponse {
	result := &postaradminv1.UpdateSpaceResponse{
		Space: fromSpace(space),
	}

	return result
}

func newGetSpaceResponse(space *model.Space) *postaradminv1.GetSpaceResponse {
	result := &postaradminv1.GetSpaceResponse{
		Space: fromSpace(space),
	}

	return result
}

func newListSpacesResponse(spaces []*model.Space, nextPageToken string) *postaradminv1.ListSpacesResponse {
	result := &postaradminv1.ListSpacesResponse{
		Spaces:        fromSpaces(spaces),
		NextPageToken: nextPageToken,
	}

	return result
}

func (gs *GrpcServer) CreateSpace(ctx context.Context, request *postaradminv1.CreateSpaceRequest) (response *postaradminv1.CreateSpaceResponse, err error) {
	space := newSpace(request.Space)

	createdSpace, err := gs.spaceService.CreateSpace(ctx, space)
	if err != nil {
		return nil, err
	}

	response = newCreateSpaceResponse(createdSpace)
	return response, nil
}

func (gs *GrpcServer) UpdateSpace(ctx context.Context, request *postaradminv1.UpdateSpaceRequest) (response *postaradminv1.UpdateSpaceResponse, err error) {
	space := newSpace(request.Space)

	updatedSpace, err := gs.spaceService.UpdateSpace(ctx, space)
	if err != nil {
		return nil, err
	}

	response = newUpdateSpaceResponse(updatedSpace)
	return response, nil
}

func (gs *GrpcServer) GetSpace(ctx context.Context, request *postaradminv1.GetSpaceRequest) (response *postaradminv1.GetSpaceResponse, err error) {
	space, err := gs.spaceService.GetSpace(ctx, request.SpaceId, request.WithToken)
	if err != nil {
		return nil, err
	}

	response = newGetSpaceResponse(space)
	return response, nil
}

func (gs *GrpcServer) ListSpaces(ctx context.Context, request *postaradminv1.ListSpacesRequest) (response *postaradminv1.ListSpacesResponse, err error) {
	filter := newListSpacesFilter(request.Filter)

	spaces, nextPageToken, err := gs.spaceService.ListSpaces(ctx, request.PageSize, request.PageToken, filter)
	if err != nil {
		return nil, err
	}

	response = newListSpacesResponse(spaces, nextPageToken)
	return response, nil
}
