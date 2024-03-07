// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package aes

import (
	"testing"

	"github.com/infra-io/servicex/rand"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNew$
func TestAES(t *testing.T) {
	str := rand.GenerateString(64)

	aesKey := "123456788765432112345678"
	aesIV := "1234567887654321"

	encrypted, err := Encrypt(aesKey, aesIV, str)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(str), str)
	t.Log(len(encrypted), encrypted)

	decrypted, err := Decrypt(aesKey, aesIV, encrypted)
	if err != nil {
		t.Fatal(err)
	}

	if decrypted != str {
		t.Fatalf("decrypted %s != str %s", decrypted, str)
	}
}
