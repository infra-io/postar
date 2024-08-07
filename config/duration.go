// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import "time"

type Duration time.Duration

func (d *Duration) Standard() time.Duration {
	return time.Duration(*d)
}

func (d *Duration) UnmarshalText(text []byte) error {
	parsed, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}

	*d = Duration(parsed)
	return nil
}
