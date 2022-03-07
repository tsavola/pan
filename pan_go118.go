// Copyright (c) 2022 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

package pan

// Must panics if err is not nil.  Otherwise it returns x.  See Error, Fatal
// and Recover.
func Must[T any](x T, err error) T {
	Check(err)
	return x
}
