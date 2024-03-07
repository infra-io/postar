// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

const (
	EmailContentTypePlain EmailContentType = 1
	EmailContentTypeHTML  EmailContentType = 2
)

const (
	TemplateStateDeleted TemplateState = 1
	TemplateStateEnabled TemplateState = 3
)

type EmailContentType int32

func (ect EmailContentType) String() string {
	if ect == EmailContentTypePlain {
		return "text/plain"
	}

	if ect == EmailContentTypeHTML {
		return "text/html"
	}

	return "text/plain"
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

type Template struct {
	ID        int64         `json:"id"`
	AccountID int32         `json:"account_id"`
	Name      string        `json:"name"`
	Email     TemplateEmail `json:"email"`
	State     TemplateState `json:"state"`
}

func (t *Template) Enabled() bool {
	return t.State == TemplateStateEnabled
}
