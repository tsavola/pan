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

// Must2 panics if err is not nil.  Otherwise it returns x1 and x2.  See Error,
// Fatal and Recover.
func Must2[T1 any, T2 any](x1 T1, x2 T2, err error) (T1, T2) {
	Check(err)
	return x1, x2
}
