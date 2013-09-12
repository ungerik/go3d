package mat4x4d

import (
	"fmt"
	"math"

	"github.com/ungerik/go3d/genericd"
	"github.com/ungerik/go3d/mat2x2d"
	"github.com/ungerik/go3d/mat3x3d"
	"github.com/ungerik/go3d/quaterniond"
	"github.com/ungerik/go3d/vec3d"
	"github.com/ungerik/go3d/vec4d"
)

var (
	Zero  = T{}
	Ident = T{
		vec4d.T{1, 0, 0, 0},
		vec4d.T{0, 1, 0, 0},
		vec4d.T{0, 0, 1, 0},
		vec4d.T{0, 0, 0, 1},
	}
)

type T [4]vec4d.T

// From copies a T from a generic.T implementation.
func From(other genericd.T) T {
	r := Ident
	cols := other.Cols()
	rows := other.Rows()
	if !((cols == 2 && rows == 2) || (cols == 3 && rows == 3) || (cols == 4 && rows == 4)) {
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
		"%f %f %f %f %f %f %f %f %f %f %f %f %f %f %f %f",
		&r[0][0], &r[0][1], &r[0][2], &r[0][3],
		&r[1][0], &r[1][1], &r[1][2], &r[1][3],
		&r[2][0], &r[2][1], &r[2][2], &r[2][3],
		&r[3][0], &r[3][1], &r[3][2], &r[3][3],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (self *T) String() string {
	return fmt.Sprintf("%s %s %s %s", self[0].String(), self[1].String(), self[2].String(), self[3].String())
}

// Rows returns the number of rows of the matrix.
func (self *T) Rows() int {
	return 4
}

// Cols returns the number of columns of the matrix.
func (self *T) Cols() int {
	return 4
}

// Size returns the number elements of the matrix.
func (self *T) Size() int {
	return 16
}

// Slice returns the elements of the matrix as slice.
func (self *T) Slice() []float64 {
	return []float64{
		self[0][0], self[0][1], self[0][2], self[0][3],
		self[1][0], self[1][1], self[1][2], self[1][3],
		self[2][0], self[2][1], self[2][2], self[2][3],
		self[3][0], self[3][1], self[3][2], self[3][3],
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

func (self *T) Trace() float64 {
	return self[0][0] + self[1][1] + self[2][2] + self[3][3]
}

func (self *T) Trace3() float64 {
	return self[0][0] + self[1][1] + self[2][2]
}

func (self *T) AssignMat2x2(m *mat2x2d.T) *T {
	*self = T{
		vec4d.T{m[0][0], m[1][0], 0, 0},
		vec4d.T{m[0][1], m[1][1], 0, 0},
		vec4d.T{0, 0, 1, 0},
		vec4d.T{0, 0, 0, 1},
	}
	return self
}

func (self *T) AssignMat3x3(m *mat3x3d.T) *T {
	*self = T{
		vec4d.T{m[0][0], m[1][0], m[2][0], 0},
		vec4d.T{m[0][1], m[1][1], m[2][1], 0},
		vec4d.T{m[0][2], m[1][2], m[2][2], 0},
		vec4d.T{0, 0, 0, 1},
	}
	return self
}

func (self *T) AssignMul(a, b *T) *T {
	self[0] = a.MulVec4(&b[0])
	self[1] = a.MulVec4(&b[1])
	self[2] = a.MulVec4(&b[2])
	self[3] = a.MulVec4(&b[3])
	return self
}

func (self *T) MulVec4(vec *vec4d.T) vec4d.T {
	return vec4d.T{
		self[0][0]*vec[0] + self[1][0]*vec[1] + self[2][0]*vec[2] + self[3][0]*vec[3],
		self[0][1]*vec[1] + self[1][1]*vec[1] + self[2][1]*vec[2] + self[3][1]*vec[3],
		self[0][2]*vec[2] + self[1][2]*vec[1] + self[2][2]*vec[2] + self[3][2]*vec[3],
		self[0][3]*vec[3] + self[1][3]*vec[1] + self[2][3]*vec[2] + self[3][3]*vec[3],
	}
}

func (self *T) MulVec3(v *vec3d.T) vec3d.T {
	v4 := vec4d.FromVec3(v)
	v4 = self.MulVec4(&v4)
	return v4.Vec3DividedByW()
}

func (self *T) SetTranslation(v *vec3d.T) *T {
	self[3][0] = v[0]
	self[3][1] = v[1]
	self[3][2] = v[2]
	return self
}

func (self *T) Translate(v *vec3d.T) *T {
	self[3][0] += v[0]
	self[3][1] += v[1]
	self[3][2] += v[2]
	return self
}

func (self *T) TranslateX(d float64) *T {
	self[3][0] += d
	return self
}

func (self *T) TranslateY(d float64) *T {
	self[3][1] += d
	return self
}

func (self *T) TranslateZ(d float64) *T {
	self[3][2] += d
	return self
}

func (self *T) Scaling() vec4d.T {
	return vec4d.T{self[0][0], self[1][1], self[2][2], self[3][3]}
}

func (self *T) SetScaling(s *vec4d.T) *T {
	self[0][0] = s[0]
	self[1][1] = s[1]
	self[2][2] = s[2]
	self[3][3] = s[3]
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
	self[3][0] = 0

	self[0][1] = xy + wz
	self[1][1] = 1 - (xx + zz)
	self[2][1] = yz - wx
	self[3][1] = 0

	self[0][2] = xz - wy
	self[1][2] = yz + wx
	self[2][2] = 1 - (xx + yy)
	self[3][2] = 0

	self[0][3] = 0
	self[1][3] = 0
	self[2][3] = 0
	self[3][3] = 1

	return self
}

func (self *T) AssignXRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	self[0][0] = 1
	self[1][0] = 0
	self[2][0] = 0
	self[3][0] = 0

	self[0][1] = 0
	self[1][1] = cosine
	self[2][1] = -sine
	self[3][1] = 0

	self[0][2] = 0
	self[1][2] = sine
	self[2][2] = cosine
	self[3][2] = 0

	self[0][3] = 0
	self[1][3] = 0
	self[2][3] = 0
	self[3][3] = 1

	return self
}

func (self *T) AssignYRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	self[0][0] = cosine
	self[1][0] = 0
	self[2][0] = sine
	self[3][0] = 0

	self[0][1] = 0
	self[1][1] = 1
	self[2][1] = 0
	self[3][1] = 0

	self[0][2] = -sine
	self[1][2] = 0
	self[2][2] = cosine
	self[3][2] = 0

	self[0][3] = 0
	self[1][3] = 0
	self[2][3] = 0
	self[3][3] = 1

	return self
}

func (self *T) AssignZRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	self[0][0] = cosine
	self[1][0] = -sine
	self[2][0] = 0
	self[3][0] = 0

	self[0][1] = sine
	self[1][1] = cosine
	self[2][1] = 0
	self[3][1] = 0

	self[0][2] = 0
	self[1][2] = 0
	self[2][2] = 1
	self[3][2] = 0

	self[0][3] = 0
	self[1][3] = 0
	self[2][3] = 0
	self[3][3] = 1

	return self
}

func (self *T) AssignCoordinateSystem(x, y, z *vec3d.T) *T {
	self[0][0] = x[0]
	self[1][0] = x[1]
	self[2][0] = x[2]
	self[3][0] = 0

	self[0][1] = y[0]
	self[1][1] = y[1]
	self[2][1] = y[2]
	self[3][1] = 0

	self[0][2] = z[0]
	self[1][2] = z[1]
	self[2][2] = z[2]
	self[3][2] = 0

	self[0][3] = 0
	self[1][3] = 0
	self[2][3] = 0
	self[3][3] = 1

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
	self[3][0] = 0

	self[0][1] = sinR*cosH + cosR*sinP*sinH
	self[1][1] = cosR * cosP
	self[2][1] = sinR*sinH - cosR*sinP*cosH
	self[3][1] = 0

	self[0][2] = -cosP * sinH
	self[1][2] = sinP
	self[2][2] = cosP * cosH
	self[3][2] = 0

	self[0][3] = 0
	self[1][3] = 0
	self[2][3] = 0
	self[3][3] = 1

	return self
}

func (self *T) AssignPerspectiveProjection(left, right, bottom, top, znear, zfar float64) *T {
	near2 := znear + znear
	oo_far_near := 1 / (zfar - znear)

	self[0][0] = near2 / (right - left)
	self[1][0] = 0
	self[2][0] = (right + left) / (right - left)
	self[3][0] = 0

	self[0][1] = 0
	self[1][1] = near2 / (top - bottom)
	self[2][1] = (top + bottom) / (top - bottom)
	self[3][1] = 0

	self[0][2] = 0
	self[1][2] = 0
	self[2][2] = -(zfar + znear) * oo_far_near
	self[3][2] = -2 * zfar * znear * oo_far_near

	self[0][3] = 0
	self[1][3] = 0
	self[2][3] = -1
	self[3][3] = 0

	return self
}

func (self *T) AssignOrthogonalProjection(left, right, bottom, top, znear, zfar float64) *T {
	oo_right_left := 1 / (right - left)
	oo_top_bottom := 1 / (top - bottom)
	oo_far_near := 1 / (zfar - znear)

	self[0][0] = 2 * oo_right_left
	self[1][0] = 0
	self[2][0] = 0
	self[3][0] = -(right + left) * oo_right_left

	self[0][1] = 0
	self[1][1] = 2 * oo_top_bottom
	self[2][1] = 0
	self[3][1] = -(top + bottom) * oo_top_bottom

	self[0][2] = 0
	self[1][2] = 0
	self[2][2] = -2 * oo_far_near
	self[3][2] = -(zfar + znear) * oo_far_near

	self[0][3] = 0
	self[1][3] = 0
	self[2][3] = 0
	self[3][3] = 1

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

func (self *T) Determinant3x3() float64 {
	return self[0][0]*self[1][1]*self[2][2] +
		self[1][0]*self[2][1]*self[0][2] +
		self[2][0]*self[0][1]*self[1][2] -
		self[2][0]*self[1][1]*self[0][2] -
		self[1][0]*self[0][1]*self[2][2] -
		self[0][0]*self[2][1]*self[1][2]
}

func (self *T) IsReflective() bool {
	return self.Determinant3x3() < 0
}

func swap(a, b *float64) {
	temp := *a
	*a = *b
	*b = temp
}

func (self *T) Transpose() *T {
	swap(&self[3][0], &self[0][3])
	swap(&self[3][1], &self[1][3])
	swap(&self[3][2], &self[2][3])
	return self.Transpose3x3()
}

func (self *T) Transpose3x3() *T {
	swap(&self[1][0], &self[0][1])
	swap(&self[2][0], &self[0][2])
	swap(&self[2][1], &self[1][2])
	return self
}
