// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/26 21:05:47

package based

import (
	"sync"

	"github.com/avino-plan/postar/src/core"
)

// BasedServer is the based server for implements.
type BasedServer struct {

	// wg is for waiting these servers.
	wg *sync.WaitGroup
}

// InitServer initializes servers with given two ports.
// We should pass the sub struct's two functions to Init because Go doesn't have inherit mechanism.
func (bs *BasedServer) Init(initServerForService func(port string, beforeServing func(), cleanUp func()), initServerForShutdown func(port string, cleanUp func()), port string, closedPort string) *sync.WaitGroup {

	// Create a wait group to wait these servers.
	bs.wg = &sync.WaitGroup{}

	// Notice that wg.Add must be executed before wg.Done, so they can't code in go func.
	bs.wg.Add(1)
	initServerForService(
		port,
		func() {
			bs.wg.Add(1)
			initServerForShutdown(closedPort, func() {
				core.Logger().Debug("Add 1 to wg in initServerForShutdown...")
				bs.wg.Done()
			})
		},
		func() {
			core.Logger().Debug("Add 1 to wg in initServerForService...")
			bs.wg.Done()
		},
	)

	core.Logger().Infof("The main service is using port %s.", port)
	core.Logger().Infof("The closed service is using port %s.", closedPort)
	return bs.wg
}
