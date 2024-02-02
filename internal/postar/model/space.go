// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

const (
	SpaceStateEnabled SpaceState = 2
)

type SpaceState int32

type Space struct {
	ID    int32      `json:"id"`
	Name  string     `json:"name"`
	Token string     `json:"-"`
	State SpaceState `json:"state"`
}

func (s *Space) Enabled() bool {
	return s.State == SpaceStateEnabled
}
