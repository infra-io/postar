// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 21:58:02

package errors

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	codeBadRequest      = http.StatusBadRequest
	codeSendTimeout     = 1100
	codeSendEmailFailed = 11000
)

// Error is error with code.
type Error struct {
	err  error
	code int32
}

// Error returns error message.
func (e *Error) Error() string {
	if e == nil || e.err == nil {
		return ""
	}
	return fmt.Sprintf("%d (%s)", e.code, e.err.Error())
}

// Is returns if e is the target's type.
func (e *Error) Is(target error) bool {
	if e == nil {
		return e == target
	}

	err, ok := target.(*Error)
	if !ok {
		return e.err == target
	}

	return e.code == err.code
}

// Unwrap unwraps e.
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

// WithCode wraps err with code.
func WithCode(err error, code int32) error {
	if err == nil {
		return nil
	}
	return &Error{err: err, code: code}
}

// Is returns if err has a code and its code equals to code.
func Is(err error, code int32) bool {
	for {
		if err == nil {
			return false
		}

		e, ok := err.(*Error)
		if !ok {
			return false
		}

		if e.code == code {
			return true
		}
		err = errors.Unwrap(err)
	}
}

// BadRequest returns a bad request error.
func BadRequest(err error) error {
	return WithCode(err, codeBadRequest)
}

// IsBadRequest returns if err is bad request.
func IsBadRequest(err error) bool {
	return Is(err, codeBadRequest)
}

// SendTimeoutErr returns a send timeout error.
func SendTimeoutErr(err error) error {
	return WithCode(err, codeSendTimeout)
}

// IsSendTimeout returns if err is send timeout.
func IsSendTimeout(err error) bool {
	return Is(err, codeSendTimeout)
}

// SendEmailFailedErr returns a send email failed error.
func SendEmailFailedErr(err error) error {
	return WithCode(err, codeSendEmailFailed)
}

// IsSendEmailFailed returns if err is send email failed.
func IsSendEmailFailed(err error) bool {
	return Is(err, codeSendEmailFailed)
}
