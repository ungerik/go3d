package mat3x3

import (
	"fmt"
	"math"

	"github.com/ungerik/go3d/float64/generic"
	"github.com/ungerik/go3d/float64/mat2x2"
	"github.com/ungerik/go3d/float64/quaternion"
	"github.com/ungerik/go3d/float64/vec2"
	"github.com/ungerik/go3d/float64/vec3"
)

var (
	Zero  = T{}
	Ident = T{
		vec3.T{1, 0, 0},
		vec3.T{0, 1, 0},
		vec3.T{0, 0, 1},
	}
)

type T [3]vec3.T

// From copies a T from a generic.T implementation.
func From(other generic.T) T {
	r := Ident
	cols := other.Cols()
	rows := other.Rows()
	if cols == 4 && rows == 4 {
		cols = 3
		rows = 3
	} else if !((cols == 2 && rows == 2) || (cols == 3 && rows == 3)) {
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
		"%f %f %f %f %f %f %f %f %f",
		&r[0][0], &r[0][1], &r[0][2],
		&r[1][0], &r[1][1], &r[1][2],
		&r[2][0], &r[2][1], &r[2][2],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (self *T) String() string {
	return fmt.Sprintf("%s %s %s", self[0].String(), self[1].String(), self[2].String())
}

// Rows returns the number of rows of the matrix.
func (self *T) Rows() int {
	return 3
}

// Cols returns the number of columns of the matrix.
func (self *T) Cols() int {
	return 3
}

// Size returns the number elements of the matrix.
func (self *T) Size() int {
	return 9
}

// Slice returns the elements of the matrix as slice.
func (self *T) Slice() []float64 {
	return []float64{
		self[0][0], self[0][1], self[0][2],
		self[1][0], self[1][1], self[1][2],
		self[2][0], self[2][1], self[2][2],
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
	self[2][2] *= f
	return self
}

// Scaled returns a copy of the matrix with the diagonal scale elements multiplied by f.
func (self *T) Scaled(f float64) T {
	r := *self
	return *r.Scale(f)
}

// Scaling returns the scaling diagonal of the matrix.
func (self *T) Scaling() vec3.T {
	return vec3.T{self[0][0], self[1][1], self[2][2]}
}

// SetScaling sets the scaling diagonal of the matrix.
func (self *T) SetScaling(s *vec3.T) *T {
	self[0][0] = s[0]
	self[1][1] = s[1]
	self[2][2] = s[2]
	return self
}

// ScaleVec3 multiplies the 2D scaling diagonal of the matrix by s.
func (self *T) ScaleVec2(s *vec2.T) *T {
	self[0][0] *= s[0]
	self[1][1] *= s[1]
	return self
}

// SetTranslation sets the 2D translation elements of the matrix.
func (self *T) SetTranslation(v *vec2.T) *T {
	self[2][0] = v[0]
	self[2][1] = v[1]
	return self
}

// Translate adds v to the 2D translation part of the matrix.
func (self *T) Translate(v *vec2.T) *T {
	self[2][0] += v[0]
	self[2][1] += v[1]
	return self
}

// Translate adds dx to the 2D X-translation element of the matrix.
func (self *T) TranslateX(dx float64) *T {
	self[2][0] += dx
	return self
}

// Translate adds dy to the 2D Y-translation element of the matrix.
func (self *T) TranslateY(dy float64) *T {
	self[2][1] += dy
	return self
}

func (self *T) Trace() float64 {
	return self[0][0] + self[1][1] + self[2][2]
}

// AssignMul multiplies a and b and assigns the result to self.
func (self *T) AssignMul(a, b *T) *T {
	self[0] = a.MulVec3(&b[0])
	self[1] = a.MulVec3(&b[1])
	self[2] = a.MulVec3(&b[2])
	return self
}

// AssignMat2x2 assigns a 2x2 sub-matrix and sets the rest of the matrix to the ident value.
func (self *T) AssignMat2x2(m *mat2x2.T) *T {
	*self = T{
		vec3.T{m[0][0], m[1][0], 0},
		vec3.T{m[0][1], m[1][1], 0},
		vec3.T{0, 0, 1},
	}
	return self
}

func (self *T) MulVec3(vec *vec3.T) vec3.T {
	return vec3.T{
		self[0][0]*vec[0] + self[1][0]*vec[1] + self[2][0]*vec[2],
		self[0][1]*vec[1] + self[1][1]*vec[1] + self[2][1]*vec[2],
		self[0][2]*vec[2] + self[1][2]*vec[1] + self[2][2]*vec[2],
	}
}

func (self *T) Quaternion() quaternion.T {
	tr := self.Trace()

	s := math.Sqrt(tr + 1)
	w := s * 0.5
	s = 0.5 / s

	q := quaternion.T{
		(self[1][2] - self[2][1]) * s,
		(self[2][0] - self[0][2]) * s,
		(self[0][1] - self[1][0]) * s,
		w,
	}
	return q.Normalized()
}

func (self *T) AssignQuaternion(q *quaternion.T) *T {
	xx := q[0] * q[0] * 2
	yy := q[1] * q[1] * 2
	zz := q[2] * q[2] * 2
	xy := q[0] * q[1] * 2
	xz := q[0] * q[2] * 2
	yz := q[1] * q[2] * 2
	wx := q[3] * q[0] * 2
	wy := q[3] * q[1] * 2
	wz := q[3] * q[2] * 2

	self[0][0] = 1 - (yy + zz)
	self[1][0] = xy - wz
	self[2][0] = xz + wy

	self[0][1] = xy + wz
	self[1][1] = 1 - (xx + zz)
	self[2][1] = yz - wx

	self[0][2] = xz - wy
	self[1][2] = yz + wx
	self[2][2] = 1 - (xx + yy)

	return self
}

func (self *T) AssignXRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	self[0][0] = 1
	self[1][0] = 0
	self[2][0] = 0

	self[0][1] = 0
	self[1][1] = cosine
	self[2][1] = -sine

	self[0][2] = 0
	self[1][2] = sine
	self[2][2] = cosine

	return self
}

func (self *T) AssignYRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	self[0][0] = cosine
	self[1][0] = 0
	self[2][0] = sine

	self[0][1] = 0
	self[1][1] = 1
	self[2][1] = 0

	self[0][2] = -sine
	self[1][2] = 0
	self[2][2] = cosine

	return self
}

func (self *T) AssignZRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	self[0][0] = cosine
	self[1][0] = -sine
	self[2][0] = 0

	self[0][1] = sine
	self[1][1] = cosine
	self[2][1] = 0

	self[0][2] = 0
	self[1][2] = 0
	self[2][2] = 1

	return self
}

