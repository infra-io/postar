// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package aes

import (
	"github.com/FishGoddess/cryptox"
	"github.com/FishGoddess/cryptox/aes"
)

func Encrypt(aesKey string, aesIV string, str string) (encrypted string, err error) {
	key := []byte(aesKey)
	iv := []byte(aesIV)
	bs := []byte(str)

	bs, err = aes.EncryptCTR(key, iv, cryptox.PaddingNone, bs)
	if err != nil {
		return "", err
	}

	encrypted = cryptox.Bytes(bs).Base64()
	return encrypted, nil
}

func Decrypt(aesKey string, aesIV string, str string) (decrypted string, err error) {
	key := []byte(aesKey)
	iv := []byte(aesIV)

	bs, err := cryptox.ParseBase64(str)
	if err != nil {
		return "", err
	}

	bs, err = aes.DecryptCTR(key, iv, cryptox.PaddingNone, bs)
	if err != nil {
		return "", err
	}

	decrypted = string(bs)
	return decrypted, nil
}
