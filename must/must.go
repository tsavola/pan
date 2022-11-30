// Copyright (c) 2022 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package must will export only symbols prefixed with word Must.
package must

import (
	"import.name/pan"
)

// Must is same as pan.Must.
func Must[T any](x T, err error) T {
	return pan.Must(x, err)
}

// Must2 is same as pan.Must2.
func Must2[T1, T2 any](x1 T1, x2 T2, err error) (T1, T2) {
	return pan.Must2(x1, x2, err)
}

// Must3 is same as pan.Must3.
func Must3[T1, T2, T3 any](x1 T1, x2 T2, x3 T3, err error) (T1, T2, T3) {
	return pan.Must3(x1, x2, x3, err)
}
