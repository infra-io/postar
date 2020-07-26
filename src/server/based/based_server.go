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
	Wg *sync.WaitGroup
}

// initServerForService initializes the server for service.
func (bs *BasedServer) initServerForService(port string, beforeServing func(), cleanUp func()) {
	// Implement me!
}

// initServerForShutdown initializes the server for shutdown.
func (bs *BasedServer) initServerForShutdown(port string, cleanUp func()) {
	// Implement me!
}

// InitServer initializes servers with given two ports.
func (bs *BasedServer) Init(port string, closedPort string) *sync.WaitGroup {

	// Create a wait group to wait these servers.
	bs.Wg = &sync.WaitGroup{}

	// Notice that wg.Add must be executed before wg.Done, so they can't code in go func.
	bs.Wg.Add(1)
	bs.initServerForService(
		port,
		func() {
			bs.Wg.Add(1)
			bs.initServerForShutdown(closedPort, func() {
				core.Logger().Debug("Add 1 to wg in initServerForShutdown...")
				bs.Wg.Done()
			})
		},
		func() {
			core.Logger().Debug("Add 1 to wg in initServerForService...")
			bs.Wg.Done()
		},
	)

	core.Logger().Infof("The main service is using port %s.", port)
	core.Logger().Infof("The closed service is using port %s.", closedPort)
	return bs.Wg
}

// StopServer stops the running servers.
func (bs *BasedServer) Stop(closedPort string) error {
	// Implement me!
	return nil
}
