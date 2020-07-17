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
// Created at 2020/07/08 23:41:46
package main

import (
	"fmt"
	"time"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/server"
)

func main() {

	// Before booting.
	printSymbol()
	core.Logger().Infof("Postar %s is booting, please wait a moment...", core.Version)

	// Boot and record the time it takes.
	beginTime := time.Now()
	wg := server.RunServer()
	endTime := time.Now()

	// After booting.
	core.Logger().Infof("Postar is ready! It takes %s to boot it.", endTime.Sub(beginTime))
	wg.Wait()
}

func printSymbol() {
	fmt.Println(`
*******************************************************************
*   _____      ____      _____   ________     ____     ______     *
*  (  __ \    / __ \    / ____\ (___  ___)   (    )   (   __ \    *
*   ) )_) )  / /  \ \  ( (___       ) )      / /\ \    ) (__) )   *
*  (  ___/  ( ()  () )  \___ \     ( (      ( (__) )  (    __/    *
*   ) )     ( ()  () )      ) )     ) )      )    (    ) \ \  _   *
*  ( (       \ \__/ /   ___/ /     ( (      /  /\  \  ( ( \ \_))  *
*  /__\       \____/   /____/      /__\    /__(  )__\  )_) \__/   *
*                                                                 *
*******************************************************************`)
}
