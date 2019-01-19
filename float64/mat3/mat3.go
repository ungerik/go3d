// Package mat3 contains a 3x3 float64 matrix type T and functions.
package mat3

import (
	"fmt"
	"math"

	"github.com/gmlewis/go3d/float64/generic"
	"github.com/gmlewis/go3d/float64/mat2"
	"github.com/gmlewis/go3d/float64/quaternion"
	"github.com/gmlewis/go3d/float64/vec2"
	"github.com/gmlewis/go3d/float64/vec3"
)

var (
	// Zero holds a zero matrix.
	Zero = T{}

	// Ident holds an ident matrix.
	Ident = T{
		vec3.T{1, 0, 0},
		vec3.T{0, 1, 0},
		vec3.T{0, 0, 1},
	}
)

// T represents a 3x3 matrix.
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
	_, err = fmt.Sscan(s,
		&r[0][0], &r[0][1], &r[0][2],
		&r[1][0], &r[1][1], &r[1][2],
		&r[2][0], &r[2][1], &r[2][2],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (mat *T) String() string {
	return fmt.Sprintf("%s %s %s", mat[0].String(), mat[1].String(), mat[2].String())
}

// Rows returns the number of rows of the matrix.
func (mat *T) Rows() int {
	return 3
}

// Cols returns the number of columns of the matrix.
func (mat *T) Cols() int {
	return 3
}

// Size returns the number elements of the matrix.
func (mat *T) Size() int {
	return 9
}

// Slice returns the elements of the matrix as slice.
func (mat *T) Slice() []float64 {
	return mat.Array()[:]
}

// Get returns one element of the matrix.
func (mat *T) Get(col, row int) float64 {
	return mat[col][row]
}

// IsZero checks if all elements of the matrix are zero.
func (mat *T) IsZero() bool {
	return *mat == Zero
}

// Scale multiplies the diagonal scale elements by f returns mat.
func (mat *T) Scale(f float64) *T {
	mat[0][0] *= f
	mat[1][1] *= f
	mat[2][2] *= f
	return mat
}

// Scaled returns a copy of the matrix with the diagonal scale elements multiplied by f.
func (mat *T) Scaled(f float64) T {
	r := *mat
	return *r.Scale(f)
}

// Scaling returns the scaling diagonal of the matrix.
func (mat *T) Scaling() vec3.T {
	return vec3.T{mat[0][0], mat[1][1], mat[2][2]}
}

// SetScaling sets the scaling diagonal of the matrix.
func (mat *T) SetScaling(s *vec3.T) *T {
	mat[0][0] = s[0]
	mat[1][1] = s[1]
	mat[2][2] = s[2]
	return mat
}

// ScaleVec2 multiplies the 2D scaling diagonal of the matrix by s.
func (mat *T) ScaleVec2(s *vec2.T) *T {
	mat[0][0] *= s[0]
	mat[1][1] *= s[1]
	return mat
}

// SetTranslation sets the 2D translation elements of the matrix.
func (mat *T) SetTranslation(v *vec2.T) *T {
	mat[2][0] = v[0]
	mat[2][1] = v[1]
	return mat
}

// Translate adds v to the 2D translation part of the matrix.
func (mat *T) Translate(v *vec2.T) *T {
	mat[2][0] += v[0]
	mat[2][1] += v[1]
	return mat
}

// TranslateX adds dx to the 2D X-translation element of the matrix.
func (mat *T) TranslateX(dx float64) *T {
	mat[2][0] += dx
	return mat
}

// TranslateY adds dy to the 2D Y-translation element of the matrix.
func (mat *T) TranslateY(dy float64) *T {
	mat[2][1] += dy
	return mat
}

// Trace returns the trace value for the matrix.
func (mat *T) Trace() float64 {
	return mat[0][0] + mat[1][1] + mat[2][2]
}

// AssignMul multiplies a and b and assigns the result to mat.
func (mat *T) AssignMul(a, b *T) *T {
	mat[0] = a.MulVec3(&b[0])
	mat[1] = a.MulVec3(&b[1])
	mat[2] = a.MulVec3(&b[2])
	return mat
}

// AssignMat2x2 assigns a 2x2 sub-matrix and sets the rest of the matrix to the ident value.
func (mat *T) AssignMat2x2(m *mat2.T) *T {
	*mat = T{
		vec3.T{m[0][0], m[1][0], 0},
		vec3.T{m[0][1], m[1][1], 0},
		vec3.T{0, 0, 1},
	}
	return mat
}

// MulVec3 multiplies v with T.
func (mat *T) MulVec3(v *vec3.T) vec3.T {
	return vec3.T{
		mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2],
		mat[0][1]*v[1] + mat[1][1]*v[1] + mat[2][1]*v[2],
		mat[0][2]*v[2] + mat[1][2]*v[1] + mat[2][2]*v[2],
	}
}

// TransformVec3 multiplies v with mat and saves the result in v.
func (mat *T) TransformVec3(v *vec3.T) {
	// Use intermediate variables to not alter further computations.
	x := mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2]
	y := mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2]
	v[2] = mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2]
	v[0] = x
	v[1] = y
}

// Quaternion extracts a quaternion from the rotation part of the matrix.
func (mat *T) Quaternion() quaternion.T {
	tr := mat.Trace()

	s := math.Sqrt(tr + 1)
	w := s * 0.5
	s = 0.5 / s

	q := quaternion.T{
		(mat[1][2] - mat[2][1]) * s,
		(mat[2][0] - mat[0][2]) * s,
		(mat[0][1] - mat[1][0]) * s,
		w,
	}
	return q.Normalized()
}

