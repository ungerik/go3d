package mat2x2

import (
	"fmt"

	"github.com/ungerik/go3d/float64/generic"
	"github.com/ungerik/go3d/float64/vec2"
)

var (
	Zero  = T{}
	Ident = T{
		vec2.T{1, 0},
		vec2.T{0, 1},
	}
)

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
func (self *T) String() string {
	return fmt.Sprintf("%s %s", self[0].String(), self[1].String())
}

// Rows returns the number of rows of the matrix.
func (self *T) Rows() int {
	return 2
}

// Cols returns the number of columns of the matrix.
func (self *T) Cols() int {
	return 2
}

// Size returns the number elements of the matrix.
func (self *T) Size() int {
	return 4
}

// Slice returns the elements of the matrix as slice.
func (self *T) Slice() []float64 {
	return []float64{
		self[0][0], self[0][1],
		self[1][0], self[1][1],
	}
}

// Get returns one element of the matrix.
func (self *T) Get(col, row int) float64 {
	return self[col][row]
}

// IsZero checks if all elements of the matrix are zero.
func (self *T) IsZero() bool {
	return *self == Zero
}

// Scale multiplies the diagonal scale elements by f returns self.
func (self *T) Scale(f float64) *T {
	self[0][0] *= f
	self[1][1] *= f
	return self
}

// Scaled returns a copy of the matrix with the diagonal scale elements multiplied by f.
func (self *T) Scaled(f float64) T {
	r := *self
	return *r.Scale(f)
}

// Scaling returns the scaling diagonal of the matrix.
func (self *T) Scaling() vec2.T {
	return vec2.T{self[0][0], self[1][1]}
}

// SetScaling sets the scaling diagonal of the matrix.
func (self *T) SetScaling(s *vec2.T) *T {
	self[0][0] = s[0]
	self[1][1] = s[1]
	return self
}

// Trace returns the trace value for the matrix.
func (self *T) Trace() float64 {
	return self[0][0] + self[1][1]
}

// AssignMul multiplies a and b and assigns the result to self.
func (self *T) AssignMul(a, b *T) *T {
	self[0] = a.MulVec2(&b[0])
	self[1] = a.MulVec2(&b[1])
	return self
}

// MulVec2 multiplies vec with self.
func (self *T) MulVec2(vec *vec2.T) vec2.T {
	return vec2.T{
		self[0][0]*vec[0] + self[1][0]*vec[1],
		self[0][1]*vec[1] + self[1][1]*vec[1],
	}
}
