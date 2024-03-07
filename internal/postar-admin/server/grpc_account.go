// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	postaradminv1 "github.com/infra-io/postar/api/genproto/postaradmin/v1"
	"github.com/infra-io/postar/internal/postar-admin/model"
	"github.com/infra-io/postar/pkg/grpc/contextutil"
)

func newAccount(account *postaradminv1.Account) *model.Account {
	if account == nil {
		return new(model.Account)
	}

	result := &model.Account{
		ID:         account.Id,
		Host:       account.Host,
		Port:       account.Port,
		Username:   account.Username,
		Password:   account.Password,
		SMTPAuth:   model.SMTPAuth(account.SmtpAuth),
		State:      model.AccountState(account.State),
		CreateTime: account.CreateTime,
		UpdateTime: account.UpdateTime,
	}

	return result
}

func fromAccount(account *model.Account) *postaradminv1.Account {
	if account == nil {
		return new(postaradminv1.Account)
	}

	result := &postaradminv1.Account{
		Id:         account.ID,
		Host:       account.Host,
		Port:       account.Port,
		Username:   account.Username,
		Password:   account.Password,
		SmtpAuth:   postaradminv1.SMTPAuth(account.SMTPAuth),
		State:      postaradminv1.AccountState(account.State),
		CreateTime: account.CreateTime,
		UpdateTime: account.UpdateTime,
	}

	return result
}

func fromAccounts(accounts []*model.Account) []*postaradminv1.Account {
	result := make([]*postaradminv1.Account, 0, len(accounts))
	for _, account := range accounts {
		result = append(result, fromAccount(account))
	}

	return result
}

func newListAccountsFilter(filter *postaradminv1.ListAccountsFilter) *model.ListAccountsFilter {
	if filter == nil {
		return new(model.ListAccountsFilter)
	}

	result := &model.ListAccountsFilter{
		AccountID:       filter.AccountId,
		AccountUsername: filter.AccountUsername,
		AccountHost:     filter.AccountHost,
		AccountState:    model.AccountState(filter.AccountState),
	}

	return result
}

func newCreateAccountResponse(account *model.Account) *postaradminv1.CreateAccountResponse {
	result := &postaradminv1.CreateAccountResponse{
		Account: fromAccount(account),
	}

	return result
}

func newUpdateAccountResponse(account *model.Account) *postaradminv1.UpdateAccountResponse {
	result := &postaradminv1.UpdateAccountResponse{
		Account: fromAccount(account),
	}

	return result
}

func newGetAccountResponse(account *model.Account) *postaradminv1.GetAccountResponse {
	result := &postaradminv1.GetAccountResponse{
		Account: fromAccount(account),
	}

	return result
}

func newListAccountsResponse(accounts []*model.Account, nextPageToken string) *postaradminv1.ListAccountsResponse {
	result := &postaradminv1.ListAccountsResponse{
		Accounts:      fromAccounts(accounts),
		NextPageToken: nextPageToken,
	}

	return result
}

func (gs *GrpcServer) CreateAccount(ctx context.Context, request *postaradminv1.CreateAccountRequest) (response *postaradminv1.CreateAccountResponse, err error) {
	spaceID := contextutil.GetSpaceID(ctx)
	account := newAccount(request.Account)

	createdAccount, err := gs.accountService.CreateAccount(ctx, spaceID, account)
	if err != nil {
		return nil, err
	}

	response = newCreateAccountResponse(createdAccount)
	return response, nil
}

func (gs *GrpcServer) UpdateAccount(ctx context.Context, request *postaradminv1.UpdateAccountRequest) (response *postaradminv1.UpdateAccountResponse, err error) {
	spaceID := contextutil.GetSpaceID(ctx)
	account := newAccount(request.Account)

	updatedAccount, err := gs.accountService.UpdateAccount(ctx, spaceID, account)
	if err != nil {
		return nil, err
	}

	response = newUpdateAccountResponse(updatedAccount)
	return response, nil
}

func (gs *GrpcServer) GetAccount(ctx context.Context, request *postaradminv1.GetAccountRequest) (response *postaradminv1.GetAccountResponse, err error) {
	spaceID := contextutil.GetSpaceID(ctx)

	account, err := gs.accountService.GetAccount(ctx, spaceID, request.AccountId, request.WithPassword)
	if err != nil {
		return nil, err
	}

	response = newGetAccountResponse(account)
	return response, nil
}

func (gs *GrpcServer) ListAccounts(ctx context.Context, request *postaradminv1.ListAccountsRequest) (response *postaradminv1.ListAccountsResponse, err error) {
	spaceID := contextutil.GetSpaceID(ctx)
	filter := newListAccountsFilter(request.Filter)

	accounts, nextPageToken, err := gs.accountService.ListAccounts(ctx, spaceID, request.PageSize, request.PageToken, filter)
	if err != nil {
		return nil, err
	}

	response = newListAccountsResponse(accounts, nextPageToken)
	return response, nil
}
