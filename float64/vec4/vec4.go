package vec4

import (
	"fmt"
	"math"

	"github.com/ungerik/go3d/float64/generic"
	"github.com/ungerik/go3d/float64/vec3"
)

var (
	Zero = T{}

	UnitXW = T{1, 0, 0, 1}
	UnitYW = T{0, 1, 0, 1}
	UnitZW = T{0, 0, 1, 1}
	UnitW  = T{0, 0, 0, 1}

	Red   = T{1, 0, 0, 1}
	Green = T{0, 1, 0, 1}
	Blue  = T{0, 0, 1, 1}
	Black = T{0, 0, 0, 1}
	White = T{1, 1, 1, 1}

	MinVal = T{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64, 1}
	MaxVal = T{+math.MaxFloat64, +math.MaxFloat64, +math.MaxFloat64, 1}
)

type T [4]float64

// From copies a T from a generic.T implementation.
func From(other generic.T) T {
	switch other.Size() {
	case 2:
		return T{other.Get(0, 0), other.Get(0, 1), 0, 1}
	case 3:
		return T{other.Get(0, 0), other.Get(0, 1), other.Get(0, 2), 1}
	case 4:
		return T{other.Get(0, 0), other.Get(0, 1), other.Get(0, 2), other.Get(0, 3)}
	default:
		panic("Unsupported type")
	}
}

func FromVec3(other *vec3.T) T {
	return T{other[0], other[1], other[2], 1}
}

// Parse parses T from a string. See also String()
func Parse(s string) (r T, err error) {
	_, err = fmt.Sscanf(s, "%f %f %f %f", &r[0], &r[1], &r[2], &r[3])
	return r, err
}

// String formats T as string. See also Parse().
func (self *T) String() string {
	return fmt.Sprintf("%f %f %f %f", self[0], self[1], self[2], self[3])
}

// Rows returns the number of rows of the vector.
func (self *T) Rows() int {
	return 4
}

// Cols returns the number of columns of the vector.
func (self *T) Cols() int {
	return 1
}

// Size returns the number elements of the vector.
func (self *T) Size() int {
	return 4
}

// Slice returns the elements of the vector as slice.
func (self *T) Slice() []float64 {
	return []float64{self[0], self[1], self[2], self[3]}
}

// Get returns one element of the vector.
func (self *T) Get(col, row int) float64 {
	return self[row]
}

// IsZero checks if all elements of the vector are zero.
func (self *T) IsZero() bool {
	return self[0] == 0 && self[1] == 0 && self[2] == 0 && self[3] == 0
}

// Length returns the length of the vector.
// See also LengthSqr and Normalize.
func (self *T) Length() float64 {
	v3 := self.Vec3DividedByW()
	return v3.Length()
}

// Length returns the squared length of the vector.
// See also Length and Normalize.
func (self *T) LengthSqr() float64 {
	v3 := self.Vec3DividedByW()
	return v3.LengthSqr()
}

// Scale multiplies the first 3 element of the vector by f and returns self.
func (self *T) Scale(f float64) *T {
	self[0] *= f
	self[1] *= f
	self[2] *= f
	return self
}

// Scaled returns a copy of self with the first 3 elements multiplies by f.
func (self *T) Scaled(f float64) T {
	return T{self[0] * f, self[1] * f, self[2] * f, self[3]}
}

// Invert inverts the vector.
func (self *T) Invert() {
	self[0] = -self[0]
	self[1] = -self[1]
	self[2] = -self[2]
}

// Inverted returns an inverted copy of the vector.
func (self *T) Inverted() T {
	return T{-self[0], -self[1], -self[2], self[3]}
}

// Normalize normalizes the vector to unit length.
func (self *T) Normalize() *T {
	v3 := self.Vec3DividedByW()
	v3.Normalize()
	self[0] = v3[0]
	self[1] = v3[1]
	self[2] = v3[2]
	self[3] = 1
	return self
}

