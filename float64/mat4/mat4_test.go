package mat4

import (
	"math"
	"testing"

	"github.com/ungerik/go3d/float64/vec3"
	"github.com/ungerik/go3d/float64/vec4"
)

const EPSILON = 0.0000001

func TestIsZeroEps(t *testing.T) {
	tests := []struct {
		name    string
		mat     T
		epsilon float64
		want    bool
	}{
		{"exact zero", Zero, 0.0001, true},
		{"within epsilon", T{vec4.T{0.00001, -0.00001, 0.00001, -0.00001}, vec4.T{-0.00001, 0.00001, -0.00001, 0.00001}, vec4.T{0.00001, -0.00001, 0.00001, -0.00001}, vec4.T{-0.00001, 0.00001, -0.00001, 0.00001}}, 0.0001, true},
		{"at epsilon boundary", T{vec4.T{0.0001, 0.0001, 0.0001, 0.0001}, vec4.T{0.0001, 0.0001, 0.0001, 0.0001}, vec4.T{0.0001, 0.0001, 0.0001, 0.0001}, vec4.T{0.0001, 0.0001, 0.0001, 0.0001}}, 0.0001, true},
		{"outside epsilon", T{vec4.T{0.001, 0, 0, 0}, vec4.T{0, 0, 0, 0}, vec4.T{0, 0, 0, 0}, vec4.T{0, 0, 0, 0}}, 0.0001, false},
		{"one element outside", T{vec4.T{0.00001, 0.00001, 0.00001, 0.00001}, vec4.T{0.001, 0.00001, 0.00001, 0.00001}, vec4.T{0.00001, 0.00001, 0.00001, 0.00001}, vec4.T{0.00001, 0.00001, 0.00001, 0.00001}}, 0.0001, false},
		{"negative outside epsilon", T{vec4.T{-0.001, 0, 0, 0}, vec4.T{0, 0, 0, 0}, vec4.T{0, 0, 0, 0}, vec4.T{0, 0, 0, 0}}, 0.0001, false},
		{"identity matrix", Ident, 0.0001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mat.IsZeroEps(tt.epsilon); got != tt.want {
				t.Errorf("IsZeroEps() = %v, want %v for mat with epsilon %v", got, tt.want, tt.epsilon)
			}
		})
	}
}

func TestScale(t *testing.T) {
	m := Ident
	m.Scale(2.5)

	expected := Ident
	expected[0][0] = 2.5
	expected[1][1] = 2.5
	expected[2][2] = 2.5

	if m != expected {
		t.Errorf("Scale() failed: got %v, want %v", m, expected)
	}
}

func TestScaled(t *testing.T) {
	m := Ident
	result := m.Scaled(2.5)

	// Original should be unchanged
	if m != Ident {
		t.Errorf("Scaled() modified original matrix")
	}

	expected := Ident
	expected[0][0] = 2.5
	expected[1][1] = 2.5
	expected[2][2] = 2.5

	if result != expected {
		t.Errorf("Scaled() failed: got %v, want %v", result, expected)
	}
}

func TestTrace(t *testing.T) {
	m := Ident
	if trace := m.Trace(); trace != 4.0 {
		t.Errorf("Trace() of identity should be 4.0, got %f", trace)
	}

	m[0][0] = 2
	m[1][1] = 3
	m[2][2] = 5
	m[3][3] = 7
	if trace := m.Trace(); trace != 17.0 {
		t.Errorf("Trace() should be 17.0, got %f", trace)
	}
}

func TestTrace3(t *testing.T) {
	m := Ident
	if trace := m.Trace3(); trace != 3.0 {
		t.Errorf("Trace3() of identity should be 3.0, got %f", trace)
	}

	m[0][0] = 2
	m[1][1] = 3
	m[2][2] = 5
	m[3][3] = 7 // Should not be included in Trace3
	if trace := m.Trace3(); trace != 10.0 {
		t.Errorf("Trace3() should be 10.0, got %f", trace)
	}
}

func TestAssignMul(t *testing.T) {
	m1, _ := Parse("1 0 0 0 0 1 0 0 0 0 1 0 3 5 7 1")
	m2, _ := Parse("1 0 0 0 0 1 0 0 0 0 1 0 2 4 8 1")
	var result T
	result.AssignMul(&m1, &m2)
	expected, _ := Parse("1 0 0 0 0 1 0 0 0 0 1 0 5 9 15 1")

	if result != expected {
		t.Errorf("AssignMul() incorrect: got %v, want %v", result, expected)
	}
}

func TestMulVec4vsTransformVec4(t *testing.T) {
	m, _ := Parse("2 0 0 0 0 3 0 0 0 0 4 0 1 2 3 1")
	v := vec4.T{1, 2, 3, 1}

	// Method 1: MulVec4
	result1 := m.MulVec4(&v)

	// Method 2: TransformVec4
	result2 := v
	m.TransformVec4(&result2)

	if math.Abs(result1[0]-result2[0]) > EPSILON ||
		math.Abs(result1[1]-result2[1]) > EPSILON ||
		math.Abs(result1[2]-result2[2]) > EPSILON ||
		math.Abs(result1[3]-result2[3]) > EPSILON {
		t.Errorf("MulVec4 and TransformVec4 differ: %v vs %v", result1, result2)
	}
}

func TestMulVec3vsTransformVec3(t *testing.T) {
	m, _ := Parse("2 0 0 0 0 3 0 0 0 0 4 0 1 2 3 1")
	v := vec3.T{1, 2, 3}

	// Method 1: MulVec3
	result1 := m.MulVec3(&v)

	// Method 2: TransformVec3
	result2 := v
	m.TransformVec3(&result2)

	if math.Abs(result1[0]-result2[0]) > EPSILON ||
		math.Abs(result1[1]-result2[1]) > EPSILON ||
		math.Abs(result1[2]-result2[2]) > EPSILON {
		t.Errorf("MulVec3 and TransformVec3 differ: %v vs %v", result1, result2)
	}
}

func TestTranspose(t *testing.T) {
	m, _ := Parse("1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16")
	original := m
	m.Transpose()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if m[i][j] != original[j][i] {
				t.Errorf("Transpose failed at [%d][%d]: got %f, want %f", i, j, m[i][j], original[j][i])
			}
		}
	}
}

func TestParseAndString(t *testing.T) {
	original := "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16"
	m, err := Parse(original)
	if err != nil {
		t.Errorf("Parse() failed: %v", err)
	}

	// Parse and stringify should preserve values
	m2, err := Parse(m.String())
	if err != nil {
		t.Errorf("Parse() of String() failed: %v", err)
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if math.Abs(m[i][j]-m2[i][j]) > EPSILON {
				t.Errorf("Parse/String roundtrip failed at [%d][%d]", i, j)
			}
		}
	}
}

func TestIdentity(t *testing.T) {
	// Identity matrix should have 1s on diagonal, 0s elsewhere
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			expected := 0.0
			if i == j {
				expected = 1.0
			}
			if Ident[i][j] != expected {
				t.Errorf("Identity matrix incorrect at [%d][%d]: got %f, want %f", i, j, Ident[i][j], expected)
			}
		}
	}
}

func TestZero(t *testing.T) {
	// Zero matrix should have all 0s
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if Zero[i][j] != 0 {
				t.Errorf("Zero matrix incorrect at [%d][%d]: got %f, want 0", i, j, Zero[i][j])
			}
		}
	}
}
