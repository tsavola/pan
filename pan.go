// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pan can be used to implement internal error propagation via
// panic and recover.  The benefit over naive panic/recover usage is that
// runtime errors (and other error values passed to panic) are propagated
// as usual.
//
//
// Example
//
//     import "import.name/pan"
//
//     func check(err error)                 { pan.Check(err) }
//     func checked[T any](x T, err error) T { pan.Check(err); return x }
//
//     func internal() string {
//         check(os.Chdir("/nonexistent"))
//         return checked(os.Getwd())
//     }
//
//     func Public() (s string, err error) {
//         err = pan.Recover(func() {
//             s = internal()
//         })
//         return
//     }
//
package pan

import (
	"fmt"
	"os"
)

type wrapper struct {
	err error
}

func (w wrapper) Error() string { return w.err.Error() }
func (w wrapper) Unwrap() error { return w.err }

// Check panics unless err is nil.
func Check(err error) {
	if err != nil {
		panic(wrapper{err})
	}
}

// Recover invokes f and returns any error value passed to Check.
func Recover(f func()) (err error) {
	defer func() { err = Error(recover()) }()
	f()
	return
}

// Error returns an error if x is a panic value from Check.  If x is nil,
// nil is returned.  If x is something else, Error panics with x as the
// panic value.
func Error(x interface{}) error {
	if x == nil {
		return nil
	}
	if w, ok := x.(wrapper); ok {
		return w.err
	}
	panic(x)
}

// Fatal is like Error, but the error is written to stderr and the program
// terminates with exit status 1.
func Fatal(x interface{}) {
	if err := Error(x); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
