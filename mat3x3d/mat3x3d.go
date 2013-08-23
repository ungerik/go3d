package mat3x3d

import (
	"fmt"
	"math"

	"github.com/ungerik/go3d/mat2x2d"
	"github.com/ungerik/go3d/quaterniond"
	"github.com/ungerik/go3d/vec3d"
)

var (
	Zero  = T{}
	Ident = T{
		vec3d.T{1, 0, 0},
		vec3d.T{0, 1, 0},
		vec3d.T{0, 0, 1},
	}
)

type T [3]vec3d.T

func Parse(s string) (r T, err error) {
	_, err = fmt.Sscanf(s,
		"%f %f %f %f %f %f %f %f %f",
		&r[0][0], &r[0][1], &r[0][2],
		&r[1][0], &r[1][1], &r[1][2],
		&r[2][0], &r[2][1], &r[2][2],
	)
	return r, err
}

func (self *T) String() string {
	return fmt.Sprintf("%s %s %s", self[0].String(), self[1].String(), self[2].String())
}

func (self *T) Rows() int {
	return 3
}

func (self *T) Cols() int {
	return 3
}

func (self *T) Size() int {
	return 9
}

func (self *T) Slice() []float64 {
	return []float64{
		self[0][0], self[0][1], self[0][2],
		self[1][0], self[1][1], self[1][2],
		self[2][0], self[2][1], self[2][2],
	}
}

func (self *T) Get(col, row int) float64 {
	return self[col][row]
}

func (self *T) Trace() float64 {
	return self[0][0] + self[1][1] + self[2][2]
}

func (self *T) AssignMul(a, b *T) *T {
	self[0] = a.MulVec3(&b[0])
	self[1] = a.MulVec3(&b[1])
	self[2] = a.MulVec3(&b[2])
	return self
}

func (self *T) AssignMat2x2(m *mat2x2d.T) *T {
	*self = T{
		vec3d.T{m[0][0], m[1][0], 0},
		vec3d.T{m[0][1], m[1][1], 0},
		vec3d.T{0, 0, 1},
	}
	return self
}

func (self *T) MulVec3(vec *vec3d.T) vec3d.T {
	return vec3d.T{
		self[0][0]*vec[0] + self[1][0]*vec[1] + self[2][0]*vec[2],
		self[0][1]*vec[1] + self[1][1]*vec[1] + self[2][1]*vec[2],
		self[0][2]*vec[2] + self[1][2]*vec[1] + self[2][2]*vec[2],
	}
}

func (self *T) Scaling() vec3d.T {
	return vec3d.T{self[0][0], self[1][1], self[2][2]}
}

func (self *T) SetScaling(s *vec3d.T) *T {
	self[0][0] = s[0]
	self[1][1] = s[1]
	self[2][2] = s[2]
	return self
}

func (self *T) Scale(s float64) *T {
	self[0][0] *= s
	self[1][1] *= s
	self[2][2] *= s
	return self
}

func (self *T) ScaleVec3(s *vec3d.T) *T {
	self[0][0] *= s[0]
	self[1][1] *= s[1]
	self[2][2] *= s[2]
	return self
}

func (self *T) Quaternion() quaterniond.T {
	tr := self.Trace()

	s := math.Sqrt(tr + 1)
	w := s * 0.5
	s = 0.5 / s

	q := quaterniond.T{
		(self[1][2] - self[2][1]) * s,
		(self[2][0] - self[0][2]) * s,
		(self[0][1] - self[1][0]) * s,
		w,
	}
	return q.Normalized()
}

func (self *T) AssignQuaternion(q *quaterniond.T) *T {
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

func (self *T) AssignCoordinateSystem(x, y, z *vec3d.T) *T {
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
