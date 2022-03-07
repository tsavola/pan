// Copyright (c) 2022 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

// Package must will export only symbols prefixed with word Must.
package must

import (
	"import.name/pan"
)

// Must is same as pan.Must.
func Must[T any](x T, err error) T {
	return pan.Must(x, err)
}
