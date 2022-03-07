// Copyright (c) 2022 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package check will export only symbols prefixed with word Check.
package check

import (
	"import.name/pan"
)

// Check is same as pan.Check.
func Check(err error) {
	pan.Check(err)
}
