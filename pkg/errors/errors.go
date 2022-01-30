// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 21:58:02

package errors

import "github.com/FishGoddess/errors"

const (
	codeBadRequest      = 400
	codeSendEmailFailed = 11000
)

// TimeoutErr returns a timeout error.
func TimeoutErr(err error) error {
	return errors.Timeout(err)
}

// IsTimeout returns if err is timeout.
func IsTimeout(err error) bool {
	return errors.IsTimeout(err)
}

// BadRequestErr returns a bad request error.
func BadRequestErr(err error) error {
	return errors.Wrap(err, codeBadRequest)
}

// IsBadRequest returns if err is bad request.
func IsBadRequest(err error) bool {
	return errors.Is(err, codeBadRequest)
}

// SendEmailFailedErr returns a send email failed error.
func SendEmailFailedErr(err error) error {
	return errors.Wrap(err, codeSendEmailFailed)
}

// IsSendEmailFailed returns if err is send email failed.
func IsSendEmailFailed(err error) bool {
	return errors.Is(err, codeSendEmailFailed)
}
