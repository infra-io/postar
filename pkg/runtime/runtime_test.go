// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package runtime

import "testing"

// go test -v -cover -count=1 -test.cpu=1 -run=^TestStack$
func TestStack(t *testing.T) {
	stack := Stack()
	t.Log(stack)
}
