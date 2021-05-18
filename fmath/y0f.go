// Copyright 2011 Arne Vansteenkiste. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides a trivial implementation based on
// Go's float64 math library. It may be overridden by an
// assembly implementation when available for the platform.

package fmath

import "math"

// float32 version of math.Y0
func Y0(x float32) float32 {
	return float32(math.Y0(float64(x)))
}
