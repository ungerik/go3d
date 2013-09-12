package vec2d

import (
	"fmt"
	"math"

	"github.com/ungerik/go3d/genericd"
)

var (
	Zero = T{}

	UnitX = T{1, 0}
	UnitY = T{0, 1}

	MinVal = T{-math.MaxFloat64, -math.MaxFloat64}
	MaxVal = T{+math.MaxFloat64, +math.MaxFloat64}
)

type T [2]float64

// From copies a T from a generic.T implementation.
func From(other genericd.T) T {
	return T{other.Get(0, 0), other.Get(0, 1)}
}

// Parse parses T from a string. See also String()
func Parse(s string) (r T, err error) {
	_, err = fmt.Sscanf(s, "%f %f", &r[0], &r[1])
	return r, err
}

// String formats T as string. See also Parse().
func (self *T) String() string {
	return fmt.Sprintf("%f %f", self[0], self[1])
}

// Rows returns the number of rows of the vector.
func (self *T) Rows() int {
	return 2
}

// Cols returns the number of columns of the vector.
func (self *T) Cols() int {
	return 1
}

// Size returns the number elements of the vector.
func (self *T) Size() int {
	return 2
}

// Slice returns the elements of the vector as slice.
func (self *T) Slice() []float64 {
	return []float64{self[0], self[1]}
}

// Get returns one element of the vector.
func (self *T) Get(col, row int) float64 {
	return self[row]
}

// IsZero checks if all elements of the vector are zero.
func (self *T) IsZero() bool {
	return self[0] == 0 && self[1] == 0
}

// Length returns the length of the vector.
// See also LengthSqr and Normalize.
func (self *T) Length() float64 {
	return float64(math.Sqrt(self.LengthSqr()))
}

// Length returns the squared length of the vector.
// See also Length and Normalize.
func (self *T) LengthSqr() float64 {
	return self[0]*self[0] + self[1]*self[1]
}

// Scale multiplies all element of the vector by f and returns self.
func (self *T) Scale(f float64) *T {
	self[0] *= f
	self[1] *= f
	return self
}

// Scaled returns a copy of self with all elements multiplies by f.
func (self *T) Scaled(f float64) T {
	return T{self[0] * f, self[1] * f}
}

func (self *T) Invert() *T {
	self[0] = -self[0]
	self[1] = -self[1]
	return self
}

func (self *T) Inverted() T {
	return T{-self[0], -self[1]}
}

func (self *T) Normalize() *T {
	sl := self.LengthSqr()
	if sl == 0 || sl == 1 {
		return self
	}
	return self.Scale(1 / math.Sqrt(sl))
}

func (self *T) Normalized() T {
	v := *self
	v.Normalize()
	return v
}

func (self *T) Add(v *T) *T {
	self[0] += v[0]
	self[1] += v[1]
	return self
}

func (self *T) Sub(v *T) *T {
	self[0] -= v[0]
	self[1] -= v[1]
	return self
}

func (self *T) Mul(v *T) *T {
	self[0] *= v[0]
	self[1] *= v[1]
	return self
}

func (self *T) Rotated(angle float64) T {
	sinus := math.Sin(angle)
	cosinus := math.Cos(angle)
	return T{
		self[0]*cosinus - self[1]*sinus,
		self[0]*sinus + self[1]*cosinus,
	}
}

func (self *T) Rotate(angle float64) *T {
	*self = self.Rotated(angle)
	return self
}

func (self *T) RotateAroundPoint(point *T, angle float64) *T {
	return self.Sub(point).Rotate(angle).Add(point)
}

func (self *T) Rotate90DegLeft() *T {
	temp := self[0]
	self[0] = -self[1]
	self[1] = temp
	return self
}

func (self *T) Rotate90DegRight() *T {
	temp := self[0]
	self[0] = self[1]
	self[1] = -temp
	return self
}

func (self *T) Angle() float64 {
	return math.Atan2(self[1], self[0])
}

func Add(a, b *T) T {
	return T{a[0] + b[0], a[1] + b[1]}
}

func Sub(a, b *T) T {
	return T{a[0] - b[0], a[1] - b[1]}
}

func Mul(a, b *T) T {
	return T{a[0] * b[0], a[1] * b[1]}
}

func Dot(a, b *T) float64 {
	return a[0]*b[0] + a[1]*b[1]
}

func Cross(a, b *T) T {
	return T{
		a[1]*b[0] - a[0]*b[1],
		a[0]*b[1] - a[1]*b[0],
	}
}

func Angle(a, b *T) float64 {
	return math.Acos(Dot(a, b))
}

func IsLeftWinding(a, b *T) bool {
	ab := b.Rotated(-a.Angle())
	return ab.Angle() > 0
}

func IsRightWinding(a, b *T) bool {
	ab := b.Rotated(-a.Angle())
	return ab.Angle() < 0
}

func Min(a, b *T) T {
	min := *a
	if b[0] < min[0] {
		min[0] = b[0]
	}
	if b[1] < min[1] {
		min[1] = b[1]
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
	return max
}
