// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import "github.com/FishGoddess/errors"

const (
	codeSendEmailFailed = 11000
)

// SendEmailFailedErr returns a send email failed error.
func SendEmailFailedErr(err error) error {
	return errors.Wrap(err, codeSendEmailFailed)
}

// IsSendEmailFailed returns if err is send email failed.
func IsSendEmailFailed(err error) bool {
	return errors.Is(err, codeSendEmailFailed)
}
