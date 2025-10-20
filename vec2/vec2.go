// Package vec2 contains a 2D float32 vector type T and functions.
package vec2

import (
	"fmt"

	math "github.com/ungerik/go3d/fmath"
	"github.com/ungerik/go3d/generic"
)

var (
	// Zero holds a zero vector.
	Zero = T{}

	// UnitX holds a vector with X set to one.
	UnitX = T{1, 0}
	// UnitY holds a vector with Y set to one.
	UnitY = T{0, 1}
	// UnitXY holds a vector with X and Y set to one.
	UnitXY = T{1, 1}

	// MinVal holds a vector with the smallest possible component values.
	MinVal = T{-math.MaxFloat32, -math.MaxFloat32}
	// MaxVal holds a vector with the highest possible component values.
	MaxVal = T{+math.MaxFloat32, +math.MaxFloat32}
)

// T represents a 2D vector.
type T [2]float32

// From copies a T from a generic.T implementation.
func From(other generic.T) T {
	return T{other.Get(0, 0), other.Get(0, 1)}
}

// Parse parses T from a string. See also String()
func Parse(s string) (r T, err error) {
	_, err = fmt.Sscan(s, &r[0], &r[1])
	return r, err
}

// String formats T as string. See also Parse().
func (vec *T) String() string {
	return fmt.Sprint(vec[0], vec[1])
}

// Rows returns the number of rows of the vector.
func (vec *T) Rows() int {
	return 2
}

// Cols returns the number of columns of the vector.
func (vec *T) Cols() int {
	return 1
}

// Size returns the number elements of the vector.
func (vec *T) Size() int {
	return 2
}

// Slice returns the elements of the vector as slice.
func (vec *T) Slice() []float32 {
	return vec[:]
}

// Get returns one element of the vector.
func (vec *T) Get(col, row int) float32 {
	return vec[row]
}

// IsZero checks if all elements of the vector are exactly zero.
// Uses exact equality comparison, which may not be suitable for floating-point math results.
// For tolerance-based comparison, use IsZeroEps instead.
func (vec *T) IsZero() bool {
	return vec[0] == 0 && vec[1] == 0
}

// IsZeroEps checks if all elements of the vector are zero within the given epsilon tolerance.
// This is the recommended method for comparing floating-point vectors that result from calculations.
// For exact zero comparison, use IsZero instead.
func (vec *T) IsZeroEps(epsilon float32) bool {
	return math.Abs(vec[0]) <= epsilon && math.Abs(vec[1]) <= epsilon
}

// Length returns the length of the vector.
// See also LengthSqr and Normalize.
func (vec *T) Length() float32 {
	return math.Hypot(vec[0], vec[1])
}

// LengthSqr returns the squared length of the vector.
// See also Length and Normalize.
func (vec *T) LengthSqr() float32 {
	return vec[0]*vec[0] + vec[1]*vec[1]
}

// Scale multiplies all element of the vector by f and returns vec.
func (vec *T) Scale(f float32) *T {
	vec[0] *= f
	vec[1] *= f
	return vec
}

// Scaled returns a copy of vec with all elements multiplies by f.
func (vec *T) Scaled(f float32) T {
	return T{vec[0] * f, vec[1] * f}
}

// PracticallyEquals compares two vectors if they are equal with each other within a delta tolerance.
func (vec *T) PracticallyEquals(compareVector *T, allowedDelta float32) bool {
	return (math.Abs(vec[0]-compareVector[0]) <= allowedDelta) &&
		(math.Abs(vec[1]-compareVector[1]) <= allowedDelta)
}

// PracticallyEquals compares two values if they are equal with each other within a delta tolerance.
func PracticallyEquals(v1, v2, allowedDelta float32) bool {
	return math.Abs(v1-v2) <= allowedDelta
}

// Invert inverts the vector.
func (vec *T) Invert() *T {
	vec[0] = -vec[0]
	vec[1] = -vec[1]
	return vec
}

// Inverted returns an inverted copy of the vector.
func (vec *T) Inverted() T {
	return T{-vec[0], -vec[1]}
}

// Abs sets every component of the vector to its absolute value.
func (vec *T) Abs() *T {
	vec[0] = math.Abs(vec[0])
	vec[1] = math.Abs(vec[1])
	return vec
}

// Absed returns a copy of the vector containing the absolute values.
func (vec *T) Absed() T {
	return T{math.Abs(vec[0]), math.Abs(vec[1])}
}

// Normalize normalizes the vector to unit length.
// Uses the package Epsilon variable for numerical stability:
// - Vectors with squared length < Epsilon are considered zero and left unchanged
// - Vectors with squared length within Epsilon of 1.0 are considered already normalized
func (vec *T) Normalize() *T {
	sl := vec.LengthSqr()
	if sl < Epsilon {
		// Vector is effectively zero
		return vec
	}
	if math.Abs(sl-1) < Epsilon {
		// Vector is already normalized
		return vec
	}
	return vec.Scale(1.0 / math.Sqrt(sl))
}

// Normalized returns a unit length normalized copy of the vector.
// Uses the package Epsilon variable for numerical stability. See Normalize() for details.
func (vec *T) Normalized() T {
	v := *vec
	v.Normalize()
	return v
}

// Normal returns a new normalized orthogonal vector.
// The normal is orthogonal clockwise to the vector.
// See also function Rotate90DegRight.
func (vec *T) Normal() T {
	n := *vec
	n[0], n[1] = n[1], -n[0]
	return *n.Normalize()
}

