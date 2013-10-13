// Package generic contains an interface T that
// that all float32 vector and matrix types implement.
package generic

// T is an interface that all float32 vector and matrix types implement.
type T interface {

	// Cols returns the number of columns of the vector or matrix.
	Cols() int

	// Rows returns the number of rows of the vector or matrix.
	Rows() int

	// Size returns the number elements of the vector or matrix.
	Size() int

	// Slice returns the elements of the vector or matrix as slice.
	Slice() []float32

	// Get returns one element of the vector or matrix.
	Get(row, col int) float32

	// IsZero checks if all elements of the vector or matrix are zero.
	IsZero() bool
}
