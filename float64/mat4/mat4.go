// Package mat4 contains a 4x4 float64 matrix type T and functions.
package mat4

import (
	"fmt"
	"math"

	"github.com/ungerik/go3d/float64/generic"
	"github.com/ungerik/go3d/float64/mat2"
	"github.com/ungerik/go3d/float64/mat3"
	"github.com/ungerik/go3d/float64/quaternion"
	"github.com/ungerik/go3d/float64/vec3"
	"github.com/ungerik/go3d/float64/vec4"
)

var (
	// Zero holds a zero matrix.
	Zero = T{}

	// Ident holds an ident matrix.
	Ident = T{
		vec4.T{1, 0, 0, 0},
		vec4.T{0, 1, 0, 0},
		vec4.T{0, 0, 1, 0},
		vec4.T{0, 0, 0, 1},
	}
)

// T represents a 4x4 matrix.
type T [4]vec4.T

// From copies a T from a generic.T implementation.
func From(other generic.T) T {
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
	_, err = fmt.Sscan(s,
		&r[0][0], &r[0][1], &r[0][2], &r[0][3],
		&r[1][0], &r[1][1], &r[1][2], &r[1][3],
		&r[2][0], &r[2][1], &r[2][2], &r[2][3],
		&r[3][0], &r[3][1], &r[3][2], &r[3][3],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (mat *T) String() string {
	return fmt.Sprintf("%s %s %s %s", mat[0].String(), mat[1].String(), mat[2].String(), mat[3].String())
}

// Rows returns the number of rows of the matrix.
func (mat *T) Rows() int {
	return 4
}

// Cols returns the number of columns of the matrix.
func (mat *T) Cols() int {
	return 4
}

// Size returns the number elements of the matrix.
func (mat *T) Size() int {
	return 16
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

// Trace returns the trace value for the matrix.
func (mat *T) Trace() float64 {
	return mat[0][0] + mat[1][1] + mat[2][2] + mat[3][3]
}

// Trace3 returns the trace value for the 3x3 sub-matrix.
func (mat *T) Trace3() float64 {
	return mat[0][0] + mat[1][1] + mat[2][2]
}

// AssignMat2x2 assigns a 2x2 sub-matrix and sets the rest of the matrix to the ident value.
func (mat *T) AssignMat2x2(m *mat2.T) *T {
	*mat = T{
		vec4.T{m[0][0], m[1][0], 0, 0},
		vec4.T{m[0][1], m[1][1], 0, 0},
		vec4.T{0, 0, 1, 0},
		vec4.T{0, 0, 0, 1},
	}
	return mat
}

// AssignMat3x3 assigns a 3x3 sub-matrix and sets the rest of the matrix to the ident value.
func (mat *T) AssignMat3x3(m *mat3.T) *T {
	*mat = T{
		vec4.T{m[0][0], m[1][0], m[2][0], 0},
		vec4.T{m[0][1], m[1][1], m[2][1], 0},
		vec4.T{m[0][2], m[1][2], m[2][2], 0},
		vec4.T{0, 0, 0, 1},
	}
	return mat
}

// AssignMul multiplies a and b and assigns the result to T.
func (mat *T) AssignMul(a, b *T) *T {
	mat[0] = a.MulVec4(&b[0])
	mat[1] = a.MulVec4(&b[1])
	mat[2] = a.MulVec4(&b[2])
	mat[3] = a.MulVec4(&b[3])
	return mat
}

// MulVec4 multiplies v with mat and returns a new vector v' = M * v.
func (mat *T) MulVec4(v *vec4.T) vec4.T {
	return vec4.T{
		mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2] + mat[3][0]*v[3],
		mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2] + mat[3][1]*v[3],
		mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2] + mat[3][2]*v[3],
		mat[0][3]*v[0] + mat[1][3]*v[1] + mat[2][3]*v[2] + mat[3][3]*v[3],
	}
}

// TransformVec4 multiplies v with mat and saves the result in v.
func (mat *T) TransformVec4(v *vec4.T) {
	// Use intermediate variables to not alter further computations.
	x := mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2] + mat[3][0]*v[3]
	y := mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2] + mat[3][1]*v[3]
	z := mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2] + mat[3][2]*v[3]
	v[3] = mat[0][3]*v[0] + mat[1][3]*v[1] + mat[2][3]*v[2] + mat[3][3]*v[3]
	v[0] = x
	v[1] = y
	v[2] = z
}