func (self *T) AssignCoordinateSystem(x, y, z *vec3.T) *T {
	self[0][0] = x[0]
	self[1][0] = x[1]
	self[2][0] = x[2]

	self[0][1] = y[0]
	self[1][1] = y[1]
	self[2][1] = y[2]

	self[0][2] = z[0]
	self[1][2] = z[1]
	self[2][2] = z[2]

	return self
}

func (self *T) AssignEulerRotation(yHead, xPitch, zRoll float64) *T {
	sinH := math.Sin(yHead)
	cosH := math.Cos(yHead)
	sinP := math.Sin(xPitch)
	cosP := math.Cos(xPitch)
	sinR := math.Sin(zRoll)
	cosR := math.Cos(zRoll)

	self[0][0] = cosR*cosH - sinR*sinP*sinH
	self[1][0] = -sinR * cosP
	self[2][0] = cosR*sinH + sinR*sinP*cosH

	self[0][1] = sinR*cosH + cosR*sinP*sinH
	self[1][1] = cosR * cosP
	self[2][1] = sinR*sinH - cosR*sinP*cosH

	self[0][2] = -cosP * sinH
	self[1][2] = sinP
	self[2][2] = cosP * cosH

	return self
}

func (self *T) ExtractEulerAngles() (yHead, xPitch, zRoll float64) {
	xPitch = math.Asin(self[1][2])
	f12 := math.Abs(self[1][2])
	if f12 > (1.0-0.0001) && f12 < (1.0+0.0001) { // f12 == 1.0
		yHead = 0.0
		zRoll = math.Atan2(self[0][1], self[0][0])
	} else {
		yHead = math.Atan2(-self[0][2], self[2][2])
		zRoll = math.Atan2(-self[1][0], self[1][1])
	}
	return yHead, xPitch, zRoll
}

func (self *T) Determinant() float64 {
	return self[0][0]*self[1][1]*self[2][2] +
		self[1][0]*self[2][1]*self[0][2] +
		self[2][0]*self[0][1]*self[1][2] -
		self[2][0]*self[1][1]*self[0][2] -
		self[1][0]*self[0][1]*self[2][2] -
		self[0][0]*self[2][1]*self[1][2]
}

func (self *T) IsReflective() bool {
	return self.Determinant() < 0
}

func swap(a, b *float64) {
	temp := *a
	*a = *b
	*b = temp
}

func (self *T) Transpose() *T {
	swap(&self[1][0], &self[0][1])
	swap(&self[2][0], &self[0][2])
	swap(&self[2][1], &self[1][2])
	return self
}
