// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 22:58:39

package trace

import "testing"

// go test -v -cover -run=^TestTimeHex$
func TestTimeHex(t *testing.T) {
	timeHex := TimeHex()
	if len(timeHex) != 8 {
		t.Errorf("length of TimeHex is wrong with %s, %d", timeHex, len(timeHex))
	}
}

// go test -v -cover -run=^TestPidHex$
func TestPidHex(t *testing.T) {
	pidHex := PidHex()
	if len(pidHex) != 8 {
		t.Errorf("length of PidHex is wrong with %s, %d", pidHex, len(pidHex))
	}
}

// go test -v -cover -run=^TestRandomString$
func TestRandomString(t *testing.T) {
	length := 16
	str := RandomString(length)
	if len(str) != length {
		t.Errorf("length of RandomString is wrong with %d", len(str))
	}
}
