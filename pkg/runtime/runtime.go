// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package runtime

import (
	"runtime"
	"strconv"

	"github.com/FishGoddess/logit"
	"go.uber.org/automaxprocs/maxprocs"
)

func init() {
	undo, err := maxprocs.Set(maxprocs.Logger(logit.Printf))
	if err != nil {
		logit.Error("set maxprocs failed", "err", err)
		undo()
	}
}

func Callers() []string {
	pcs := make([]uintptr, 16)
	n := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	var callers []string
	for {
		frame, more := frames.Next()

		caller := frame.File + ":" + strconv.Itoa(frame.Line)
		callers = append(callers, caller)

		if !more {
			break
		}
	}

	return callers
}
