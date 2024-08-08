// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package trace

import (
	"strconv"
	"unsafe"

	"github.com/FishGoddess/cryptox"
	"github.com/FishGoddess/logit/extension/fastclock"
)

func now() int64 {
	now := fastclock.NowNanos()
	now = now / 1000_000_000

	return now
}

func ID() string {
	bs := make([]byte, 0, 24)
	bs = strconv.AppendInt(bs, now(), 16) // 8 bytes appended before 2106-02-07.
	bs = cryptox.AppendBytes(bs, 16)

	bsPtr := unsafe.SliceData(bs)
	return unsafe.String(bsPtr, len(bs))
}
