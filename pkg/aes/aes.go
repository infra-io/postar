// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package aes

import (
	"github.com/FishGoddess/cryptox"
	"github.com/FishGoddess/cryptox/aes"
)

func Encrypt(aesKey string, aesIV string, str string) (encrypted string, err error) {
	raw := cryptox.FromString(str)
	aes := aes.New(cryptox.FromString(aesKey))
	iv := cryptox.FromString(aesIV)

	bs, err := aes.EncryptCTR(cryptox.PaddingNone, iv, raw)
	if err != nil {
		return "", err
	}

	return bs.Base64(), nil
}

func Decrypt(aesKey string, aesIV string, str string) (decrypted string, err error) {
	raw, err := cryptox.FromBase64(str)
	if err != nil {
		return "", err
	}

	aes := aes.New(cryptox.FromString(aesKey))
	iv := cryptox.FromString(aesIV)

	bs, err := aes.DecryptCTR(cryptox.UnPaddingNone, iv, raw)
	if err != nil {
		return "", err
	}

	return bs.String(), nil
}
