// +build !netgo

package mat4

import "unsafe"

// Array returns the elements of the matrix as array pointer.
// The data may be a copy depending on the platform implementation.
func (mat *T) Array() *[16]float64 {
	return (*[16]float64)(unsafe.Pointer(mat))
}
