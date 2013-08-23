package mat2x2d

import (
	"fmt"
	"github.com/ungerik/go3d/vec2d"
)

var (
	Zero  = T{}
	Ident = T{
		vec2d.T{1, 0},
		vec2d.T{0, 1},
	}
)

type T [2]vec2d.T

func Parse(s string) (r T, err error) {
	_, err = fmt.Sscanf(s,
		"%f %f %f %f",
		&r[0][0], &r[0][1],
		&r[1][0], &r[1][1],
	)
	return r, err
}

func (self *T) String() string {
	return fmt.Sprintf("%s %s", self[0].String(), self[1].String())
}

func (self *T) Rows() int {
	return 2
}

func (self *T) Cols() int {
	return 2
}

func (self *T) Size() int {
	return 4
}

func (self *T) Slice() []float64 {
	return []float64{
		self[0][0], self[0][1],
		self[1][0], self[1][1],
	}
}

func (self *T) Get(col, row int) float64 {
	return self[col][row]
}

func (self *T) Trace() float64 {
	return self[0][0] + self[1][1]
}

func (self *T) AssignMul(a, b *T) *T {
	self[0] = a.MulVec2(&b[0])
	self[1] = a.MulVec2(&b[1])
	return self
}

func (self *T) MulVec2(vec *vec2d.T) vec2d.T {
	return vec2d.T{
		self[0][0]*vec[0] + self[1][0]*vec[1],
		self[0][1]*vec[1] + self[1][1]*vec[1],
	}
}
