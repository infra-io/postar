// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

type ListSpacesFilter struct {
	SpaceID    int32      `json:"space_id"`
	SpaceName  string     `json:"space_name"`
	SpaceState SpaceState `json:"space_state"`
}

type ListAccountsFilter struct {
	AccountID       int32        `json:"account_id"`
	AccountUsername string       `json:"account_username"`
	AccountHost     string       `json:"account_host"`
	AccountState    AccountState `json:"account_state"`
}

type ListTemplatesFilter struct {
	AccountID     int32         `json:"account_id"`
	TemplateID    int64         `json:"template_id"`
	TemplateName  string        `json:"template_name"`
	TemplateState TemplateState `json:"template_state"`
	EmailSubject  string        `json:"email_subject"`
}
