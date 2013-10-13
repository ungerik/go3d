// Package mat2x2 contains a 2x2 float32 matrix type T and functions.
package mat2x2

import (
	"fmt"

	"github.com/ungerik/go3d/generic"
	"github.com/ungerik/go3d/vec2"
)

var (
	// Zero holds a zero matrix.
	Zero  = T{}
	
	// Ident holds an ident matrix.
	Ident = T{
		vec2.T{1, 0},
		vec2.T{0, 1},
	}
)

// T represents a 2x2 matrix.
type T [2]vec2.T

// From copies a T from a generic.T implementation.
func From(other generic.T) T {
	r := Ident
	cols := other.Cols()
	rows := other.Rows()
	if (cols == 3 && rows == 3) || (cols == 4 && rows == 4) {
		cols = 2
		rows = 2
	} else if !(cols == 2 && rows == 2) {
		panic("Unsupported type")
	}
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			r[col][row] = other.Get(col, row)
		}
	}
	return r
}

// Parse parses T from a string. See also String()
func Parse(s string) (r T, err error) {
	_, err = fmt.Sscanf(s,
		"%f %f %f %f",
		&r[0][0], &r[0][1],
		&r[1][0], &r[1][1],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (mat *T) String() string {
	return fmt.Sprintf("%s %s", mat[0].String(), mat[1].String())
}

// Rows returns the number of rows of the matrix.
func (mat *T) Rows() int {
	return 2
}

// Cols returns the number of columns of the matrix.
func (mat *T) Cols() int {
	return 2
}

// Size returns the number elements of the matrix.
func (mat *T) Size() int {
	return 4
}

// Slice returns the elements of the matrix as slice.
func (mat *T) Slice() []float32 {
	return []float32{
		mat[0][0], mat[0][1],
		mat[1][0], mat[1][1],
	}
}

// Get returns one element of the matrix.
func (mat *T) Get(col, row int) float32 {
	return mat[col][row]
}

// IsZero checks if all elements of the matrix are zero.
func (mat *T) IsZero() bool {
	return *mat == Zero
}

// Scale multiplies the diagonal scale elements by f returns mat.
func (mat *T) Scale(f float32) *T {
	mat[0][0] *= f
	mat[1][1] *= f
	return mat
}

// Scaled returns a copy of the matrix with the diagonal scale elements multiplied by f.
func (mat *T) Scaled(f float32) T {
	r := *mat
	return *r.Scale(f)
}

// Scaling returns the scaling diagonal of the matrix.
func (mat *T) Scaling() vec2.T {
	return vec2.T{mat[0][0], mat[1][1]}
}

// SetScaling sets the scaling diagonal of the matrix.
func (mat *T) SetScaling(s *vec2.T) *T {
	mat[0][0] = s[0]
	mat[1][1] = s[1]
	return mat
}

// Trace returns the trace value for the matrix.
func (mat *T) Trace() float32 {
	return mat[0][0] + mat[1][1]
}

// AssignMul multiplies a and b and assigns the result to mat.
func (mat *T) AssignMul(a, b *T) *T {
	mat[0] = a.MulVec2(&b[0])
	mat[1] = a.MulVec2(&b[1])
	return mat
}

// MulVec2 multiplies vec with mat.
func (mat *T) MulVec2(vec *vec2.T) vec2.T {
	return vec2.T{
		mat[0][0]*vec[0] + mat[1][0]*vec[1],
		mat[0][1]*vec[1] + mat[1][1]*vec[1],
	}
}

// Transpose transposes the matrix.
func (mat *T) Transpose() *T {
	temp := mat[0][1]
	mat[0][1] = mat[1][0]
	mat[1][0] = temp
	return mat
}
