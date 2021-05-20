// Copyright 2011 Arne Vansteenkiste. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// float32 math package.
//
// This library provides float32 counterparts for Go's float64 math functions. E.g.:
//		Sqrtf(x float32) float32.
//
// The implementation partially uses assembly code (see, e.g., sqrtf_amd64.s) for fast computation. However, when no assembly implementation exists yet a generic implementation is used which uses Go's float64 math functions underneath.
//
// Note: the assembly code is not goinstallable on all platforms. Therefore, goinstall will compile the portable implementation. If you manually execute "make install", you will get the faster implementation.
package fmath
