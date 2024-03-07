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
	AccountStateEnabled AccountState = 2
)

type SMTPAuth int32

func (sa SMTPAuth) String() string {
	if sa == SMTPAuthPlain {
		return "PLAIN"
	}

	if sa == SMTPAuthLogin {
		return "LOGIN"
	}

	if sa == SMTPAuthCramMD5 {
		return "CRAM-MD5"
	}

	if sa == SMTPAuthXOAUTH2 {
		return "XOAUTH2"
	}

	return "PLAIN"
}

type AccountState int32

type Account struct {
	ID       int32        `json:"id"`
	Host     string       `json:"host"`
	Port     int32        `json:"port"`
	Username string       `json:"username"`
	Password string       `json:"-"`
	SMTPAuth SMTPAuth     `json:"smtp_auth"`
	State    AccountState `json:"state"`
}

func (a Account) Enabled() bool {
	return a.State == AccountStateEnabled
}