// MulVec3 multiplies v (converted to a vec4 as (v_1, v_2, v_3, 1))
// with mat and divides the result by w. Returns a new vec3.
func (mat *T) MulVec3(v *vec3.T) vec3.T {
	v4 := vec4.FromVec3(v)
	v4 = mat.MulVec4(&v4)
	return v4.Vec3DividedByW()
}

// TransformVec3 multiplies v (converted to a vec4 as (v_1, v_2, v_3, 1))
// with mat, divides the result by w and saves the result in v.
func (mat *T) TransformVec3(v *vec3.T) {
	x := mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2] + mat[3][0]
	y := mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2] + mat[3][1]
	z := mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2] + mat[3][2]
	w := mat[0][3]*v[0] + mat[1][3]*v[1] + mat[2][3]*v[2] + mat[3][3]
	oow := 1 / w
	v[0] = x * oow
	v[1] = y * oow
	v[2] = z * oow
}

// MulVec3W multiplies v with mat with w as fourth component of the vector.
// Useful to differentiate between vectors (w = 0) and points (w = 1)
// without transforming them to vec4.
func (mat *T) MulVec3W(v *vec3.T, w float64) vec3.T {
	result := *v
	mat.TransformVec3W(&result, w)
	return result
}

// TransformVec3W multiplies v with mat with w as fourth component of the vector and
// saves the result in v.
// Useful to differentiate between vectors (w = 0) and points (w = 1)
// without transforming them to vec4.
func (mat *T) TransformVec3W(v *vec3.T, w float64) {
	// use intermediate variables to not alter further computations
	x := mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2] + mat[3][0]*w
	y := mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2] + mat[3][1]*w
	v[2] = mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2] + mat[3][2]*w
	v[0] = x
	v[1] = y
}

// SetTranslation sets the translation elements of the matrix.
func (mat *T) SetTranslation(v *vec3.T) *T {
	mat[3][0] = v[0]
	mat[3][1] = v[1]
	mat[3][2] = v[2]
	return mat
}

// Translate adds v to the translation part of the matrix.
func (mat *T) Translate(v *vec3.T) *T {
	mat[3][0] += v[0]
	mat[3][1] += v[1]
	mat[3][2] += v[2]
	return mat
}

// TranslateX adds dx to the X-translation element of the matrix.
func (mat *T) TranslateX(dx float64) *T {
	mat[3][0] += dx
	return mat
}

// TranslateY adds dy to the Y-translation element of the matrix.
func (mat *T) TranslateY(dy float64) *T {
	mat[3][1] += dy
	return mat
}

// TranslateZ adds dz to the Z-translation element of the matrix.
func (mat *T) TranslateZ(dz float64) *T {
	mat[3][2] += dz
	return mat
}

// Scaling returns the scaling diagonal of the matrix.
func (mat *T) Scaling() vec4.T {
	return vec4.T{mat[0][0], mat[1][1], mat[2][2], mat[3][3]}
}

// SetScaling sets the scaling diagonal of the matrix.
func (mat *T) SetScaling(s *vec4.T) *T {
	mat[0][0] = s[0]
	mat[1][1] = s[1]
	mat[2][2] = s[2]
	mat[3][3] = s[3]
	return mat
}

// ScaleVec3 multiplies the scaling diagonal of the matrix by s.
func (mat *T) ScaleVec3(s *vec3.T) *T {
	mat[0][0] *= s[0]
	mat[1][1] *= s[1]
	mat[2][2] *= s[2]
	return mat
}

