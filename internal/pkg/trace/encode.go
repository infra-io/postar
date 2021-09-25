// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/09/25 22:39:17

package trace

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	// letters includes 0-9 a-z A-Z.
	letters = [62]byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}

	// random 用于生成随机数
	random = rand.New(rand.NewSource(time.Now().Unix()))

	// pid 是程序的 pid
	pid = uint64(os.Getpid())
)

// TimeHex returns now time in hex.
func TimeHex() string {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(time.Now().Unix()))
	return fmt.Sprintf("%x", b[4:])
}

// PidHex returns pid in hex.
func PidHex() string {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, pid)
	return fmt.Sprintf("%x", b[4:])
}

// RandomString returns a string including 0-9/a-z/A-Z not longer than length.
func RandomString(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = letters[random.Intn(62)]
	}
	return string(b)
}