// NormalCCW returns a new normalized orthogonal vector.
// The normal is orthogonal counterclockwise to the vector.
// See also function Rotate90DegLeft.
func (vec *T) NormalCCW() T {
	n := *vec
	n[0], n[1] = -n[1], n[0]
	return *n.Normalize()
}

// Add adds another vector to vec.
func (vec *T) Add(v *T) *T {
	vec[0] += v[0]
	vec[1] += v[1]
	return vec
}

// Added adds another vector to vec and returns a copy of the result
func (vec *T) Added(v *T) T {
	return T{vec[0] + v[0], vec[1] + v[1]}
}

// Sub subtracts another vector from vec.
func (vec *T) Sub(v *T) *T {
	vec[0] -= v[0]
	vec[1] -= v[1]
	return vec
}

// Subed subtracts another vector from vec and returns a copy of the result
func (vec *T) Subed(v *T) T {
	return T{vec[0] - v[0], vec[1] - v[1]}
}

// Mul multiplies the components of the vector with the respective components of v.
func (vec *T) Mul(v *T) *T {
	vec[0] *= v[0]
	vec[1] *= v[1]
	return vec
}

// Muled multiplies the components of the vector with the respective components of v and returns a copy of the result
func (vec *T) Muled(v *T) T {
	return T{vec[0] * v[0], vec[1] * v[1]}
}

// Rotate rotates the vector counter-clockwise by angle.
func (vec *T) Rotate(angle float32) *T {
	*vec = vec.Rotated(angle)
	return vec
}

// Rotated returns a counter-clockwise rotated copy of the vector.
func (vec *T) Rotated(angle float32) T {
	sinus := math.Sin(angle)
	cosinus := math.Cos(angle)
	return T{
		vec[0]*cosinus - vec[1]*sinus,
		vec[0]*sinus + vec[1]*cosinus,
	}
}

// RotateAroundPoint rotates the vector counter-clockwise around a point.
func (vec *T) RotateAroundPoint(point *T, angle float32) *T {
	return vec.Sub(point).Rotate(angle).Add(point)
}

// Rotate90DegLeft rotates the vector 90 degrees left (counter-clockwise).
// See also function NormalCCW.
func (vec *T) Rotate90DegLeft() *T {
	vec[0], vec[1] = -vec[1], vec[0]
	return vec
}

// Rotate90DegRight rotates the vector 90 degrees right (clockwise).
// See also function Normal.
func (vec *T) Rotate90DegRight() *T {
	vec[0], vec[1] = vec[1], -vec[0]
	return vec
}

// Sinus returns the sinus value of the (shortest/smallest) angle between the two vectors a and b.
// The returned sine value is in the range -1.0 ≤ value ≤ 1.0.
// The angle is always considered to be in the range 0 to Pi radians and thus the sine value returned is always positive.
func Sinus(a, b *T) float32 {
	v := Cross(a, b) / math.Sqrt(a.LengthSqr()*b.LengthSqr())

	if v > 1.0 {
		return 1.0
	} else if v < -1.0 {
		return -1.0
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

// Angle returns the counter-clockwise angle of the vector from the x axis.
func (vec *T) Angle() float32 {
	return math.Atan2(vec[1], vec[0])
}

// Add adds the composants of the two vectors and returns a new vector with the sum of the two vectors.
func Add(a, b *T) T {
	return T{a[0] + b[0], a[1] + b[1]}
}

// Sub returns the difference of two vectors.
func Sub(a, b *T) T {
	return T{a[0] - b[0], a[1] - b[1]}
}

// Mul returns the component wise product of two vectors.
func Mul(a, b *T) T {
	return T{a[0] * b[0], a[1] * b[1]}
}

// Dot returns the dot product of two vectors.
func Dot(a, b *T) float32 {
	return a[0]*b[0] + a[1]*b[1]
}

// Cross returns the "cross product" of two vectors.
// In 2D space it is a scalar value.
// It is the same as the determinant value of the 2D matrix constructed by the two vectors.
// Cross product in 2D is not well-defined but this is the implementation stated at https://mathworld.wolfram.com/CrossProduct.html .
func Cross(a, b *T) float32 {
	return a[0]*b[1] - a[1]*b[0]
}

// Angle returns the angle value of the (shortest/smallest) angle between the two vectors a and b.
// The returned value is in the range 0 ≤ angle ≤ Pi radians.
func Angle(a, b *T) float32 {
	return math.Acos(Cosine(a, b))
}

// IsLeftWinding returns if the angle from a to b is left winding.
// Two parallell or anti parallell vectors will give a false result.
func IsLeftWinding(a, b *T) bool {
	return Cross(a, b) > 0 // It's really the sign changing part of the Sinus(a, b) function
}

// IsRightWinding returns if the angle from a to b is right winding.
// Two parallell or anti parallell vectors will give a false result.
func IsRightWinding(a, b *T) bool {
	return Cross(a, b) < 0 // It's really the sign changing part of the Sinus(a, b) function
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
	return max
}

// Interpolate interpolates between a and b at t (0,1).
func Interpolate(a, b *T, t float32) T {
	t1 := 1.0 - t
	return T{
		a[0]*t1 + b[0]*t,
		a[1]*t1 + b[1]*t,
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
	return vec.Clamp(&Zero, &UnitXY)
}

// Clamped01 returns a copy of the vector with the components clamped to be in the range of 0 to 1.
func (vec *T) Clamped01() T {
	result := *vec
	result.Clamp01()
	return result
}
