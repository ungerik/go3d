//go:build !netgo

package mat3

import "unsafe"

// Array returns the elements of the matrix as array pointer.
// The data may be a copy depending on the platform implementation.
func (mat *T) Array() *[9]float64 {
	return (*[9]float64)(unsafe.Pointer(mat)) //#nosec G103 -- unsafe OK
}
