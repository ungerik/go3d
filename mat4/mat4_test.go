package mat4

import (
	math "github.com/barnex/fmath"
	"github.com/ungerik/go3d/mat3"
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec4"
	"testing"
)

const EPSILON = 0.0001

// Some matrices used in multiple tests.
var (
	TEST_MATRIX1 = T{vec4.T{0.38016528, -0.0661157, -0.008264462, -0},
		vec4.T{-0.19834709, 0.33884296, -0.08264463, 0},
		vec4.T{0.11570247, -0.28099173, 0.21487603, -0},
		vec4.T{18.958677, -33.471073, 8.066115, 0.99999994}}

	TEST_MATRIX2 = T{vec4.T{23, -4, -0.5, -0},
		vec4.T{-12, 20.5, -5, 0},
		vec4.T{7, -17, 13, -0},
		vec4.T{1147, -2025, 488, 60.5}}

	ROW_123_CHANGED, _ = Parse("3 1 0.5 0 2 5 2 0 1 6 7 0 2 100 1 1")
)

func TestDeterminant(t *testing.T) {
	detId := Ident.Determinant()
	if detId != 1 {
		t.Errorf("Wrong determinant for identity matrix: %f", detId)
	}

	detTwo := Ident
	detTwo[0][0] = 2
	if det := detTwo.Determinant(); det != 2 {
		t.Errorf("Wrong determinant: %f", det)
	}

	scale2 := Ident.Scale(2)
	if det := scale2.Determinant(); det != 2*2*2*1 {
		t.Errorf("Wrong determinant: %f", det)
	}

	row1changed, _ := Parse("3 0 0 0 2 2 0 0 1 0 2 0 2 0 0 1")
	if det := row1changed.Determinant(); det != 12 {
		t.Errorf("Wrong determinant: %f", det)
	}

	row12changed, _ := Parse("3 1 0 0 2 5 0 0 1 6 2 0 2 100 0 1")
	if det := row12changed.Determinant(); det != 26 {
		t.Errorf("Wrong determinant: %f", det)
	}

	row123changed := ROW_123_CHANGED
	if det := row123changed.Determinant3x3(); det != 60.500 {
		t.Errorf("Wrong determinant for 3x3 matrix: %f", det)
	}
	if det := row123changed.Determinant(); det != 60.500 {
		t.Errorf("Wrong determinant: %f", det)
	}
	randomMatrix, err := Parse("0.43685 0.81673 0.63721 0.23421 0.16600 0.40608 0.53479 0.43210 0.37328 0.36436 0.56356 0.66830 0.32475 0.14294 0.42137 0.98046")
	randomMatrix.Transpose() //transpose for easy comparability with octave output
	if err != nil {
		t.Errorf("Could not parse random matrix: %v", err)
	}
	if det := randomMatrix.Determinant3x3(); math.Abs(det-0.043437) > EPSILON {
		t.Errorf("Wrong determinant for random sub 3x3 matrix: %f", det)
	}

	if det := randomMatrix.Determinant(); math.Abs(det-0.012208) > EPSILON {
		t.Errorf("Wrong determinant for random matrix: %f", det)
	}
}

func TestMaskedBlock(t *testing.T) {
	m := ROW_123_CHANGED
	blocked_expected := mat3.T{vec3.T{5, 2, 0}, vec3.T{6, 7, 0}, vec3.T{100, 1, 1}}
	if blocked := m.maskedBlock(0, 0); *blocked != blocked_expected {
		t.Errorf("Did not block 0,0 correctly: %#v", blocked)
	}
}

func TestAdjugate(t *testing.T) {
	adj := ROW_123_CHANGED
	adj.Adjugate()
	// Computed in octave:
	adj_expected := T{vec4.T{23, -4, -0.5, -0}, vec4.T{-12, 20.5, -5, 0}, vec4.T{7, -17, 13, -0}, vec4.T{1147, -2025, 488, 60.5}}
	if adj != adj_expected {
		t.Errorf("Adjugate not computed correctly: %#v", adj)
	}
}

func TestInvert(t *testing.T) {
	inv := ROW_123_CHANGED
	inv.Invert()
	// Computed in octave:
	inv_expected := T{vec4.T{0.38016528, -0.0661157, -0.008264462, -0}, vec4.T{-0.19834709, 0.33884296, -0.08264463, 0}, vec4.T{0.11570247, -0.28099173, 0.21487603, -0}, vec4.T{18.958677, -33.471073, 8.066115, 0.99999994}}
	if inv != inv_expected {
		t.Errorf("Inverse not computed correctly: %#v", inv)
	}
}

func TestMultSimpleMatrices(t *testing.T) {
	m1 := T{vec4.T{1, 0, 0, 2},
		vec4.T{0, 1, 2, 0},
		vec4.T{0, 2, 1, 0},
		vec4.T{2, 0, 0, 1}}
	m2 := m1
	var mMult T
	mMult.AssignMul(&m1, &m2)
	t.Log(&m1)
	t.Log(&m2)
	m1.MultMatrix(&m2)
	if m1 != mMult {
		t.Errorf("Multiplication of matrices above failed, expected: \n%v \ngotten: \n%v", &mMult, &m1)
	}
}

func TestMultMatrixVsAssignMul(t *testing.T) {
	m1 := TEST_MATRIX1
	m2 := TEST_MATRIX2
	var mMult T
	mMult.AssignMul(&m1, &m2)
	t.Log(&m1)
	t.Log(&m2)
	m1.MultMatrix(&m2)
	if m1 != mMult {
		t.Errorf("Multiplication of matrices above failed, expected: \n%v \ngotten: \n%v", &mMult, &m1)
	}
}

func BenchmarkAssignMul(b *testing.B) {
	m1 := TEST_MATRIX1
	m2 := TEST_MATRIX2
	var mMult T
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mMult.AssignMul(&m1, &m2)
	}
}

func BenchmarkMultMatrix(b *testing.B) {
	m1 := TEST_MATRIX1
	m2 := TEST_MATRIX2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m1.MultMatrix(&m2)
	}
}

func TestMulVec4vsTransformVec4(t *testing.T) {
	m1 := TEST_MATRIX1
	v := vec4.T{1, 1.5, 2, 2.5}
	v_1 := m1.MulVec4(&v)
	v_2 := m1.MulVec4(&v_1)

	m1.TransformVec4(&v)
	m1.TransformVec4(&v)

	if v_2 != v {
		t.Error(v_2, v)
	}

}

func BenchmarkMulVec4(b *testing.B) {
	m1 := TEST_MATRIX1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := vec4.T{1, 1.5, 2, 2.5}
		v_1 := m1.MulVec4(&v)
		m1.MulVec4(&v_1)
	}
}

func BenchmarkTransformVec4(b *testing.B) {
	m1 := TEST_MATRIX1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := vec4.T{1, 1.5, 2, 2.5}
		m1.TransformVec4(&v)
		m1.TransformVec4(&v)
	}
}
