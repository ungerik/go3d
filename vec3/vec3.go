// Package vec3 contains a 3D float32 vector type T and functions.
package vec3

import (
	"fmt"

	math "github.com/ungerik/go3d/fmath"
	"github.com/ungerik/go3d/generic"
)

var (
	// Zero holds a zero vector.
	Zero = T{}

	// UnitX holds a vector with X set to one.
	UnitX = T{1, 0, 0}
	// UnitY holds a vector with Y set to one.
	UnitY = T{0, 1, 0}
	// UnitZ holds a vector with Z set to one.
	UnitZ = T{0, 0, 1}
	// UnitXYZ holds a vector with X, Y, Z set to one.
	UnitXYZ = T{1, 1, 1}

	// Red holds the color red.
	Red = T{1, 0, 0}
	// Green holds the color green.
	Green = T{0, 1, 0}
	// Blue holds the color black.
	Blue = T{0, 0, 1}
	// Black holds the color black.
	Black = T{0, 0, 0}
	// White holds the color white.
	White = T{1, 1, 1}

	// MinVal holds a vector with the smallest possible component values.
	MinVal = T{-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32}
	// MaxVal holds a vector with the highest possible component values.
	MaxVal = T{+math.MaxFloat32, +math.MaxFloat32, +math.MaxFloat32}
)

// T represents a 3D vector.
type T [3]float32

// From copies a T from a generic.T implementation.
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

// Parse parses T from a string. See also String()
func Parse(s string) (r T, err error) {
	_, err = fmt.Sscan(s, &r[0], &r[1], &r[2])
	return r, err
}

// String formats T as string. See also Parse().
func (vec *T) String() string {
	return fmt.Sprint(vec[0], vec[1], vec[2])
}

// Rows returns the number of rows of the vector.
func (vec *T) Rows() int {
	return 3
}

// Cols returns the number of columns of the vector.
func (vec *T) Cols() int {
	return 1
}

// Size returns the number elements of the vector.
func (vec *T) Size() int {
	return 3
}

// Slice returns the elements of the vector as slice.
func (vec *T) Slice() []float32 {
	return vec[:]
}

// Get returns one element of the vector.
func (vec *T) Get(col, row int) float32 {
	return vec[row]
}

// IsZero checks if all elements of the vector are zero.
func (vec *T) IsZero() bool {
	return vec[0] == 0 && vec[1] == 0 && vec[2] == 0
}

// Length returns the length of the vector.
// See also LengthSqr and Normalize.
func (vec *T) Length() float32 {
	return float32(math.Sqrt(vec.LengthSqr()))
}

// LengthSqr returns the squared length of the vector.
// See also Length and Normalize.
func (vec *T) LengthSqr() float32 {
	return vec[0]*vec[0] + vec[1]*vec[1] + vec[2]*vec[2]
}

// Scale multiplies all element of the vector by f and returns vec.
func (vec *T) Scale(f float32) *T {
	vec[0] *= f
	vec[1] *= f
	vec[2] *= f
	return vec
}

// Scaled returns a copy of vec with all elements multiplies by f.
func (vec *T) Scaled(f float32) T {
	return T{vec[0] * f, vec[1] * f, vec[2] * f}
}

// PracticallyEquals compares two vectors if they are equal with each other within a delta tolerance.
func (vec *T) PracticallyEquals(compareVector *T, allowedDelta float32) bool {
	return (math.Abs(vec[0]-compareVector[0]) <= allowedDelta) &&
		(math.Abs(vec[1]-compareVector[1]) <= allowedDelta) &&
		(math.Abs(vec[2]-compareVector[2]) <= allowedDelta)
}

// PracticallyEquals compares two values if they are equal with each other within a delta tolerance.
func PracticallyEquals(v1, v2, allowedDelta float32) bool {
	return math.Abs(v1-v2) <= allowedDelta
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
	return T{-vec[0], -vec[1], -vec[2]}
}

// Abs sets every component of the vector to its absolute value.
func (vec *T) Abs() *T {
	vec[0] = math.Abs(vec[0])
	vec[1] = math.Abs(vec[1])
	vec[2] = math.Abs(vec[2])
	return vec
}

// Absed returns a copy of the vector containing the absolute values.
func (vec *T) Absed() T {
	return T{math.Abs(vec[0]), math.Abs(vec[1]), math.Abs(vec[2])}
}

// Normalize normalizes the vector to unit length.
func (vec *T) Normalize() *T {
	sl := vec.LengthSqr()
	if sl == 0 || sl == 1 {
		return vec
	}
	return vec.Scale(1 / math.Sqrt(sl))
}

// Normalized returns a unit length normalized copy of the vector.
func (vec *T) Normalized() T {
	v := *vec
	v.Normalize()
	return v
}

