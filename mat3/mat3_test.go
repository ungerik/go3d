package mat3

import (
	"math"
	"testing"

	"github.com/ungerik/go3d/mat2"
	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/vec3"
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

func Test_ColsAndRows(t *testing.T) {
	A := testMatrix2

	a11 := A.Get(0, 0)
	a21 := A.Get(0, 1)
	a31 := A.Get(0, 2)

	a12 := A.Get(1, 0)
	a22 := A.Get(1, 1)
	a32 := A.Get(1, 2)

	a13 := A.Get(2, 0)
	a23 := A.Get(2, 1)
	a33 := A.Get(2, 2)

	correctReference := a11 == 23 && a21 == -4 && a31 == -0.5 &&
		a12 == -12 && a22 == 20.5 && a32 == -5 &&
		a13 == 7 && a23 == -17 && a33 == 13

	if !correctReference {
		t.Errorf("matrix ill referenced")
	}
}

func TestT_Transposed(t *testing.T) {
	matrix := T{
		vec3.T{1, 2, 3},
		vec3.T{4, 5, 6},
		vec3.T{7, 8, 9},
	}
	expectedMatrix := T{
		vec3.T{1, 4, 7},
		vec3.T{2, 5, 8},
		vec3.T{3, 6, 9},
	}

	transposedMatrix := matrix.Transposed()

	if transposedMatrix != expectedMatrix {
		t.Errorf("matrix transposed wrong: %v --> %v", matrix, transposedMatrix)
	}
}

func TestT_Transpose(t *testing.T) {
	matrix := T{
		vec3.T{10, 20, 30},
		vec3.T{40, 50, 60},
		vec3.T{70, 80, 90},
	}

	expectedMatrix := T{
		vec3.T{10, 40, 70},
		vec3.T{20, 50, 80},
		vec3.T{30, 60, 90},
	}

	transposedMatrix := matrix
	transposedMatrix.Transpose()

	if transposedMatrix != expectedMatrix {
		t.Errorf("matrix transposed wrong: %v --> %v", matrix, transposedMatrix)
	}
}

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
	if det := randomMatrix.Determinant(); PracticallyEquals(det, 0.043437) {
		t.Errorf("Wrong determinant for random sub 3x3 matrix: %f", det)
	}
}

func PracticallyEquals(value1 float32, value2 float32) bool {
	return math.Abs(float64(value1-value2)) > EPSILON
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

	// Computed in octave:
	adjExpected := T{
		vec3.T{23, -4, -0.5},
		vec3.T{-12, 20.5, -5},
		vec3.T{7, -17, 13},
	}

	adj.Adjugate()

	if adj != adjExpected {
		t.Errorf("Adjugate not computed correctly: %#v", adj)
	}
}

func TestAdjugated(t *testing.T) {
	sqrt2 := float32(math.Sqrt(2))
	A := T{
		vec3.T{1, 0, -1},
		vec3.T{0, sqrt2, 0},
		vec3.T{1, 0, 1},
	}

	expectedAdjugated := T{
		vec3.T{1.4142135623730951, -0, 1.4142135623730951},
		vec3.T{-0, 2, -0},
		vec3.T{-1.4142135623730951, -0, 1.4142135623730951},
	}

	adjA := A.Adjugated()

	if adjA != expectedAdjugated {
		t.Errorf("Adjugate not computed correctly: %v", adjA)
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

func TestInvert_ok2(t *testing.T) {
	sqrt2 := float32(math.Sqrt(2))
	A := T{
		vec3.T{1, 0, -1},
		vec3.T{0, sqrt2, 0},
		vec3.T{1, 0, 1},
	}

	expectedInverted := T{
		vec3.T{0.5, 0, 0.5},
		vec3.T{0, 0.7071067811865475, 0},
		vec3.T{-0.5, 0, 0.5},
	}

	invA, err := A.Inverted()
	if err != nil {
		t.Error("Inverse not computed correctly", err)
	}

	if !invA.PracticallyEquals(&expectedInverted, EPSILON) {
		t.Errorf("Inverse not computed correctly: %v", A)
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
