// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 22:58:39

package encode

import (
	"os"
	"testing"
	"time"
)

// go test -v -cover -run=^TestPidHex$
func TestPidHex(t *testing.T) {
	pidHex := PIDHex()
	if len(pidHex) != 4 {
		t.Errorf("length of PIDHex is wrong with %s, %d", pidHex, len(pidHex))
	}

	pid := uint64(os.Getpid())
	t.Log(pid, numberHex(pid, 0, 0), pidHex)
}

// go test -v -cover -run=^NowTimeHex$
func TestNowTimeHex(t *testing.T) {
	timeHex := NowHex()
	if len(timeHex) != 8 {
		t.Errorf("length of TimeHex is wrong with %s, %d", timeHex, len(timeHex))
	}

	now := uint64(time.Now().Unix())
	t.Log(now, numberHex(now, 0, 0), timeHex)
}

// go test -v -cover -run=^TestStringHex$
func TestStringHex(t *testing.T) {
	length := 16
	str := StringHex(length)
	if len(str) != length {
		t.Errorf("length of StringHex is wrong with %d", len(str))
	}
}
