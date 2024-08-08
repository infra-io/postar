// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"time"

	"github.com/FishGoddess/cachego"
	"github.com/FishGoddess/logit"
)

func newCache(name string) cachego.Cache {
	opts := []cachego.Option{
		cachego.WithCacheName(name),
		cachego.WithShardings(4),
		cachego.WithLFU(10000),
		cachego.WithGC(5 * time.Minute),
		cachego.WithRecordLoad(false),
		cachego.WithReportGC(reportCacheGC),
	}

	cache, _ := cachego.NewCacheWithReport(opts...)
	return cache
}

func reportCacheGC(reporter *cachego.Reporter, cost time.Duration, cleans int) {
	name := reporter.CacheName()
	size := reporter.CacheSize()
	hitRate := reporter.HitRate()

	logit.Info("report cache gc", "name", name, "cost", cost, "cleans", cleans, "size", size, "hit_rate", hitRate)
}
