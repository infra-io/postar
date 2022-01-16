// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 22:39:17

package encode

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	// letters includes 0-9, a-z, A-Z.
	letters = [62]byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}

	pid    = strconv.Itoa(os.Getpid())
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

// numberHex returns num in hex string.
// The hex string will be cut with start and end.
func numberHex(num uint64, start int, end int) string {
	size := 8
	if start < 0 || start > size {
		start = 0
	}

	if end <= 0 || end > size {
		end = size
	}

	b := make([]byte, size)
	binary.BigEndian.PutUint64(b, num)
	return fmt.Sprintf("%x", b[start:end])
}

// Now returns in current time in hex string.
func Now() string {
	return numberHex(uint64(time.Now().Unix()), 4, 0)
}

// PID returns pid in string.
func PID() string {
	return pid
}

// RandomString returns a string including 0-9/a-z/A-Z not longer than length.
func RandomString(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = letters[random.Intn(62)]
	}
	return string(b)
}
