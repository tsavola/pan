// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pan (short for panic) can be used to implement internal error
// propagation via panic and recover.  A benefit over plain panic/recover usage
// is that unrelated panics are automatically disregarded.
//
// Example
//
//	import "import.name/pan"
//
//	var z = new(pan.Zone)
//
//	func must[T any](value T, err error) T {
//		z.Check(err)
//		return value
//	}
//
//	func changeDirectory(dir string) string {
//		z.Check(os.Chdir(dir))
//		return must(os.Getwd())
//	}
//
//	func ChangeDirectory(dir string) (string, error) {
//		return pan.Recover1(z, func() string {
//			return changeDirectory(dir)
//		})
//	}
package pan

import (
	"runtime"
)

type wrapper struct {
	err error
	z   *Zone
}

func (w wrapper) Error() string { return w.err.Error() }
func (w wrapper) Unwrap() error { return w.err }

// Zone for error propagation.  Allocate a private instance of Zone for your
// package or module.  A panic within a zone is not receovered within another
// zone.
type Zone struct{}

// Wrap an error.  panic(Wrap(err)) is equivalent to Panic(err), but can be
// used to work around compilation errors.
func (z *Zone) Wrap(err error) error {
	if err == nil {
		err = new(runtime.PanicNilError)
	}
	return wrapper{err, z}
}

// Panic unconditionally.
func (z *Zone) Panic(err error) {
	panic(z.Wrap(err))
}

// Check err and panic unless it's nil.
func (z *Zone) Check(err error) {
	if err != nil {
		panic(wrapper{err, z})
	}
}

// Error returns the original error if x is a wrapper value created by Check,
// Panic or Wrap within this zone.  If x is nil, nil is returned.  If x is
// something else, Error panics with x as the panic value.
func (z *Zone) Error(x any) error {
	if x == nil {
		return nil
	}
	if w, ok := x.(wrapper); ok && w.z == z {
		return w.err
	}
	panic(x)
}

// Recover invokes f, returning any error raised within this zone.
func (z *Zone) Recover(f func()) (err error) {
	defer func() { err = z.Error(recover()) }()
	f()
	return
}

// Recover1 invokes f, returning any error raised within this zone.
func Recover1[T any](z *Zone, f func() T) (x T, err error) {
	defer func() { err = z.Error(recover()) }()
	x = f()
	return
}

// Recover2 invokes f, returning any error raised within this zone.
func Recover2[T1, T2 any](z *Zone, f func() (T1, T2)) (x1 T1, x2 T2, err error) {
	defer func() { err = z.Error(recover()) }()
	x1, x2 = f()
	return
}

// Recover3 invokes f, returning any error raised within this zone.
func Recover3[T1, T2, T3 any](z *Zone, f func() (T1, T2, T3)) (x1 T1, x2 T2, x3 T3, err error) {
	defer func() { err = z.Error(recover()) }()
	x1, x2, x3 = f()
	return
}
