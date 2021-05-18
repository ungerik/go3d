// Copyright 2011 Arne Vansteenkiste. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides a trivial implementation based on
// Go's float64 math library. It may be overridden by an
// assembly implementation when available for the platform.

package fmath

import "math"

// float32 version of math.Sincos
func Sincos(x float32) (sin, cos float32) {
	sin64, cos64 := math.Sincos(float64(x))
	sin = float32(sin64)
	cos = float32(cos64)
	return
}
