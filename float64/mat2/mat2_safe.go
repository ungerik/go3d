// +build netgo

package mat2

// Array returns the elements of the matrix as array pointer.
// The data may be a copy depending on the platform implementation.
func (mat *T) Array() *[4]float64 {
	return &[...]float64{
		mat[0][0], mat[0][1],
		mat[1][0], mat[1][1],
	}
}
