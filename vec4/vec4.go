// Package vec4 contains a 4 float32 components vector type T and functions.
package vec4

import (
	"fmt"
	"math"

	"github.com/barnex/fmath"
	"github.com/ungerik/go3d/generic"
	"github.com/ungerik/go3d/vec3"
)

var (
	// Zero holds a zero vector.
	Zero = T{}

	// UnitXW holds a vector with X and W set to one.
	UnitXW = T{1, 0, 0, 1}
	// UnitYW holds a vector with Y and W set to one.
	UnitYW = T{0, 1, 0, 1}
	// UnitZW holds a vector with Z and W set to one.
	UnitZW = T{0, 0, 1, 1}
	// UnitW holds a vector with W set to one.
	UnitW = T{0, 0, 0, 1}

	// Red holds the color red.
	Red = T{1, 0, 0, 1}
	// Green holds the color green.
	Green = T{0, 1, 0, 1}
	// Blue holds the color blue.
	Blue = T{0, 0, 1, 1}
	// Black holds the color black.
	Black = T{0, 0, 0, 1}
	// White holds the color white.
	White = T{1, 1, 1, 1}

	// MinVal holds a vector with the smallest possible component values.
	MinVal = T{-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32, 1}
	// MaxVal holds a vector with the highest possible component values.
	MaxVal = T{+math.MaxFloat32, +math.MaxFloat32, +math.MaxFloat32, 1}
)

// T represents a 4 component vector.
type T [4]float32

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

// FromVec3 returns a vector with the first 3 components copied from a vec3.T.
func FromVec3(other *vec3.T) T {
	return T{other[0], other[1], other[2], 1}
}

// Parse parses T from a string. See also String()
func Parse(s string) (r T, err error) {
	_, err = fmt.Sscanf(s, "%f %f %f %f", &r[0], &r[1], &r[2], &r[3])
	return r, err
}

// String formats T as string. See also Parse().
func (vec *T) String() string {
	return fmt.Sprintf("%f %f %f %f", vec[0], vec[1], vec[2], vec[3])
}

// Rows returns the number of rows of the vector.
func (vec *T) Rows() int {
	return 4
}

// Cols returns the number of columns of the vector.
func (vec *T) Cols() int {
	return 1
}

// Size returns the number elements of the vector.
func (vec *T) Size() int {
	return 4
}

// Slice returns the elements of the vector as slice.
func (vec *T) Slice() []float32 {
	return []float32{vec[0], vec[1], vec[2], vec[3]}
}

// Get returns one element of the vector.
func (vec *T) Get(col, row int) float32 {
	return vec[row]
}

// IsZero checks if all elements of the vector are zero.
func (vec *T) IsZero() bool {
	return vec[0] == 0 && vec[1] == 0 && vec[2] == 0 && vec[3] == 0
}

// Length returns the length of the vector.
// See also LengthSqr and Normalize.
func (vec *T) Length() float32 {
	v3 := vec.Vec3DividedByW()
	return v3.Length()
}

// Length returns the squared length of the vector.
// See also Length and Normalize.
func (vec *T) LengthSqr() float32 {
	v3 := vec.Vec3DividedByW()
	return v3.LengthSqr()
}

// Scale multiplies the first 3 element of the vector by f and returns vec.
func (vec *T) Scale(f float32) *T {
	vec[0] *= f
	vec[1] *= f
	vec[2] *= f
	return vec
}

// Scaled returns a copy of vec with the first 3 elements multiplies by f.
func (vec *T) Scaled(f float32) T {
	return T{vec[0] * f, vec[1] * f, vec[2] * f, vec[3]}
}

// Invert inverts the vector.
func (vec *T) Invert() *T {
	vec[0] = -vec[0]
	vec[1] = -vec[1]
	vec[2] = -vec[2]
	return vec
}

// Inverted returns an inverted copy of the vector.
func (vec *T) Inverted() T {
	return T{-vec[0], -vec[1], -vec[2], vec[3]}
}

// Normalize normalizes the vector to unit length.
func (vec *T) Normalize() *T {
	v3 := vec.Vec3DividedByW()
	v3.Normalize()
	vec[0] = v3[0]
	vec[1] = v3[1]
	vec[2] = v3[2]
	vec[3] = 1
	return vec
}

// Normalized returns a unit length normalized copy of the vector.
func (vec *T) Normalized() T {
	v := *vec
	v.Normalize()
	return v
}

// Normal returns an orthogonal vector.
func (vec *T) Normal() T {
	v3 := vec.Vec3()
	n3 := v3.Normal()
	return T{n3[0], n3[1], n3[2], 1}
}

// DivideByW divides the first three components (XYZ) by the last one (W).
func (vec *T) DivideByW() *T {
	oow := 1 / vec[3]
	vec[0] *= oow
	vec[1] *= oow
	vec[2] *= oow
	vec[3] = 1
	return vec
}

// DividedByW returns a copy of the vector with the first three components (XYZ) divided by the last one (W).
func (vec *T) DividedByW() T {
	oow := 1 / vec[3]
	return T{vec[0] * oow, vec[1] * oow, vec[2] * oow, 1}
}

// Vec3DividedByW returns a vec3.T version of the vector by dividing the first three vector components (XYZ) by the last one (W).
func (vec *T) Vec3DividedByW() vec3.T {
	oow := 1 / vec[3]
	return vec3.T{vec[0] * oow, vec[1] * oow, vec[2] * oow}
}

// Vec3 returns a vec3.T with the first three components of the vector.
// See also Vec3DividedByW
func (vec *T) Vec3() vec3.T {
	return vec3.T{vec[0], vec[1], vec[2]}
}

// AssignVec3 assigns v to the first three components and sets the fourth to 1.
func (vec *T) AssignVec3(v *vec3.T) *T {
	vec[0] = v[0]
	vec[1] = v[1]
	vec[2] = v[2]
	vec[3] = 1
	return vec
}

// Add adds another vector to vec.
func (vec *T) Add(v *T) *T {
	if v[3] == vec[3] {
		vec[0] += v[0]
		vec[1] += v[1]
		vec[2] += v[2]
	} else {
		vec.DividedByW()
		v3 := v.Vec3DividedByW()
		vec[0] += v3[0]
		vec[1] += v3[1]
		vec[2] += v3[2]
	}
	return vec
}

// Sub subtracts another vector from vec.
func (vec *T) Sub(v *T) *T {
	if v[3] == vec[3] {
		vec[0] -= v[0]
		vec[1] -= v[1]
		vec[2] -= v[2]
	} else {
		vec.DividedByW()
		v3 := v.Vec3DividedByW()
		vec[0] -= v3[0]
		vec[1] -= v3[1]
		vec[2] -= v3[2]
	}
	return vec
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
func Dot(a, b *T) float32 {
	a3 := a.Vec3DividedByW()
	b3 := b.Vec3DividedByW()
	return vec3.Dot(&a3, &b3)
}

// Dot returns the 4 element dot product of two vectors.
func Dot4(a, b *T) float32 {
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
func Angle(a, b *T) float32 {
	return fmath.Acos(Dot(a, b))
}
