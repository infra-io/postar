// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

const (
	EmailContentTypePlain EmailContentType = 1
	EmailContentTypeHTML  EmailContentType = 2
)

const (
	TemplateStateDeleted  TemplateState = 1
	TemplateStateDisabled TemplateState = 2
	TemplateStateEnabled  TemplateState = 3
)

type EmailContentType int32

func (ect EmailContentType) Valid() bool {
	return ect == EmailContentTypePlain || ect == EmailContentTypeHTML
}

type TemplateEmail struct {
	Subject     string           `json:"subject"`
	To          []string         `json:"to"`
	CC          []string         `json:"cc"`
	BCC         []string         `json:"bcc"`
	ContentType EmailContentType `json:"content_type"`
	Content     string           `json:"content"`
}

type TemplateState int32

func (ts TemplateState) Valid() bool {
	return ts == TemplateStateDeleted || ts == TemplateStateDisabled || ts == TemplateStateEnabled
}

type Template struct {
	ID          int64         `json:"id"`
	AccountID   int32         `json:"account_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Email       TemplateEmail `json:"email"`
	State       TemplateState `json:"state"`
	CreateTime  int64         `json:"create_time"`
	UpdateTime  int64         `json:"update_time"`
}
