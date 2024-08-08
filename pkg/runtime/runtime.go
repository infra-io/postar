// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package runtime

import (
	"runtime"
	"unsafe"

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

const (
	maxStackSize = 4096 // 4KB
)

func Stack() string {
	stack := make([]byte, maxStackSize)
	n := runtime.Stack(stack, false)
	bs := stack[:n]

	bsPtr := unsafe.SliceData(bs)
	return unsafe.String(bsPtr, len(bs))
}
