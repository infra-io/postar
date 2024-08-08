// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package trace

import (
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestID$
func TestID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(ID())
	}
}

// go test -v -run=^$ -bench=^BenchmarkID$ -benchtime=1s
// BenchmarkID-2           9329845               128.8 ns/op            24 B/op          1 allocs/op
func BenchmarkID(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ID()
	}
}
