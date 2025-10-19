package vec2

import (
	"math"
	"strconv"
	"testing"
)

func TestAbs(t *testing.T) {
	v1 := T{1.5, -2.6}

	expectedV1 := T{1.5, 2.6}
	expectedV2 := T{1.5, 2.6}

	v2 := v1.Abs()

	if v1 != expectedV1 {
		t.Fail()
	}
	if *v2 != expectedV2 {
		t.Fail()
	}
}

func TestAbsed(t *testing.T) {
	v1 := T{1.5, -2.6}

	expectedV1 := T{1.5, -2.6}
	expectedV2 := T{1.5, 2.6}

	v2 := v1.Absed()

	if v1 != expectedV1 {
		t.Fail()
	}
	if v2 != expectedV2 {
		t.Fail()
	}
}

func TestNormal(t *testing.T) {
	v1 := T{1.0, 1.0}

	v1Length := math.Sqrt(1*1 + 1*1)
	expectedV1n := T{1 / v1Length, -1 / v1Length}

	v1n := v1.Normal()

	if v1n != expectedV1n {
		t.Fail()
	}
}

func TestNormal2(t *testing.T) {
	v1 := T{4, 6}

	v1Length := math.Sqrt(4*4 + 6*6)
	expectedV1n := T{6.0 / v1Length, -4 / v1Length}

	v1n := v1.Normal()

	if v1n != expectedV1n {
		t.Fail()
	}
}

func TestNormalCCW(t *testing.T) {
	v1 := T{1.0, 1.0}

	v1Length := math.Sqrt(1*1 + 1*1)
	expectedV1n := T{-1 / v1Length, 1 / v1Length}

	v1n := v1.NormalCCW()

	if v1n != expectedV1n {
		t.Fail()
	}
}

