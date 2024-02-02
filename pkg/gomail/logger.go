// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gomail

import (
	"fmt"

	"github.com/FishGoddess/logit"
)

type Logger struct{}

func (l Logger) Debugf(format string, v ...interface{}) {
	if len(v) > 0 {
		format = fmt.Sprintf(format, v...)
	}

	logit.Debug(format)
}

func (l Logger) Infof(format string, v ...interface{}) {
	if len(v) > 0 {
		format = fmt.Sprintf(format, v...)
	}

	logit.Info(format)
}

func (l Logger) Warnf(format string, v ...interface{}) {
	if len(v) > 0 {
		format = fmt.Sprintf(format, v...)
	}

	logit.Warn(format)
}

func (l Logger) Errorf(format string, v ...interface{}) {
	if len(v) > 0 {
		format = fmt.Sprintf(format, v...)
	}

	logit.Error(format)
}
