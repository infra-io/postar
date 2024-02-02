// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

const (
	SMTPAuthPlain   SMTPAuth = 1
	SMTPAuthLogin   SMTPAuth = 2
	SMTPAuthCramMD5 SMTPAuth = 3
	SMTPAuthXOAUTH2 SMTPAuth = 4
)

const (
	AccountStateDisabled AccountState = 1
	AccountStateEnabled  AccountState = 2
)

type SMTPAuth int32

type AccountState int32

func (as AccountState) Valid() bool {
	return as == AccountStateDisabled || as == AccountStateEnabled
}

type Account struct {
	ID         int32        `json:"id"`
	Host       string       `json:"host"`
	Port       int32        `json:"port"`
	Username   string       `json:"username"`
	Password   string       `json:"-"`
	SMTPAuth   SMTPAuth     `json:"smtp_auth"`
	State      AccountState `json:"state"`
	CreateTime int64        `json:"create_time"`
	UpdateTime int64        `json:"update_time"`
}
