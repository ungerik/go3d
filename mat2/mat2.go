// Package mat2 contains a 2x2 float32 matrix type T and functions.
package mat2

import (
	"errors"
	"fmt"

	math "github.com/chewxy/math32"
	"github.com/ungerik/go3d/generic"
	"github.com/ungerik/go3d/vec2"
)

var (
	// Zero holds a zero matrix.
	Zero = T{
		vec2.T{0, 0},
		vec2.T{0, 0},
	}

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
	_, err = fmt.Sscan(s,
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
// The data may be a copy depending on the platform implementation.
func (mat *T) Slice() []float32 {
	return mat.Array()[:]
}

// Get returns one element of the matrix.
// Matrices are defined by (two) column vectors.
//
// Note that this function use the opposite reference order of rows and columns to the mathematical matrix indexing.
//
// A value in this matrix is referenced by <col><row> where both row and column is in the range [0..1].
// This notation and range reflects the underlying representation.
//
// A value in a matrix A is mathematically referenced by A<row><col>
// where both row and column is in the range [1..2].
// (It is really the lower case 'a' followed by <row><col> but this documentation syntax is somewhat limited.)
//
// matrixA.Get(0, 1) == matrixA[0][1] ( == A21 in mathematical indexing)
func (mat *T) Get(col, row int) float32 {
	return mat[col][row]
}

// IsZero checks if all elements of the matrix are exactly zero.
// Uses exact equality comparison, which may not be suitable for floating-point math results.
// For tolerance-based comparison, use IsZeroEps instead.
func (mat *T) IsZero() bool {
	return *mat == Zero
}

// IsZeroEps checks if all elements of the matrix are zero within the given epsilon tolerance.
// This is the recommended method for comparing floating-point matrices that result from calculations.
// For exact zero comparison, use IsZero instead.
func (mat *T) IsZeroEps(epsilon float32) bool {
	return math.Abs(mat[0][0]) <= epsilon && math.Abs(mat[0][1]) <= epsilon &&
		math.Abs(mat[1][0]) <= epsilon && math.Abs(mat[1][1]) <= epsilon
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

// TransformVec2 multiplies v with mat and saves the result in v.
func (mat *T) TransformVec2(v *vec2.T) {
	// Use intermediate variables to not alter further computations.
	x := mat[0][0]*v[0] + mat[1][0]*v[1]
	v[1] = mat[0][1]*v[0] + mat[1][1]*v[1]
	v[0] = x
}

func (mat *T) Determinant() float32 {
	return mat[0][0]*mat[1][1] - mat[1][0]*mat[0][1]
}

// PracticallyEquals compares two matrices if they are equal with each other within a delta tolerance.
func (mat *T) PracticallyEquals(matrix *T, allowedDelta float32) bool {
	return mat[0].PracticallyEquals(&matrix[0], allowedDelta) &&
		mat[1].PracticallyEquals(&matrix[1], allowedDelta)
}

// Transpose transposes the matrix.
func (mat *T) Transpose() *T {
	mat[0][1], mat[1][0] = mat[1][0], mat[0][1]
	return mat
}

// Transposed returns a transposed copy the matrix.
func (mat *T) Transposed() T {
	result := *mat
	result.Transpose()
	return result
}

// Invert inverts the given matrix. Destructive operation.
// Does not check if matrix is singular and may lead to strange results!
func (mat *T) Invert() (*T, error) {
	determinant := mat.Determinant()
	if determinant == 0 {
		return &Zero, errors.New("can not create inverted matrix as determinant is 0")
	}

	invDet := 1.0 / determinant

	mat[0][0], mat[1][1] = invDet*mat[1][1], invDet*mat[0][0]
	mat[0][1] = -invDet * mat[0][1]
	mat[1][0] = -invDet * mat[1][0]

	return mat, nil
}

// Inverted inverts a copy of the given matrix.
// Does not check if matrix is singular and may lead to strange results!
func (mat *T) Inverted() (T, error) {
	result := *mat
	_, err := result.Invert()
	return result, err
}
