package quaternion

import (
	"fmt"
	"testing"

	math "github.com/ungerik/go3d/fmath"
	"github.com/ungerik/go3d/vec3"
)

// RotateVec3 rotates v by the rotation represented by the quaternion.
func rotateAndNormalizeVec3(quat *T, v *vec3.T) {
	qv := T{v[0], v[1], v[2], 0}
	inv := quat.Inverted()
	q := Mul3(quat, &qv, &inv)
	v[0] = q[0]
	v[1] = q[1]
	v[2] = q[2]
}

func TestQuaternionRotateVec3(t *testing.T) {
	eularAngles := []vec3.T{
		{90, 20, 21},
		{-90, 0, 0},
		{28, 1043, -38},
	}
	vecs := []vec3.T{
		{2, 3, 4},
		{1, 3, -2},
		{-6, 2, 9},
	}
	for _, vec := range vecs {
		for _, eularAngle := range eularAngles {
			func() {
				q := FromEulerAngles(eularAngle[1]*math.Pi/180.0, eularAngle[0]*math.Pi/180.0, eularAngle[2]*math.Pi/180.0)
				vec_r1 := vec
				vec_r2 := vec
				magSqr := vec_r1.LengthSqr()
				rotateAndNormalizeVec3(&q, &vec_r2)
				q.RotateVec3(&vec_r1)
				vecd := q.RotatedVec3(&vec)
				magSqr2 := vec_r1.LengthSqr()

				if !vecd.PracticallyEquals(&vec_r1, 0.000001) {
					t.Logf("test case %v rotates %v failed - vector rotation: %+v, %+v\n", eularAngle, vec, vecd, vec_r1)
					t.Fail()
				}

				angle := vec3.Angle(&vec_r1, &vec_r2)
				length := math.Abs(magSqr - magSqr2)

				if angle > 0.001 {
					t.Logf("test case %v rotates %v failed - angle difference to large\n", eularAngle, vec)
					t.Logf("vectors: %+v, %+v\n", vec_r1, vec_r2)
					t.Logf("angle: %v\n", angle)
					t.Fail()
				}

				if length > 0.0001 {
					t.Logf("test case %v rotates %v failed - squared length difference to large\n", eularAngle, vec)
					t.Logf("vectors: %+v %+v\n", vec_r1, vec_r2)
					t.Logf("squared lengths: %v, %v\n", magSqr, magSqr2)
					t.Fail()
				}
			}()
		}
	}
}

func TestToEulerAngles(t *testing.T) {
	specialValues := []float32{-5, -math.Pi, -2, -math.Pi / 2, 0, math.Pi / 2, 2.4, math.Pi, 3.9}
	for _, x := range specialValues {
		for _, y := range specialValues {
			for _, z := range specialValues {
				quat1 := FromEulerAngles(y, x, z)
				ry, rx, rz := quat1.ToEulerAngles()
				quat2 := FromEulerAngles(ry, rx, rz)
				// quat must be equivalent
				const e32 = 1e-6
				cond1 := math.Abs(quat1[0]-quat2[0]) < e32 && math.Abs(quat1[1]-quat2[1]) < e32 && math.Abs(quat1[2]-quat2[2]) < e32 && math.Abs(quat1[3]-quat2[3]) < e32
				cond2 := math.Abs(quat1[0]+quat2[0]) < e32 && math.Abs(quat1[1]+quat2[1]) < e32 && math.Abs(quat1[2]+quat2[2]) < e32 && math.Abs(quat1[3]+quat2[3]) < e32
				if !cond1 && !cond2 {
					fmt.Printf("test case %v, %v, %v failed\n", x, y, z)
					fmt.Printf("result is %v, %v, %v\n", rx, ry, rz)
					fmt.Printf("quat1 is %v\n", quat1)
					fmt.Printf("quat2 is %v\n", quat2)
					t.Fail()
				}
			}
		}
	}
}
