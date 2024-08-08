// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

type Email struct {
	TemplateID    int64             `json:"template_id"`
	To            []string          `json:"to"`
	CC            []string          `json:"cc"`
	BCC           []string          `json:"bcc"`
	SubjectParams map[string]string `json:"subject_params"`
	ContentParams map[string]string `json:"content_params"`
}
