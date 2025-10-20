package vec3

import (
	math "github.com/chewxy/math32"

	"testing"
)

func TestBoxIntersection(t *testing.T) {
	bb1 := Box{T{0, 0, 0}, T{1, 1, 1}}
	bb2 := Box{T{1, 1, 1}, T{2, 2, 2}}
	if !bb1.Intersects(&bb2) {
		t.Fail()
	}

	bb3 := Box{T{1, 2, 1}, T{2, 3, 2}}
	if bb1.Intersects(&bb3) {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	v1 := T{1, 2, 3}
	v2 := T{9, 8, 7}

	expectedV1 := T{10, 10, 10}
	expectedV2 := T{9, 8, 7}
	expectedV3 := T{10, 10, 10}

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
	v1 := T{1, 2, 3}
	v2 := T{9, 8, 7}

	expectedV1 := T{1, 2, 3}
	expectedV2 := T{9, 8, 7}
	expectedV3 := T{10, 10, 10}

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
	v1 := T{1, 2, 3}
	v2 := T{9, 8, 7}

	expectedV1 := T{-8, -6, -4}
	expectedV2 := T{9, 8, 7}
	expectedV3 := T{-8, -6, -4}

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
	v1 := T{1, 2, 3}
	v2 := T{9, 8, 7}

	expectedV1 := T{1, 2, 3}
	expectedV2 := T{9, 8, 7}
	expectedV3 := T{-8, -6, -4}

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
	v1 := T{1, 2, 3}
	v2 := T{9, 8, 7}

	expectedV1 := T{9, 16, 21}
	expectedV2 := T{9, 8, 7}
	expectedV3 := T{9, 16, 21}

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
	v1 := T{1, 2, 3}
	v2 := T{9, 8, 7}

	expectedV1 := T{1, 2, 3}
	expectedV2 := T{9, 8, 7}
	expectedV3 := T{9, 16, 21}

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
	radFor45deg := float32(math.Pi / 4.0)
	testSetups := []struct {
		a, b          T
		expectedAngle float32
		name          string
	}{
		{a: T{1, 0, 0}, b: T{1, 0, 0}, expectedAngle: 0 * radFor45deg, name: "0/360 degree angle, equal/parallell vectors"},
		{a: T{1, 0, 0}, b: T{1, 1, 0}, expectedAngle: 1 * radFor45deg, name: "45 degree angle"},
		{a: T{1, 0, 0}, b: T{0, 1, 0}, expectedAngle: 2 * radFor45deg, name: "90 degree angle, orthogonal vectors"},
		{a: T{1, 0, 0}, b: T{-1, 1, 0}, expectedAngle: 3 * radFor45deg, name: "135 degree angle"},
		{a: T{1, 0, 0}, b: T{-1, 0, 0}, expectedAngle: 4 * radFor45deg, name: "180 degree angle, inverted/anti parallell vectors"},
		{a: T{1, 0, 0}, b: T{-1, -1, 0}, expectedAngle: (8 - 5) * radFor45deg, name: "225 degree angle"},
		{a: T{1, 0, 0}, b: T{0, -1, 0}, expectedAngle: (8 - 6) * radFor45deg, name: "270 degree angle, orthogonal vectors"},
		{a: T{1, 0, 0}, b: T{1, -1, 0}, expectedAngle: (8 - 7) * radFor45deg, name: "315 degree angle"},
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
	radFor45deg := float32(math.Pi / 4.0)
	testSetups := []struct {
		a, b           T
		expectedCosine float32
		name           string
	}{
		{a: T{1, 0, 0}, b: T{1, 0, 0}, expectedCosine: math.Cos(0 * radFor45deg), name: "0/360 degree angle, equal/parallell vectors"},
		{a: T{1, 0, 0}, b: T{1, 1, 0}, expectedCosine: math.Cos(1 * radFor45deg), name: "45 degree angle"},
		{a: T{1, 0, 0}, b: T{0, 1, 0}, expectedCosine: math.Cos(2 * radFor45deg), name: "90 degree angle, orthogonal vectors"},
		{a: T{1, 0, 0}, b: T{-1, 1, 0}, expectedCosine: math.Cos(3 * radFor45deg), name: "135 degree angle"},
		{a: T{1, 0, 0}, b: T{-1, 0, 0}, expectedCosine: math.Cos(4 * radFor45deg), name: "180 degree angle, inverted/anti parallell vectors"},
		{a: T{1, 0, 0}, b: T{-1, -1, 0}, expectedCosine: math.Cos(5 * radFor45deg), name: "225 degree angle"},
		{a: T{1, 0, 0}, b: T{0, -1, 0}, expectedCosine: math.Cos(6 * radFor45deg), name: "270 degree angle, orthogonal vectors"},
		{a: T{1, 0, 0}, b: T{1, -1, 0}, expectedCosine: math.Cos(7 * radFor45deg), name: "315 degree angle"},
	}

	for _, testSetup := range testSetups {
		t.Run(testSetup.name, func(t *testing.T) {
			cos := Cosine(&testSetup.a, &testSetup.b)

			if !PracticallyEquals(cos, testSetup.expectedCosine, 0.000001) {
				t.Errorf("Cosine expected to be %f but was %f for test \"%s\".", testSetup.expectedCosine, cos, testSetup.name)
			}
		})
	}
}

func TestSinus(t *testing.T) {
	radFor45deg := float32(math.Pi / 4.0)
	testSetups := []struct {
		a, b         T
		expectedSine float32
		name         string
	}{
		{a: T{1, 0, 0}, b: T{1, 0, 0}, expectedSine: math.Sin(0 * radFor45deg), name: "0/360 degree angle, equal/parallell vectors"},
		{a: T{1, 0, 0}, b: T{1, 1, 0}, expectedSine: math.Sin(1 * radFor45deg), name: "45 degree angle"},
		{a: T{1, 0, 0}, b: T{0, 1, 0}, expectedSine: math.Sin(2 * radFor45deg), name: "90 degree angle, orthogonal vectors"},
		{a: T{1, 0, 0}, b: T{-1, 1, 0}, expectedSine: math.Sin(3 * radFor45deg), name: "135 degree angle"},
		{a: T{1, 0, 0}, b: T{-1, 0, 0}, expectedSine: math.Sin(4 * radFor45deg), name: "180 degree angle, inverted/anti parallell vectors"},
		{a: T{1, 0, 0}, b: T{-1, -1, 0}, expectedSine: math.Abs(math.Sin(5 * radFor45deg)), name: "225 degree angle"},
		{a: T{1, 0, 0}, b: T{0, -1, 0}, expectedSine: math.Abs(math.Sin(6 * radFor45deg)), name: "270 degree angle, orthogonal vectors"},
		{a: T{1, 0, 0}, b: T{1, -1, 0}, expectedSine: math.Abs(math.Sin(7 * radFor45deg)), name: "315 degree angle"},
	}

	for _, testSetup := range testSetups {
		t.Run(testSetup.name, func(t *testing.T) {
			sin := Sinus(&testSetup.a, &testSetup.b)

			if !PracticallyEquals(sin, testSetup.expectedSine, 0.000001) {
				t.Errorf("Sine expected to be %f but was %f for test \"%s\".", testSetup.expectedSine, sin, testSetup.name)
			}
		})
	}
}

func TestIsZeroEps(t *testing.T) {
	tests := []struct {
		name    string
		vec     T
		epsilon float32
		want    bool
	}{
		{"exact zero", T{0, 0, 0}, 0.0001, true},
		{"within epsilon", T{0.00001, -0.00001, 0.00001}, 0.0001, true},
		{"at epsilon boundary", T{0.0001, 0.0001, 0.0001}, 0.0001, true},
		{"outside epsilon", T{0.001, 0, 0}, 0.0001, false},
		{"one component outside", T{0.00001, 0.001, 0.00001}, 0.0001, false},
		{"negative outside epsilon", T{-0.001, 0, 0}, 0.0001, false},
		{"large values", T{1.0, 2.0, 3.0}, 0.0001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vec.IsZeroEps(tt.epsilon); got != tt.want {
				t.Errorf("IsZeroEps() = %v, want %v for vec %v with epsilon %v", got, tt.want, tt.vec, tt.epsilon)
			}
		})
	}
}
