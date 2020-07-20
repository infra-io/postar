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

	"github.com/FishGoddess/logit"
	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/server"
)

func main() {
	startPostar()
}

// printSymbol prints the symbol of postar.
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

// printBootingInformation prints the booting information of postar.
func printBootingInformation() {
	if core.Logger().Level() <= logit.InfoLevel {
		core.Logger().Infof("Postar %s is booting, please wait a moment...", core.Version)
		return
	}
	fmt.Printf("Postar %s is booting, please wait a moment...\n", core.Version)
}

// printReadyInformation prints the ready information of postar.
func printReadyInformation(timeSpent time.Duration) {
	if core.Logger().Level() <= logit.InfoLevel {
		core.Logger().Infof("Postar is ready! It takes %s to boot it.\n", timeSpent)
		return
	}
	fmt.Printf("Postar is ready! It takes %s to boot it.\n", timeSpent)
}

// startPostar starts postar services.
func startPostar() {

	// Before booting.
	printSymbol()
	printBootingInformation()

	// Boot and record the time it takes.
	beginTime := time.Now()
	wg := server.RunServer()
	endTime := time.Now()

	// After booting.
	printReadyInformation(endTime.Sub(beginTime))
	wg.Wait()
}

// shutdownPostar shutdowns postar services.
func shutdownPostar() {

}
