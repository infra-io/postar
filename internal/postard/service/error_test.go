// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 22:06:36

package service

import (
	"errors"
	"testing"
)

// go test -v -cover -run=^TestIsSendTimeout$
func TestIsSendTimeout(t *testing.T) {
	testCases := []struct {
		err    error
		result bool
	}{
		{errSendTimeout, true},
		{errSendEmailFailed, false},
		{errors.New("unknown error"), false},
	}

	for i, testCase := range testCases {
		if IsSendTimeout(testCase.err) != testCase.result {
			t.Errorf("testCase %d failed with err %+v, result %+v", i, testCase.err, testCase.result)
		}
	}
}

// go test -v -cover -run=^TestIsSendEmailFailed$
func TestIsSendEmailFailed(t *testing.T) {
	testCases := []struct {
		err    error
		result bool
	}{
		{errSendTimeout, false},
		{errSendEmailFailed, true},
		{errors.New("unknown error"), false},
	}

	for i, testCase := range testCases {
		if IsSendEmailFailed(testCase.err) != testCase.result {
			t.Errorf("testCase %d failed with err %+v, result %+v", i, testCase.err, testCase.result)
		}
	}
}
