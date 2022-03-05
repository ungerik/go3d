package mat3

import (
	"math"
	"testing"

	"github.com/ungerik/go3d/float64/mat2"
	"github.com/ungerik/go3d/float64/vec2"
	"github.com/ungerik/go3d/float64/vec3"
)

const EPSILON = 0.0001

// Some matrices used in multiple tests.
var (
	invertableMatrix1    = T{vec3.T{4, -2, 3}, vec3.T{8, -3, 5}, vec3.T{7, -2, 4}}
	invertedMatrix1      = T{vec3.T{-2, 2, -1}, vec3.T{3, -5, 4}, vec3.T{5, -6, 4}}
	nonInvertableMatrix1 = T{vec3.T{1, 1, 1}, vec3.T{1, 1, 1}, vec3.T{1, 1, 1}}
	nonInvertableMatrix2 = T{vec3.T{1, 1, 1}, vec3.T{1, 0, 1}, vec3.T{1, 1, 1}}

	testMatrix1 = T{
		vec3.T{0.38016528, -0.0661157, -0.008264462},
		vec3.T{-0.19834709, 0.33884296, -0.08264463},
		vec3.T{0.11570247, -0.28099173, 0.21487603},
	}

	testMatrix2 = T{
		vec3.T{23, -4, -0.5},
		vec3.T{-12, 20.5, -5},
		vec3.T{7, -17, 13},
	}

	row123Changed, _ = Parse("3 1 0.5   2 5 2   1 6 7")
)

func TestDeterminant_2(t *testing.T) {
	detTwo := Ident
	detTwo[0][0] = 2
	if det := detTwo.Determinant(); det != 2 {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_3(t *testing.T) {
	scale2 := Ident.Scaled(2)
	if det := scale2.Determinant(); det != 2*2*2*1 {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_4(t *testing.T) {
	row1changed, _ := Parse("3 0 0   2 2 0   1 0 2")
	if det := row1changed.Determinant(); det != 12 {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_5(t *testing.T) {
	row12changed, _ := Parse("3 1 0   2 5 0   1 6 2")
	if det := row12changed.Determinant(); det != 26 {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_7(t *testing.T) {
	randomMatrix, err := Parse("0.43685 0.81673 0.63721   0.16600 0.40608 0.53479   0.37328 0.36436 0.56356")
	randomMatrix.Transpose()
	if err != nil {
		t.Errorf("Could not parse random matrix: %v", err)
	}
	if det := randomMatrix.Determinant(); practicallyEqual(det, 0.043437) {
		t.Errorf("Wrong determinant for random sub 3x3 matrix: %f", det)
	}
}

func practicallyEqual(value1 float64, value2 float64) bool {
	return math.Abs(value1-value2) > EPSILON
}

func TestDeterminant_6(t *testing.T) {
	row123changed := row123Changed
	if det := row123changed.Determinant(); det != 60.500 {
		t.Errorf("Wrong determinant for 3x3 matrix: %f", det)
	}
}

func TestDeterminant_1(t *testing.T) {
	detId := Ident.Determinant()
	if detId != 1 {
		t.Errorf("Wrong determinant for identity matrix: %f", detId)
	}
}

func TestMaskedBlock(t *testing.T) {
	m := row123Changed
	blockedExpected := mat2.T{vec2.T{5, 2}, vec2.T{6, 7}}
	if blocked := m.maskedBlock(0, 0); *blocked != blockedExpected {
		t.Errorf("Did not block 0,0 correctly: %#v", blocked)
	}
}

func TestAdjugate(t *testing.T) {
	adj := row123Changed
	adj.Adjugate()
	// Computed in octave:
	adjExpected := T{vec3.T{23, -4, -0.5}, vec3.T{-12, 20.5, -5}, vec3.T{7, -17, 13}}
	if adj != adjExpected {
		t.Errorf("Adjugate not computed correctly: %#v", adj)
	}
}

func TestInvert_ok(t *testing.T) {
	inv := invertableMatrix1
	_, err := inv.Invert()

	if err != nil {
		t.Error("Inverse not computed correctly", err)
	}

	invExpected := invertedMatrix1
	if inv != invExpected {
		t.Errorf("Inverse not computed correctly: %#v", inv)
	}
}

func TestInvert_nok_1(t *testing.T) {
	inv := nonInvertableMatrix1
	_, err := inv.Inverted()
	if err == nil {
		t.Error("Inverse should not be possible", err)
	}
}

func TestInvert_nok_2(t *testing.T) {
	inv := nonInvertableMatrix2
	_, err := inv.Inverted()
	if err == nil {
		t.Error("Inverse should not be possible", err)
	}
}

func BenchmarkAssignMul(b *testing.B) {
	m1 := testMatrix1
	m2 := testMatrix2
	var mMult T
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mMult.AssignMul(&m1, &m2)
	}
}
