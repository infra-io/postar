// Copyright 2020 Ye Zi Jie. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/13 23:23:24

package system

import (
	"sync"

	"gopkg.in/gomail.v2"
)

var (
	emailDialer         *gomail.Dialer
	emailDialerInitOnce = &sync.Once{}
)

// initEmailDialerWith 可以初始化 emailDialer，并且保证多次调用也只初始化一次。
func initEmailDialerWith(host string, port int, username string, password string) {
	emailDialerInitOnce.Do(func() {
		emailDialer = gomail.NewDialer(host, port, username, password)
	})
}
