package quaternion

import (
	"fmt"
	"math"
	"testing"

	quaternion64 "github.com/ungerik/go3d/float64/quaternion"
	vec364 "github.com/ungerik/go3d/float64/vec3"
	"github.com/ungerik/go3d/fmath"
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

func TestToEulerAngles(t *testing.T) {
	specialValues := []float64{-5, -math.Pi, -2, -math.Pi / 2, 0, math.Pi / 2, 2.4, math.Pi, 3.9}
	for _, x := range specialValues {
		for _, y := range specialValues {
			for _, z := range specialValues {
				quat1 := quaternion64.FromEulerAngles(y, x, z)
				ry, rx, rz := quat1.ToEulerAngles()
				quat2 := quaternion64.FromEulerAngles(ry, rx, rz)
				// quat must be equivalent
				const e64 = 1e-14
				cond1 := math.Abs(quat1[0]-quat2[0]) < e64 && math.Abs(quat1[1]-quat2[1]) < e64 && math.Abs(quat1[2]-quat2[2]) < e64 && math.Abs(quat1[3]-quat2[3]) < e64
				cond2 := math.Abs(quat1[0]+quat2[0]) < e64 && math.Abs(quat1[1]+quat2[1]) < e64 && math.Abs(quat1[2]+quat2[2]) < e64 && math.Abs(quat1[3]+quat2[3]) < e64
				if !cond1 && !cond2 {
					fmt.Printf("test case %v, %v, %v failed\n", x, y, z)
					fmt.Printf("result is %v, %v, % v\n", rx, ry, rz)
					fmt.Printf("quat1 is %v\n", quat1)
					fmt.Printf("quat2 is %v\n", quat2)
					t.Fail()
				}
				x32 := float32(x)
				y32 := float32(y)
				z32 := float32(z)
				quat132 := FromEulerAngles(x32, y32, z32)
				ry32, rx32, rz32 := quat132.ToEulerAngles()
				quat232 := FromEulerAngles(ry32, rx32, rz32)
				// quat must be equivalent
				const e32 = 1e-6
				cond1 = fmath.Abs(quat132[0]-quat232[0]) < e32 && fmath.Abs(quat132[1]-quat232[1]) < e32 && fmath.Abs(quat132[2]-quat232[2]) < e32 && fmath.Abs(quat132[3]-quat232[3]) < e32
				cond2 = fmath.Abs(quat132[0]+quat232[0]) < e32 && fmath.Abs(quat132[1]+quat232[1]) < e32 && fmath.Abs(quat132[2]+quat232[2]) < e32 && fmath.Abs(quat132[3]+quat232[3]) < e32
				if !cond1 && !cond2 {
					fmt.Printf("test case %v, %v, %v failed\n", x32, y32, z32)
					fmt.Printf("result is %v, %v, % v\n", rx32, ry32, rz32)
					fmt.Printf("quat1 is %v\n", quat132)
					fmt.Printf("quat2 is %v\n", quat232)
					t.Fail()
				}
			}
		}
	}
}
