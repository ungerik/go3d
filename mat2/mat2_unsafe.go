//go:build !netgo
// +build !netgo

package mat2

import "unsafe"

// Array returns the elements of the matrix as array pointer.
// The data may be a copy depending on the platform implementation.
func (mat *T) Array() *[4]float32 {
	return (*[4]float32)(unsafe.Pointer(mat)) //#nosec G103 -- unsafe OK
}
