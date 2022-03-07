package mat2

import (
	"math"
	"testing"

	"github.com/ungerik/go3d/vec2"
)

const EPSILON = 0.0001

// Some matrices used in multiple tests.
var (
	invertableMatrix1    = T{vec2.T{4, -2}, vec2.T{8, -3}}
	invertedMatrix1      = T{vec2.T{-3.0 / 4.0, 1.0 / 2.0}, vec2.T{-2, 1}}
	nonInvertableMatrix1 = T{vec2.T{1, 1}, vec2.T{1, 1}}
	nonInvertableMatrix2 = T{vec2.T{2, 0}, vec2.T{1, 0}}

	testMatrix1 = T{
		vec2.T{0.38016528, -0.0661157},
		vec2.T{-0.19834709, 0.33884296},
	}

	testMatrix2 = T{
		vec2.T{23, -4},
		vec2.T{-12, 20.5},
	}

	row123Changed, _ = Parse("3 1   2 5")
)

func TestT_Transposed(t *testing.T) {
	matrix := T{
		vec2.T{1, 2},
		vec2.T{3, 4},
	}
	expectedMatrix := T{
		vec2.T{1, 3},
		vec2.T{2, 4},
	}

	transposedMatrix := matrix.Transposed()

	if transposedMatrix != expectedMatrix {
		t.Errorf("matrix trasnposed wrong: %v --> %v", matrix, transposedMatrix)
	}
}

func TestT_Transpose(t *testing.T) {
	matrix := T{
		vec2.T{10, 20},
		vec2.T{30, 40},
	}

	expectedMatrix := T{
		vec2.T{10, 30},
		vec2.T{20, 40},
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
	if det := detTwo.Determinant(); det != (2*1 - 0*0) {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_3(t *testing.T) {
	scale2 := Ident.Scaled(2)
	if det := scale2.Determinant(); det != (2*2 - 0*0) {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_4(t *testing.T) {
	row1changed, _ := Parse("3 0   2 2")
	if det := row1changed.Determinant(); det != (3*2 - 0*2) {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_5(t *testing.T) {
	row12changed, _ := Parse("3 1   2 5")
	if det := row12changed.Determinant(); det != (3*5 - 1*2) {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_7(t *testing.T) {
	randomMatrix, err := Parse("0.43685 0.81673   0.16600 0.40608")
	randomMatrix.Transpose()
	if err != nil {
		t.Errorf("Could not parse random matrix: %v", err)
	}
	if det := randomMatrix.Determinant(); PracticallyEquals(det, 0.0418189) {
		t.Errorf("Wrong determinant for random sub 3x3 matrix: %f", det)
	}
}

func PracticallyEquals(value1 float32, value2 float32) bool {
	return math.Abs(float64(value1-value2)) > EPSILON
}

func TestDeterminant_6(t *testing.T) {
	row123changed := row123Changed
	if det := row123changed.Determinant(); det != (3*5 - 2*1) {
		t.Errorf("Wrong determinant for 3x3 matrix: %f", det)
	}
}

func TestDeterminant_1(t *testing.T) {
	detId := Ident.Determinant()
	if detId != 1 {
		t.Errorf("Wrong determinant for identity matrix: %f", detId)
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