// Quaternion extracts a quaternion from the rotation part of the matrix.
func (mat *T) Quaternion() quaternion.T {
	// Use Trace3() for the 3x3 rotation part only (not the full 4x4 matrix)
	tr := mat.Trace3()

	var q quaternion.T

	// Use Shepperd's method to handle numerical stability
	// Pick the largest diagonal element to avoid division by small numbers
	if tr > 0 {
		// w is the largest component
		s := math.Sqrt(tr + 1)
		w := s * 0.5
		s = 0.5 / s
		q = quaternion.T{
			(mat[1][2] - mat[2][1]) * s,
			(mat[2][0] - mat[0][2]) * s,
			(mat[0][1] - mat[1][0]) * s,
			w,
		}
	} else if mat[0][0] > mat[1][1] && mat[0][0] > mat[2][2] {
		// x is the largest component
		s := math.Sqrt(1 + mat[0][0] - mat[1][1] - mat[2][2])
		x := s * 0.5
		s = 0.5 / s
		q = quaternion.T{
			x,
			(mat[0][1] + mat[1][0]) * s,
			(mat[2][0] + mat[0][2]) * s,
			(mat[1][2] - mat[2][1]) * s,
		}
	} else if mat[1][1] > mat[2][2] {
		// y is the largest component
		s := math.Sqrt(1 + mat[1][1] - mat[0][0] - mat[2][2])
		y := s * 0.5
		s = 0.5 / s
		q = quaternion.T{
			(mat[0][1] + mat[1][0]) * s,
			y,
			(mat[1][2] + mat[2][1]) * s,
			(mat[2][0] - mat[0][2]) * s,
		}
	} else {
		// z is the largest component
		s := math.Sqrt(1 + mat[2][2] - mat[0][0] - mat[1][1])
		z := s * 0.5
		s = 0.5 / s
		q = quaternion.T{
			(mat[2][0] + mat[0][2]) * s,
			(mat[1][2] + mat[2][1]) * s,
			z,
			(mat[0][1] - mat[1][0]) * s,
		}
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
	mat[3][0] = 0

	mat[0][1] = xy + wz
	mat[1][1] = 1 - (xx + zz)
	mat[2][1] = yz - wx
	mat[3][1] = 0

	mat[0][2] = xz - wy
	mat[1][2] = yz + wx
	mat[2][2] = 1 - (xx + yy)
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// AssignXRotation assigns a rotation around the x axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *T) AssignXRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	mat[0][0] = 1
	mat[1][0] = 0
	mat[2][0] = 0
	mat[3][0] = 0

	mat[0][1] = 0
	mat[1][1] = cosine
	mat[2][1] = -sine
	mat[3][1] = 0

	mat[0][2] = 0
	mat[1][2] = sine
	mat[2][2] = cosine
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// AssignYRotation assigns a rotation around the y axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *T) AssignYRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	mat[0][0] = cosine
	mat[1][0] = 0
	mat[2][0] = sine
	mat[3][0] = 0

	mat[0][1] = 0
	mat[1][1] = 1
	mat[2][1] = 0
	mat[3][1] = 0

	mat[0][2] = -sine
	mat[1][2] = 0
	mat[2][2] = cosine
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// AssignZRotation assigns a rotation around the z axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *T) AssignZRotation(angle float64) *T {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	mat[0][0] = cosine
	mat[1][0] = -sine
	mat[2][0] = 0
	mat[3][0] = 0

	mat[0][1] = sine
	mat[1][1] = cosine
	mat[2][1] = 0
	mat[3][1] = 0

	mat[0][2] = 0
	mat[1][2] = 0
	mat[2][2] = 1
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// AssignCoordinateSystem assigns the rotation of a orthogonal coordinates system to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *T) AssignCoordinateSystem(x, y, z *vec3.T) *T {
	mat[0][0] = x[0]
	mat[1][0] = x[1]
	mat[2][0] = x[2]
	mat[3][0] = 0

	mat[0][1] = y[0]
	mat[1][1] = y[1]
	mat[2][1] = y[2]
	mat[3][1] = 0

	mat[0][2] = z[0]
	mat[1][2] = z[1]
	mat[2][2] = z[2]
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

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
	mat[3][0] = 0

	mat[0][1] = sinR*cosH + cosR*sinP*sinH
	mat[1][1] = cosR * cosP
	mat[2][1] = sinR*sinH - cosR*sinP*cosH
	mat[3][1] = 0

	mat[0][2] = -cosP * sinH
	mat[1][2] = sinP
	mat[2][2] = cosP * cosH
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

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

// AssignFrustum assigns a frustum projection transformation.
// This creates an asymmetric perspective projection with explicit left, right, bottom, top planes.
// For a typical symmetric perspective projection, use AssignPerspective instead.
func (mat *T) AssignFrustum(left, right, bottom, top, znear, zfar float64) *T {
	near2 := znear + znear
	ooFarNear := 1 / (zfar - znear)

	mat[0][0] = near2 / (right - left)
	mat[1][0] = 0
	mat[2][0] = (right + left) / (right - left)
	mat[3][0] = 0

	mat[0][1] = 0
	mat[1][1] = near2 / (top - bottom)
	mat[2][1] = (top + bottom) / (top - bottom)
	mat[3][1] = 0

	mat[0][2] = 0
	mat[1][2] = 0
	mat[2][2] = -(zfar + znear) * ooFarNear
	mat[3][2] = -2 * zfar * znear * ooFarNear

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = -1
	mat[3][3] = 0

	return mat
}

// AssignPerspective assigns a symmetric perspective projection transformation.
// This is the typical perspective projection using field of view and aspect ratio.
// For asymmetric projections, use AssignFrustum instead.
//
// Parameters:
//   - fovy: Field of view in the y direction, in radians
//   - aspect: Aspect ratio (width / height)
//   - znear: Distance to near clipping plane (must be positive)
//   - zfar: Distance to far clipping plane (must be positive and > znear)
func (mat *T) AssignPerspective(fovy, aspect, znear, zfar float64) *T {
	tanHalfFovy := math.Tan(fovy / 2)
	ooFarNear := 1 / (zfar - znear)

	mat[0][0] = 1 / (aspect * tanHalfFovy)
	mat[1][0] = 0
	mat[2][0] = 0
	mat[3][0] = 0

	mat[0][1] = 0
	mat[1][1] = 1 / tanHalfFovy
	mat[2][1] = 0
	mat[3][1] = 0

	mat[0][2] = 0
	mat[1][2] = 0
	mat[2][2] = -(zfar + znear) * ooFarNear
	mat[3][2] = -2 * zfar * znear * ooFarNear

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = -1
	mat[3][3] = 0

	return mat
}

// AssignOrthogonalProjection assigns an orthogonal projection transformation.
func (mat *T) AssignOrthogonalProjection(left, right, bottom, top, znear, zfar float64) *T {
	ooRightLeft := 1 / (right - left)
	ooTopBottom := 1 / (top - bottom)
	ooFarNear := 1 / (zfar - znear)

	mat[0][0] = 2 * ooRightLeft
	mat[1][0] = 0
	mat[2][0] = 0
	mat[3][0] = -(right + left) * ooRightLeft

	mat[0][1] = 0
	mat[1][1] = 2 * ooTopBottom
	mat[2][1] = 0
	mat[3][1] = -(top + bottom) * ooTopBottom

	mat[0][2] = 0
	mat[1][2] = 0
	mat[2][2] = -2 * ooFarNear
	mat[3][2] = -(zfar + znear) * ooFarNear

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// Determinant3x3 returns the determinant of the 3x3 sub-matrix.
func (mat *T) Determinant3x3() float64 {
	return mat[0][0]*mat[1][1]*mat[2][2] +
		mat[1][0]*mat[2][1]*mat[0][2] +
		mat[2][0]*mat[0][1]*mat[1][2] -
		mat[2][0]*mat[1][1]*mat[0][2] -
		mat[1][0]*mat[0][1]*mat[2][2] -
		mat[0][0]*mat[2][1]*mat[1][2]
}

// IsReflective returns true if the matrix can be reflected by a plane.
func (mat *T) IsReflective() bool {
	return mat.Determinant3x3() < 0
}

func swap(a, b *float64) {
	*a, *b = *b, *a
}

// Transpose transposes the matrix.
func (mat *T) Transpose() *T {
	swap(&mat[3][0], &mat[0][3])
	swap(&mat[3][1], &mat[1][3])
	swap(&mat[3][2], &mat[2][3])
	return mat.Transpose3x3()
}

// Transpose3x3 transposes the 3x3 sub-matrix.
func (mat *T) Transpose3x3() *T {
	swap(&mat[1][0], &mat[0][1])
	swap(&mat[2][0], &mat[0][2])
	swap(&mat[2][1], &mat[1][2])
	return mat
}
