// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/08 23:41:46
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/FishGoddess/logit"
	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/server"
)

var (
	// systemCommands stores all systemCommands can be executed.
	systemCommands = map[string]func(){
		"boot": startPostar,
		"off":  shutdownPostar,
	}
)

func main() {
	systemCommand, ok := systemCommands[core.SystemCommand()]
	if !ok {
		fmt.Printf("Unsupported system command: %s. Try boot, off.\n", core.SystemCommand())
		os.Exit(1)
	}
	systemCommand()
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
	core.Logger().Debug("Exit wait status...")
}

// shutdownPostar shutdowns postar services.
func shutdownPostar() {
	server.ShutdownServer()
}
