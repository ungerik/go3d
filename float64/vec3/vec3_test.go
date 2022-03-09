package vec3

import (
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