// Normalized returns a unit length normalized copy of the vector.
func (self *T) Normalized() T {
	v := *self
	v.Normalize()
	return v
}

// Normal returns an orthogonal vector.
func (self *T) Normal() T {
	v3 := self.Vec3()
	n3 := v3.Normal()
	return T{n3[0], n3[1], n3[2], 1}
}

// DivideByW divides the first three components (XYZ) by the last one (W).
func (self *T) DivideByW() *T {
	oow := 1 / self[3]
	self[0] *= oow
	self[1] *= oow
	self[2] *= oow
	self[3] = 1
	return self
}

// DividedByW returns a copy of the vector with the first three components (XYZ) divided by the last one (W).
func (self *T) DividedByW() T {
	oow := 1 / self[3]
	return T{self[0] * oow, self[1] * oow, self[2] * oow, 1}
}

// Vec3DividedByW returns a vec3.T version of the vector by dividing the first three vector components (XYZ) by the last one (W).
func (self *T) Vec3DividedByW() vec3.T {
	oow := 1 / self[3]
	return vec3.T{self[0] * oow, self[1] * oow, self[2] * oow}
}

// Vec3 returns a vec3.T with the first three components of the vector.
// See also Vec3DividedByW
func (self *T) Vec3() vec3.T {
	return vec3.T{self[0], self[1], self[2]}
}

// AssignVec3 assigns v to the first three components and sets the fourth to 1.
func (self *T) AssignVec3(v *vec3.T) *T {
	self[0] = v[0]
	self[1] = v[1]
	self[2] = v[2]
	self[3] = 1
	return self
}

// Add adds another vector to self.
func (self *T) Add(v *T) *T {
	if v[3] == self[3] {
		self[0] += v[0]
		self[1] += v[1]
		self[2] += v[2]
	} else {
		self.DividedByW()
		v3 := v.Vec3DividedByW()
		self[0] += v3[0]
		self[1] += v3[1]
		self[2] += v3[2]
	}
	return self
}

// Sub subtracts another vector from self.
func (self *T) Sub(v *T) *T {
	if v[3] == self[3] {
		self[0] -= v[0]
		self[1] -= v[1]
		self[2] -= v[2]
	} else {
		self.DividedByW()
		v3 := v.Vec3DividedByW()
		self[0] -= v3[0]
		self[1] -= v3[1]
		self[2] -= v3[2]
	}
	return self
}

// Add returns the sum of two vectors.
func Add(a, b *T) T {
	if a[3] == b[3] {
		return T{a[0] + b[0], a[1] + b[1], a[2] + b[2], 1}
	} else {
		a3 := a.Vec3DividedByW()
		b3 := b.Vec3DividedByW()
		return T{a3[0] + b3[0], a3[1] + b3[1], a3[2] + b3[2], 1}
	}
}

// Add returns the difference of two vectors.
func Sub(a, b *T) T {
	if a[3] == b[3] {
		return T{a[0] - b[0], a[1] - b[1], a[2] - b[2], 1}
	} else {
		a3 := a.Vec3DividedByW()
		b3 := b.Vec3DividedByW()
		return T{a3[0] - b3[0], a3[1] - b3[1], a3[2] - b3[2], 1}
	}
}

// Dot returns the dot product of two (dived by w) vectors.
func Dot(a, b *T) float64 {
	a3 := a.Vec3DividedByW()
	b3 := b.Vec3DividedByW()
	return vec3.Dot(&a3, &b3)
}

// Dot returns the 4 element vdot product of two vectors.
func Dot4(a, b *T) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

// Cross returns the cross product of two vectors.
func Cross(a, b *T) T {
	a3 := a.Vec3DividedByW()
	b3 := b.Vec3DividedByW()
	c3 := vec3.Cross(&a3, &b3)
	return T{c3[0], c3[1], c3[2], 1}
}

// Angle returns the angle between two vectors.
func Angle(a, b *T) float64 {
	return math.Acos(Dot(a, b))
}
