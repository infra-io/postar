// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 22:06:36

package errors

import (
	"errors"
	"testing"
)

// go test -v -cover -run=^TestIsBadRequest$
func TestIsBadRequest(t *testing.T) {
	testCases := []struct {
		err    error
		result bool
	}{
		{BadRequest(errors.New("bad request")), true},
		{errors.New("unknown error"), false},
	}

	for i, testCase := range testCases {
		if IsBadRequest(testCase.err) != testCase.result {
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
		{SendEmailFailedErr(errors.New("send email failed")), true},
		{errors.New("unknown error"), false},
	}

	for i, testCase := range testCases {
		if IsSendEmailFailed(testCase.err) != testCase.result {
			t.Errorf("testCase %d failed with err %+v, result %+v", i, testCase.err, testCase.result)
		}
	}
}
