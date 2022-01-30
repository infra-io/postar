// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/16 02:05:02

package log

import (
	"github.com/avinoplan/postar/configs"
	"testing"
)

// go test -v -cover -run=^TestInitialize$
func TestInitialize(t *testing.T) {
	if globalLogger != nil {
		t.Errorf("globalLogger %+v != nil", globalLogger)
	}

	c := configs.NewDefaultConfig()
	err := Initialize(c)
	if err != nil {
		t.Error(err)
	}

	if globalLogger == nil {
		t.Error("globalLogger == nil")
	}
}
