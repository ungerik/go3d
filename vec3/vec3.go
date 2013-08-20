package vec3

import (
	"fmt"
	"math"

	"github.com/barnex/fmath"
	"github.com/ungerik/go3d/generic"
)

var (
	Zero = T{}

	UnitX = T{1, 0, 0}
	UnitY = T{0, 1, 0}
	UnitZ = T{0, 0, 1}

	Red   = T{1, 0, 0}
	Green = T{0, 1, 0}
	Blue  = T{0, 0, 1}
	Black = T{0, 0, 0}

	MinVal = T{-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32}
	MaxVal = T{+math.MaxFloat32, +math.MaxFloat32, +math.MaxFloat32}
)

type T [3]float32

func From(other generic.T) T {
	switch other.Size() {
	case 2:
		return T{other.Get(0, 0), other.Get(0, 1), 0}
	case 3, 4:
		return T{other.Get(0, 0), other.Get(0, 1), other.Get(0, 2)}
	default:
		panic("Unsupported type")
	}
}

func Parse(s string) (r T, err error) {
	_, err = fmt.Sscanf(s, "%f %f %f", &r[0], &r[1], &r[2])
	return r, err
}

func (self *T) String() string {
	return fmt.Sprintf("%f %f %f", self[0], self[1], self[2])
}

func (self *T) Rows() int {
	return 3
}

func (self *T) Cols() int {
	return 1
}

func (self *T) Size() int {
	return 3
}

func (self *T) Slice() []float32 {
	return []float32{self[0], self[1], self[2]}
}

func (self *T) Get(col, row int) float32 {
	return self[row]
}

func (self *T) IsZero() bool {
	return self[0] == 0 && self[1] == 0 && self[2] == 0
}

func (self *T) Length() float32 {
	return float32(fmath.Sqrt(self.LengthSqr()))
}

func (self *T) LengthSqr() float32 {
	return self[0]*self[0] + self[1]*self[1] + self[2]*self[2]
}

func (self *T) Scale(f float32) {
	self[0] *= f
	self[1] *= f
	self[2] *= f
}

func (self *T) Scaled(f float32) T {
	return T{self[0] * f, self[1] * f, self[2] * f}
}

func (self *T) Invert() {
	self[0] = -self[0]
	self[1] = -self[1]
	self[2] = -self[2]
}

func (self *T) Inverted() T {
	return T{-self[0], -self[1], -self[2]}
}

func (self *T) Normalize() {
	sl := self.LengthSqr()
	if sl == 0 || sl == 1 {
		return
	}
	self.Scale(1 / fmath.Sqrt(sl))
}

func (self *T) Normalized() T {
	v := *self
	v.Normalize()
	return v
}

func (self *T) Normal() T {
	n := Cross(self, &UnitZ)
	if n.IsZero() {
		return UnitX
	}
	return n.Normalized()
}

func (self *T) Add(v *T) {
	self[0] += v[0]
	self[1] += v[1]
	self[2] += v[2]
}

func (self *T) Sub(v *T) {
	self[0] -= v[0]
	self[1] -= v[1]
	self[2] -= v[2]
}

func (self *T) Mul(v *T) {
	self[0] *= v[0]
	self[1] *= v[1]
	self[2] *= v[2]
}

func Add(a, b *T) T {
	return T{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func Sub(a, b *T) T {
	return T{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func Mul(a, b *T) T {
	return T{a[0] * b[0], a[1] * b[1], a[2] * b[2]}
}

func Dot(a, b *T) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func Cross(a, b *T) T {
	return T{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

func Angle(a, b *T) float32 {
	return fmath.Acos(Dot(a, b))
}

func Min(a, b *T) T {
	min := *a
	if b[0] < min[0] {
		min[0] = b[0]
	}
	if b[1] < min[1] {
		min[1] = b[1]
	}
	if b[2] < min[2] {
		min[2] = b[2]
	}
	return min
}

func Max(a, b *T) T {
	max := *a
	if b[0] > max[0] {
		max[0] = b[0]
	}
	if b[1] > max[1] {
		max[1] = b[1]
	}
	if b[2] > max[2] {
		max[2] = b[2]
	}
	return max
}
