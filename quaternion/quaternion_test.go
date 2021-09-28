package quaternion

import (
	"fmt"
	"math"
	"testing"

	quaternion64 "github.com/ungerik/go3d/float64/quaternion"
	vec364 "github.com/ungerik/go3d/float64/vec3"
	"github.com/ungerik/go3d/vec3"
)

// RotateVec3 rotates v by the rotation represented by the quaternion.
func rotateAndNormalizeVec332(quat *T, v *vec3.T) {
	qv := T{v[0], v[1], v[2], 0}
	inv := quat.Inverted()
	q := Mul3(quat, &qv, &inv)
	v[0] = q[0]
	v[1] = q[1]
	v[2] = q[2]
}

// RotateVec3 rotates v by the rotation represented by the quaternion.
func rotateAndNormalizeVec364(quat *quaternion64.T, v *vec364.T) {
	qv := quaternion64.T{v[0], v[1], v[2], 0}
	inv := quat.Inverted()
	q := quaternion64.Mul3(quat, &qv, &inv)
	v[0] = q[0]
	v[1] = q[1]
	v[2] = q[2]
}

func TestQuaternionRotateVec3(t *testing.T) {
	eularAngles := []vec364.T{
		{90, 20, 21},
		{-90, 0, 0},
		{28, 1043, -38},
	}
	vecs := []vec364.T{
		{2, 3, 4},
		{1, 3, -2},
		{-6, 2, 9},
	}
	for _, vec := range vecs {
		for _, eularAngle := range eularAngles {
			func() {
				q := quaternion64.FromEulerAngles(eularAngle[1]*math.Pi/180, eularAngle[0]*math.Pi/180, eularAngle[2]*math.Pi/180)
				vec_r1 := vec
				vec_r2 := vec
				magSqr := vec_r1.LengthSqr()
				rotateAndNormalizeVec364(&q, &vec_r2)
				q.RotateVec3(&vec_r1)
				vecd := q.RotatedVec3(&vec)
				magSqr2 := vec_r1.LengthSqr()
				if vecd != vec_r1 {
					t.Fail()
				}
				if vec364.Angle(&vec_r1, &vec_r2) > 0.00000001 {
					fmt.Printf("test case %v rotates %v failed\n", eularAngle, vec)
					fmt.Println(vec_r1, vec_r2)
					fmt.Println(vec364.Angle(&vec_r1, &vec_r2))
					t.Fail()
				}
				if math.Abs(float64(magSqr-magSqr2)) > 0.000000000001 {
					fmt.Printf("test case %v rotates %v failed\n", eularAngle, vec)
					fmt.Println(vec_r1, vec_r2)
					fmt.Println(magSqr, magSqr2)
					t.Fail()
				}
			}()
			func() {
				q := FromEulerAngles(float32(eularAngle[1]*math.Pi/180), float32(eularAngle[0]*math.Pi/180), float32(eularAngle[2]*math.Pi/180))
				vec32 := vec3.T{float32(vec[0]), float32(vec[1]), float32(vec[2])}
				vec_r1 := vec32
				vec_r2 := vec32
				magSqr := vec_r1.LengthSqr()
				rotateAndNormalizeVec332(&q, &vec_r2)
				q.RotateVec3(&vec_r1)
				vecd := q.RotatedVec3(&vec32)
				magSqr2 := vec_r1.LengthSqr()
				if vecd != vec_r1 {
					t.Fail()
				}
				if vec3.Angle(&vec_r1, &vec_r2) > 0.001 {
					fmt.Printf("test case %v rotates %v failed\n", eularAngle, vec)
					fmt.Println(vec_r1, vec_r2)
					fmt.Println(vec3.Angle(&vec_r1, &vec_r2))
					t.Fail()
				}
				if math.Abs(float64(magSqr-magSqr2)) > 0.0001 {
					fmt.Printf("test case %v rotates %v failed\n", eularAngle, vec)
					fmt.Println(vec_r1, vec_r2)
					fmt.Println(magSqr, magSqr2)
					t.Fail()
				}
			}()
		}
	}
}
