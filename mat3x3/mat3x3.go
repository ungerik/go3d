package mat3x3

import (
	"fmt"

	"github.com/barnex/fmath"
	"github.com/ungerik/go3d/mat2x2"
	"github.com/ungerik/go3d/quaternion"
	"github.com/ungerik/go3d/vec3"
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

func (self *T) Slice() []float32 {
	return []float32{
		self[0][0], self[0][1], self[0][2],
		self[1][0], self[1][1], self[1][2],
		self[2][0], self[2][1], self[2][2],
	}
}

func (self *T) Get(col, row int) float32 {
	return self[col][row]
}

func (self *T) Trace() float32 {
	return self[0][0] + self[1][1] + self[2][2]
}

func (self *T) AssignMul(a, b *T) *T {
	self[0] = a.MulVec3(&b[0])
	self[1] = a.MulVec3(&b[1])
	self[2] = a.MulVec3(&b[2])
	return self
}

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

func (self *T) Scaling() vec3.T {
	return vec3.T{self[0][0], self[1][1], self[2][2]}
}

func (self *T) SetScaling(s *vec3.T) *T {
	self[0][0] = s[0]
	self[1][1] = s[1]
	self[2][2] = s[2]
	return self
}

func (self *T) Scale(s float32) *T {
	self[0][0] *= s
	self[1][1] *= s
	self[2][2] *= s
	return self
}

func (self *T) ScaleVec3(s *vec3.T) *T {
	self[0][0] *= s[0]
	self[1][1] *= s[1]
	self[2][2] *= s[2]
	return self
}

func (self *T) Quaternion() quaternion.T {
	tr := self.Trace()

	s := fmath.Sqrt(tr + 1)
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

func (self *T) AssignXRotation(angle float32) *T {
	cosine := fmath.Cos(angle)
	sine := fmath.Sin(angle)

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

func (self *T) AssignYRotation(angle float32) *T {
	cosine := fmath.Cos(angle)
	sine := fmath.Sin(angle)

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

func (self *T) AssignZRotation(angle float32) *T {
	cosine := fmath.Cos(angle)
	sine := fmath.Sin(angle)

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

func (self *T) AssignEulerRotation(yHead, xPitch, zRoll float32) *T {
	sinH := fmath.Sin(yHead)
	cosH := fmath.Cos(yHead)
	sinP := fmath.Sin(xPitch)
	cosP := fmath.Cos(xPitch)
	sinR := fmath.Sin(zRoll)
	cosR := fmath.Cos(zRoll)

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

func (self *T) ExtractEulerAngles() (yHead, xPitch, zRoll float32) {
	xPitch = fmath.Asin(self[1][2])
	f12 := fmath.Abs(self[1][2])
	if f12 > (1.0-0.0001) && f12 < (1.0+0.0001) { // f12 == 1.0
		yHead = 0.0
		zRoll = fmath.Atan2(self[0][1], self[0][0])
	} else {
		yHead = fmath.Atan2(-self[0][2], self[2][2])
		zRoll = fmath.Atan2(-self[1][0], self[1][1])
	}
	return yHead, xPitch, zRoll
}

func (self *T) Determinant() float32 {
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

func swap(a, b *float32) {
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