func TestNormalCCW2(t *testing.T) {
	v1 := T{4, 6}

	v1Length := math.Sqrt(4*4 + 6*6)
	expectedV1n := T{-6.0 / v1Length, 4 / v1Length}

	v1n := v1.NormalCCW()

	if v1n != expectedV1n {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	v1 := T{1, 2}
	v2 := T{9, 8}

	expectedV1 := T{10, 10}
	expectedV2 := T{9, 8}
	expectedV3 := T{10, 10}

	v3 := v1.Add(&v2)

	if v1 != expectedV1 {
		t.Fail()
	}
	if v2 != expectedV2 {
		t.Fail()
	}
	if *v3 != expectedV3 {
		t.Fail()
	}
}

func TestAdded(t *testing.T) {
	v1 := T{1, 2}
	v2 := T{9, 8}

	expectedV1 := T{1, 2}
	expectedV2 := T{9, 8}
	expectedV3 := T{10, 10}

	v3 := v1.Added(&v2)

	if v1 != expectedV1 {
		t.Fail()
	}
	if v2 != expectedV2 {
		t.Fail()
	}
	if v3 != expectedV3 {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	v1 := T{1, 2}
	v2 := T{9, 8}

	expectedV1 := T{-8, -6}
	expectedV2 := T{9, 8}
	expectedV3 := T{-8, -6}

	v3 := v1.Sub(&v2)

	if v1 != expectedV1 {
		t.Fail()
	}
	if v2 != expectedV2 {
		t.Fail()
	}
	if *v3 != expectedV3 {
		t.Fail()
	}
}

func TestSubed(t *testing.T) {
	v1 := T{1, 2}
	v2 := T{9, 8}

	expectedV1 := T{1, 2}
	expectedV2 := T{9, 8}
	expectedV3 := T{-8, -6}

	v3 := v1.Subed(&v2)

	if v1 != expectedV1 {
		t.Fail()
	}
	if v2 != expectedV2 {
		t.Fail()
	}
	if v3 != expectedV3 {
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	v1 := T{1, 2}
	v2 := T{9, 8}

	expectedV1 := T{9, 16}
	expectedV2 := T{9, 8}
	expectedV3 := T{9, 16}

	v3 := v1.Mul(&v2)

	if v1 != expectedV1 {
		t.Fail()
	}
	if v2 != expectedV2 {
		t.Fail()
	}
	if *v3 != expectedV3 {
		t.Fail()
	}
}

func TestMuled(t *testing.T) {
	v1 := T{1, 2}
	v2 := T{9, 8}

	expectedV1 := T{1, 2}
	expectedV2 := T{9, 8}
	expectedV3 := T{9, 16}

	v3 := v1.Muled(&v2)

	if v1 != expectedV1 {
		t.Fail()
	}
	if v2 != expectedV2 {
		t.Fail()
	}
	if v3 != expectedV3 {
		t.Fail()
	}
}

func TestAngle(t *testing.T) {
	radFor45deg := math.Pi / 4.0
	testSetups := []struct {
		a, b          T
		expectedAngle float64
		name          string
	}{
		{a: T{1, 0}, b: T{1, 0}, expectedAngle: 0 * radFor45deg, name: "0/360 degree angle, equal/parallell vectors"},
		{a: T{1, 0}, b: T{1, 1}, expectedAngle: 1 * radFor45deg, name: "45 degree angle"},
		{a: T{1, 0}, b: T{0, 1}, expectedAngle: 2 * radFor45deg, name: "90 degree angle, orthogonal vectors"},
		{a: T{1, 0}, b: T{-1, 1}, expectedAngle: 3 * radFor45deg, name: "135 degree angle"},
		{a: T{1, 0}, b: T{-1, 0}, expectedAngle: 4 * radFor45deg, name: "180 degree angle, inverted/anti parallell vectors"},
		{a: T{1, 0}, b: T{-1, -1}, expectedAngle: (8 - 5) * radFor45deg, name: "225 degree angle"},
		{a: T{1, 0}, b: T{0, -1}, expectedAngle: (8 - 6) * radFor45deg, name: "270 degree angle, orthogonal vectors"},
		{a: T{1, 0}, b: T{1, -1}, expectedAngle: (8 - 7) * radFor45deg, name: "315 degree angle"},
	}

	for _, testSetup := range testSetups {
		t.Run(testSetup.name, func(t *testing.T) {
			angle := Angle(&testSetup.a, &testSetup.b)

			if !PracticallyEquals(angle, testSetup.expectedAngle, 0.00000001) {
				t.Errorf("Angle expected to be %f but was %f for test \"%s\".", testSetup.expectedAngle, angle, testSetup.name)
			}
		})
	}
}

func TestCosine(t *testing.T) {
	radFor45deg := math.Pi / 4.0
	testSetups := []struct {
		a, b           T
		expectedCosine float64
		name           string
	}{
		{a: T{1, 0}, b: T{1, 0}, expectedCosine: math.Cos(0 * radFor45deg), name: "0/360 degree angle, equal/parallell vectors"},
		{a: T{1, 0}, b: T{1, 1}, expectedCosine: math.Cos(1 * radFor45deg), name: "45 degree angle"},
		{a: T{1, 0}, b: T{0, 1}, expectedCosine: math.Cos(2 * radFor45deg), name: "90 degree angle, orthogonal vectors"},
		{a: T{1, 0}, b: T{-1, 1}, expectedCosine: math.Cos(3 * radFor45deg), name: "135 degree angle"},
		{a: T{1, 0}, b: T{-1, 0}, expectedCosine: math.Cos(4 * radFor45deg), name: "180 degree angle, inverted/anti parallell vectors"},
		{a: T{1, 0}, b: T{-1, -1}, expectedCosine: math.Cos(5 * radFor45deg), name: "225 degree angle"},
		{a: T{1, 0}, b: T{0, -1}, expectedCosine: math.Cos(6 * radFor45deg), name: "270 degree angle, orthogonal vectors"},
		{a: T{1, 0}, b: T{1, -1}, expectedCosine: math.Cos(7 * radFor45deg), name: "315 degree angle"},
	}

	for _, testSetup := range testSetups {
		t.Run(testSetup.name, func(t *testing.T) {
			cos := Cosine(&testSetup.a, &testSetup.b)

			if !PracticallyEquals(cos, testSetup.expectedCosine, 0.00000001) {
				t.Errorf("Cosine expected to be %f but was %f for test \"%s\".", testSetup.expectedCosine, cos, testSetup.name)
			}
		})
	}
}

func TestSinus(t *testing.T) {
	radFor45deg := math.Pi / 4.0
	testSetups := []struct {
		a, b         T
		expectedSine float64
		name         string
	}{
		{a: T{1, 0}, b: T{1, 0}, expectedSine: math.Sin(0 * radFor45deg), name: "0/360 degree angle, equal/parallell vectors"},
		{a: T{1, 0}, b: T{1, 1}, expectedSine: math.Sin(1 * radFor45deg), name: "45 degree angle"},
		{a: T{1, 0}, b: T{0, 1}, expectedSine: math.Sin(2 * radFor45deg), name: "90 degree angle, orthogonal vectors"},
		{a: T{1, 0}, b: T{-1, 1}, expectedSine: math.Sin(3 * radFor45deg), name: "135 degree angle"},
		{a: T{1, 0}, b: T{-1, 0}, expectedSine: math.Sin(4 * radFor45deg), name: "180 degree angle, inverted/anti parallell vectors"},
		{a: T{1, 0}, b: T{-1, -1}, expectedSine: math.Sin(5 * radFor45deg), name: "225 degree angle"},
		{a: T{1, 0}, b: T{0, -1}, expectedSine: math.Sin(6 * radFor45deg), name: "270 degree angle, orthogonal vectors"},
		{a: T{1, 0}, b: T{1, -1}, expectedSine: math.Sin(7 * radFor45deg), name: "315 degree angle"},
	}

	for _, testSetup := range testSetups {
		t.Run(testSetup.name, func(t *testing.T) {
			sin := Sinus(&testSetup.a, &testSetup.b)

			if !PracticallyEquals(sin, testSetup.expectedSine, 0.00000001) {
				t.Errorf("Sine expected to be %f but was %f for test \"%s\".", testSetup.expectedSine, sin, testSetup.name)
			}
		})
	}
}

func TestLeftRightWinding(t *testing.T) {
	a := T{1.0, 0.0}

	for angle := 0; angle <= 360; angle += 15 {
		rad := (math.Pi / 180.0) * float64(angle)

		bx := clampDecimals(math.Cos(rad), 4)
		by := clampDecimals(math.Sin(rad), 4)
		b := T{bx, by}

		t.Run("left winding angle "+strconv.Itoa(angle), func(t *testing.T) {
			lw := IsLeftWinding(&a, &b)
			rw := IsRightWinding(&a, &b)

			if angle%180 == 0 {
				// No winding at 0, 180 and 360 degrees
				if lw || rw {
					t.Errorf("Neither left or right winding should be true on angle %d. Left winding=%t, right winding=%t", angle, lw, rw)
				}
			} else if angle < 180 {
				// Left winding at 0 < angle < 180
				if !lw || rw {
					t.Errorf("Left winding should be true (not right winding) on angle %d. Left winding=%t, right winding=%t", angle, lw, rw)
				}
			} else if angle > 180 {
				// Right winding at 180 < angle < 360
				if lw || !rw {
					t.Errorf("Right winding should be true (not left winding) on angle %d. Left winding=%t, right winding=%t", angle, lw, rw)
				}
			}
		})
	}
}

func clampDecimals(decimalValue float64, amountDecimals float64) float64 {
	factor := math.Pow(10, amountDecimals)
	return math.Round(decimalValue*factor) / factor
}

func TestIsZeroEps(t *testing.T) {
	tests := []struct {
		name    string
		vec     T
		epsilon float64
		want    bool
	}{
		{"exact zero", T{0, 0}, 0.0001, true},
		{"within epsilon", T{0.00001, -0.00001}, 0.0001, true},
		{"at epsilon boundary", T{0.0001, 0.0001}, 0.0001, true},
		{"outside epsilon", T{0.001, 0}, 0.0001, false},
		{"one component outside", T{0.00001, 0.001}, 0.0001, false},
		{"negative outside epsilon", T{-0.001, 0}, 0.0001, false},
		{"large values", T{1.0, 2.0}, 0.0001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vec.IsZeroEps(tt.epsilon); got != tt.want {
				t.Errorf("IsZeroEps() = %v, want %v for vec %v with epsilon %v", got, tt.want, tt.vec, tt.epsilon)
			}
		})
	}
}
