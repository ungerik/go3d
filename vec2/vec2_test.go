package vec2

import (
	"math"
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
	expectedV1n := T{float32(1 / v1Length), float32(-1 / v1Length)}

	v1n := v1.Normal()

	if v1n != expectedV1n {
		t.Fail()
	}
}

func TestNormal2(t *testing.T) {
	v1 := T{4, 6}

	v1Length := math.Sqrt(4*4 + 6*6)
	expectedV1n := T{float32(6.0 / v1Length), float32(-4 / v1Length)}

	v1n := v1.Normal()

	if v1n != expectedV1n {
		t.Fail()
	}
}

func TestNormalCCW(t *testing.T) {
	v1 := T{1.0, 1.0}

	v1Length := math.Sqrt(1*1 + 1*1)
	expectedV1n := T{float32(-1 / v1Length), float32(1 / v1Length)}

	v1n := v1.NormalCCW()

	if v1n != expectedV1n {
		t.Fail()
	}
}

func TestNormalCCW2(t *testing.T) {
	v1 := T{4, 6}

	v1Length := math.Sqrt(4*4 + 6*6)
	expectedV1n := T{float32(-6.0 / v1Length), float32(4 / v1Length)}

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
