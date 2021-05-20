// Copyright 2011 Arne Vansteenkiste.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// func Sqrtf(x float32) float32
TEXT ·Sqrtf+0(SB),$0-16
	SQRTSS	 x+0(FP), X0
	MOVSS   X0,r+8(FP)
	RET
