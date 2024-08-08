// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

const (
	SpaceStateDisabled SpaceState = 1
	SpaceStateEnabled  SpaceState = 2
)

type SpaceState int32

func (ss SpaceState) Valid() bool {
	return ss == SpaceStateDisabled || ss == SpaceStateEnabled
}

type Space struct {
	ID         int32      `json:"id"`
	Name       string     `json:"name"`
	Token      string     `json:"-"`
	State      SpaceState `json:"state"`
	CreateTime int64      `json:"create_time"`
	UpdateTime int64      `json:"update_time"`
}
