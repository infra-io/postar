// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package runtime

import (
	"fmt"
	"testing"
)

// BenchmarkCallers-2   	  437528	      2584 ns/op	     748 B/op	      12 allocs/op
func BenchmarkCallers(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Callers()
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestCallers$
func TestCallers(t *testing.T) {
	callers := Callers()
	fmt.Println(callers)
}