// Normal returns an orthogonal vector.
func (vec *T) Normal() T {
	n := Cross(vec, &UnitZ)
	if n.IsZero() {
		return UnitX
	}
	return n.Normalized()
}

// Add adds another vector to vec.
func (vec *T) Add(v *T) *T {
	vec[0] += v[0]
	vec[1] += v[1]
	vec[2] += v[2]
	return vec
}

// Added adds another vector to vec and returns a copy of the result
func (vec *T) Added(v *T) T {
	return T{vec[0] + v[0], vec[1] + v[1], vec[2] + v[2]}
}

// Sub subtracts another vector from vec.
func (vec *T) Sub(v *T) *T {
	vec[0] -= v[0]
	vec[1] -= v[1]
	vec[2] -= v[2]
	return vec
}

// Subed subtracts another vector from vec and returns a copy of the result
func (vec *T) Subed(v *T) T {
	return T{vec[0] - v[0], vec[1] - v[1], vec[2] - v[2]}
}

// Mul multiplies the components of the vector with the respective components of v.
func (vec *T) Mul(v *T) *T {
	vec[0] *= v[0]
	vec[1] *= v[1]
	vec[2] *= v[2]
	return vec
}

// Muled multiplies the components of the vector with the respective components of v and returns a copy of the result
func (vec *T) Muled(v *T) T {
	return T{vec[0] * v[0], vec[1] * v[1], vec[2] * v[2]}
}

// SquareDistance the distance between two vectors squared (= distance*distance)
func SquareDistance(a, b *T) float32 {
	d := Sub(a, b)
	return d.LengthSqr()
}

// Distance the distance between two vectors
func Distance(a, b *T) float32 {
	d := Sub(a, b)
	return d.Length()
}

// Add adds the composants of the two vectors and returns a new vector with the sum of the two vectors.
func Add(a, b *T) T {
	return T{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

// Sub returns the difference of two vectors.
func Sub(a, b *T) T {
	return T{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

// Mul returns the component wise product of two vectors.
func Mul(a, b *T) T {
	return T{a[0] * b[0], a[1] * b[1], a[2] * b[2]}
}

// Dot returns the dot product of two vectors.
func Dot(a, b *T) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// Cross returns the cross product of two vectors.
func Cross(a, b *T) T {
	return T{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

// Sinus returns the sinus value of the (shortest/smallest) angle between the two vectors a and b.
// The returned sine value is in the range 0.0 ≤ value ≤ 1.0.
// The angle is always considered to be in the range 0 to Pi radians and thus the sine value returned is always positive.
func Sinus(a, b *T) float32 {
	cross := Cross(a, b)
	v := cross.Length() / math.Sqrt(a.LengthSqr()*b.LengthSqr())

	if v > 1.0 {
		return 1.0
	} else if v < 0.0 {
		return 0.0
	}
	return v
}

// Cosine returns the cosine value of the angle between the two vectors.
// The returned cosine value is in the range -1.0 ≤ value ≤ 1.0.
func Cosine(a, b *T) float32 {
	v := Dot(a, b) / math.Sqrt(a.LengthSqr()*b.LengthSqr())

	if v > 1.0 {
		return 1.0
	} else if v < -1.0 {
		return -1.0
	}
	return v
}

// Angle returns the angle value of the (shortest/smallest) angle between the two vectors a and b.
// The returned value is in the range 0 ≤ angle ≤ Pi radians.
func Angle(a, b *T) float32 {
	return math.Acos(Cosine(a, b))
}

// Min returns the component wise minimum of two vectors.
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

// Max returns the component wise maximum of two vectors.
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

// Interpolate interpolates between a and b at t (0,1).
func Interpolate(a, b *T, t float32) T {
	t1 := 1 - t
	return T{
		a[0]*t1 + b[0]*t,
		a[1]*t1 + b[1]*t,
		a[2]*t1 + b[2]*t,
	}
}

// Clamp clamps the vector's components to be in the range of min to max.
func (vec *T) Clamp(min, max *T) *T {
	for i := range vec {
		if vec[i] < min[i] {
			vec[i] = min[i]
		} else if vec[i] > max[i] {
			vec[i] = max[i]
		}
	}
	return vec
}

// Clamped returns a copy of the vector with the components clamped to be in the range of min to max.
func (vec *T) Clamped(min, max *T) T {
	result := *vec
	result.Clamp(min, max)
	return result
}

// Clamp01 clamps the vector's components to be in the range of 0 to 1.
func (vec *T) Clamp01() *T {
	return vec.Clamp(&Zero, &UnitXYZ)
}

// Clamped01 returns a copy of the vector with the components clamped to be in the range of 0 to 1.
func (vec *T) Clamped01() T {
	result := *vec
	result.Clamp01()
	return result
}
