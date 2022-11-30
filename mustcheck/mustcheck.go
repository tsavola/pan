// Copyright (c) 2022 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

// Package mustcheck will export only symbols prefixed with words Check or
// Must.
package mustcheck

import (
	"import.name/pan"
)

// Check is same as pan.Check.
func Check(err error) {
	pan.Check(err)
}

// Must is same as pan.Must.
func Must[T any](x T, err error) T {
	return pan.Must(x, err)
}

// Must2 is same as pan.Must2.
func Must2[T1 any, T2 any](x1 T1, x2 T2, err error) (T1, T2) {
	return pan.Must2(x1, x2, err)
}