// AssignQuaternion assigns a quaternion to the rotations part of the matrix and sets the other elements to their ident value.
func (mat *T) AssignQuaternion(q *quaternion.T) *T {
	xx := q[0] * q[0] * 2
	yy := q[1] * q[1] * 2
	zz := q[2] * q[2] * 2
	xy := q[0] * q[1] * 2
	xz := q[0] * q[2] * 2
	yz := q[1] * q[2] * 2
	wx := q[3] * q[0] * 2
	wy := q[3] * q[1] * 2
	wz := q[3] * q[2] * 2

	mat[0][0] = 1 - (yy + zz)
	mat[1][0] = xy - wz
	mat[2][0] = xz + wy

	mat[0][1] = xy + wz
	mat[1][1] = 1 - (xx + zz)
	mat[2][1] = yz - wx

	mat[0][2] = xz - wy
	mat[1][2] = yz + wx
	mat[2][2] = 1 - (xx + yy)

	return mat
}

// AssignXRotation assigns a rotation around the x axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *T) AssignXRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	mat[0][0] = 1
	mat[1][0] = 0
	mat[2][0] = 0

	mat[0][1] = 0
	mat[1][1] = cosine
	mat[2][1] = -sine

	mat[0][2] = 0
	mat[1][2] = sine
	mat[2][2] = cosine

	return mat
}

// AssignYRotation assigns a rotation around the y axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *T) AssignYRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	mat[0][0] = cosine
	mat[1][0] = 0
	mat[2][0] = sine

	mat[0][1] = 0
	mat[1][1] = 1
	mat[2][1] = 0

	mat[0][2] = -sine
	mat[1][2] = 0
	mat[2][2] = cosine

	return mat
}

// AssignZRotation assigns a rotation around the z axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *T) AssignZRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	mat[0][0] = cosine
	mat[1][0] = -sine
	mat[2][0] = 0

	mat[0][1] = sine
	mat[1][1] = cosine
	mat[2][1] = 0

	mat[0][2] = 0
	mat[1][2] = 0
	mat[2][2] = 1

	return mat
}

// AssignCoordinateSystem assigns the rotation of a orthogonal coordinates system to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *T) AssignCoordinateSystem(x, y, z *vec3.T) *T {
	mat[0][0] = x[0]
	mat[1][0] = x[1]
	mat[2][0] = x[2]

	mat[0][1] = y[0]
	mat[1][1] = y[1]
	mat[2][1] = y[2]

	mat[0][2] = z[0]
	mat[1][2] = z[1]
	mat[2][2] = z[2]

	return mat
}

// AssignEulerRotation assigns Euler angle rotations to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *T) AssignEulerRotation(yHead, xPitch, zRoll float64) *T {
	sinH := math.Sin(yHead)
	cosH := math.Cos(yHead)
	sinP := math.Sin(xPitch)
	cosP := math.Cos(xPitch)
	sinR := math.Sin(zRoll)
	cosR := math.Cos(zRoll)

	mat[0][0] = cosR*cosH - sinR*sinP*sinH
	mat[1][0] = -sinR * cosP
	mat[2][0] = cosR*sinH + sinR*sinP*cosH

	mat[0][1] = sinR*cosH + cosR*sinP*sinH
	mat[1][1] = cosR * cosP
	mat[2][1] = sinR*sinH - cosR*sinP*cosH

	mat[0][2] = -cosP * sinH
	mat[1][2] = sinP
	mat[2][2] = cosP * cosH

	return mat
}

// ExtractEulerAngles extracts the rotation part of the matrix as Euler angle rotation values.
func (mat *T) ExtractEulerAngles() (yHead, xPitch, zRoll float64) {
	xPitch = math.Asin(mat[1][2])
	f12 := math.Abs(mat[1][2])
	if f12 > (1.0-0.0001) && f12 < (1.0+0.0001) { // f12 == 1.0
		yHead = 0.0
		zRoll = math.Atan2(mat[0][1], mat[0][0])
	} else {
		yHead = math.Atan2(-mat[0][2], mat[2][2])
		zRoll = math.Atan2(-mat[1][0], mat[1][1])
	}
	return yHead, xPitch, zRoll
}

// Determinant returns the determinant of the matrix.
func (mat *T) Determinant() float64 {
	return mat[0][0]*mat[1][1]*mat[2][2] +
		mat[1][0]*mat[2][1]*mat[0][2] +
		mat[2][0]*mat[0][1]*mat[1][2] -
		mat[2][0]*mat[1][1]*mat[0][2] -
		mat[1][0]*mat[0][1]*mat[2][2] -
		mat[0][0]*mat[2][1]*mat[1][2]
}

// IsReflective returns true if the matrix can be reflected by a plane.
func (mat *T) IsReflective() bool {
	return mat.Determinant() < 0
}

func swap(a, b *float64) {
	temp := *a
	*a = *b
	*b = temp
}

// Transpose transposes the matrix.
func (mat *T) Transpose() *T {
	swap(&mat[1][0], &mat[0][1])
	swap(&mat[2][0], &mat[0][2])
	swap(&mat[2][1], &mat[1][2])
	return mat
}
