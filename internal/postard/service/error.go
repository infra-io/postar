// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 21:58:02

package service

import "errors"

var (
	errSendTimeout     = errors.New("postard/service: send timeout")      // Send timeout
	errSendEmailFailed = errors.New("postard/service: send email failed") // Send email failed
)

// IsSendTimeout returns if err equals to errSendTimeout.
func IsSendTimeout(err error) bool {
	return errors.Is(err, errSendTimeout)
}

// IsSendEmailFailed returns if err equals to errSendEmailFailed.
func IsSendEmailFailed(err error) bool {
	return errors.Is(err, errSendEmailFailed)
}
