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

	// InitServerForService is the function for initializing main service.
	InitServerForService func(port string, beforeServing func(), cleanUp func())

	// InitServerForService is the function for initializing close service.
	InitServerForShutdown func(port string, cleanUp func())
}

// InitServer initializes servers with given two ports.
// We should pass the sub struct's two functions to Init because Go doesn't have inherit mechanism.
func (bs *BasedServer) Init(port string, closedPort string) *sync.WaitGroup {

	// Create a wait group to wait these servers.
	bs.wg = &sync.WaitGroup{}

	// Notice that wg.Add must be executed before wg.Done, so they can't code in go func.
	bs.wg.Add(1)
	bs.InitServerForService(
		port,
		func() {
			bs.wg.Add(1)
			bs.InitServerForShutdown(closedPort, func() {
				bs.wg.Done()
				core.Logger().Debug("Done with wg in initServerForShutdown...")
			})
		},
		func() {
			bs.wg.Done()
			core.Logger().Debug("Done with wg in initServerForService...")
		},
	)

	core.Logger().Infof("The main service is using port %s.", port)
	core.Logger().Infof("The closed service is using port %s.", closedPort)
	return bs.wg
}

// StopServer stops the running servers.
func (bs *BasedServer) Stop(closedPort string) error {
	return nil
}
