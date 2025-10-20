package hermit2

import (
	"math"
	"testing"

	"github.com/ungerik/go3d/vec2"
)

const EPSILON = 0.0001

func TestParseAndString(t *testing.T) {
	original := T{
		A: PointTangent{
			Point:   vec2.T{1.0, 2.0},
			Tangent: vec2.T{0.5, 0.5},
		},
		B: PointTangent{
			Point:   vec2.T{3.0, 4.0},
			Tangent: vec2.T{-0.5, 0.5},
		},
	}

	// Convert to string and parse back
	str := original.String()
	parsed, err := Parse(str)
	if err != nil {
		t.Errorf("Parse() failed: %v", err)
	}

	// Check all components
	if math.Abs(float64(parsed.A.Point[0]-original.A.Point[0])) > EPSILON ||
		math.Abs(float64(parsed.A.Point[1]-original.A.Point[1])) > EPSILON ||
		math.Abs(float64(parsed.A.Tangent[0]-original.A.Tangent[0])) > EPSILON ||
		math.Abs(float64(parsed.A.Tangent[1]-original.A.Tangent[1])) > EPSILON ||
		math.Abs(float64(parsed.B.Point[0]-original.B.Point[0])) > EPSILON ||
		math.Abs(float64(parsed.B.Point[1]-original.B.Point[1])) > EPSILON ||
		math.Abs(float64(parsed.B.Tangent[0]-original.B.Tangent[0])) > EPSILON ||
		math.Abs(float64(parsed.B.Tangent[1]-original.B.Tangent[1])) > EPSILON {
		t.Errorf("Parse/String roundtrip failed: got %v, want %v", parsed, original)
	}
}

func TestPointEndpoints(t *testing.T) {
	// Create a simple hermit spline from (0,0) to (1,1) with tangents pointing in the same direction
	herm := T{
		A: PointTangent{
			Point:   vec2.T{0.0, 0.0},
			Tangent: vec2.T{1.0, 1.0},
		},
		B: PointTangent{
			Point:   vec2.T{1.0, 1.0},
			Tangent: vec2.T{1.0, 1.0},
		},
	}

	// At t=0, should return pointA
	p0 := herm.Point(0.0)
	if math.Abs(float64(p0[0]-herm.A.Point[0])) > EPSILON ||
		math.Abs(float64(p0[1]-herm.A.Point[1])) > EPSILON {
		t.Errorf("Point(0) should equal A.Point: got %v, want %v", p0, herm.A.Point)
	}

	// At t=1, should return pointB
	p1 := herm.Point(1.0)
	if math.Abs(float64(p1[0]-herm.B.Point[0])) > EPSILON ||
		math.Abs(float64(p1[1]-herm.B.Point[1])) > EPSILON {
		t.Errorf("Point(1) should equal B.Point: got %v, want %v", p1, herm.B.Point)
	}
}

func TestPointMidpoint(t *testing.T) {
	// Horizontal line from (0,0) to (2,0) with horizontal tangents
	herm := T{
		A: PointTangent{
			Point:   vec2.T{0.0, 0.0},
			Tangent: vec2.T{1.0, 0.0},
		},
		B: PointTangent{
			Point:   vec2.T{2.0, 0.0},
			Tangent: vec2.T{1.0, 0.0},
		},
	}

	// At t=0.5, should be approximately at (1, 0) for a straight line
	p := herm.Point(0.5)

	// The y-coordinate should be exactly 0 for horizontal tangents
	if math.Abs(float64(p[1])) > EPSILON {
		t.Errorf("Point(0.5) y-coordinate should be 0: got %v", p[1])
	}

	// The x-coordinate should be approximately 1.0
	if math.Abs(float64(p[0]-1.0)) > EPSILON {
		t.Errorf("Point(0.5) x-coordinate should be ~1.0: got %v", p[0])
	}
}

func TestPointFunctionMatchesMethod(t *testing.T) {
	herm := T{
		A: PointTangent{
			Point:   vec2.T{0.0, 0.0},
			Tangent: vec2.T{1.0, 0.5},
		},
		B: PointTangent{
			Point:   vec2.T{2.0, 1.0},
			Tangent: vec2.T{1.0, -0.5},
		},
	}

	tests := []float32{0.0, 0.25, 0.5, 0.75, 1.0}
	for _, tVal := range tests {
		// Method call
		p1 := herm.Point(tVal)

		// Function call
		p2 := Point(&herm.A.Point, &herm.A.Tangent, &herm.B.Point, &herm.B.Tangent, tVal)

		if math.Abs(float64(p1[0]-p2[0])) > EPSILON ||
			math.Abs(float64(p1[1]-p2[1])) > EPSILON {
			t.Errorf("Point method and function differ at t=%f: method=%v, func=%v", tVal, p1, p2)
		}
	}
}

func TestTangentAtZero(t *testing.T) {
	herm := T{
		A: PointTangent{
			Point:   vec2.T{0.0, 0.0},
			Tangent: vec2.T{1.0, 0.5},
		},
		B: PointTangent{
			Point:   vec2.T{2.0, 1.0},
			Tangent: vec2.T{0.5, 1.0},
		},
	}

	// At t=0, the Tangent function returns tangentA (verified by examining the formula at t=0)
	tan0 := herm.Tangent(0.0)
	expectedTan0 := herm.A.Tangent
	if math.Abs(float64(tan0[0]-expectedTan0[0])) > EPSILON ||
		math.Abs(float64(tan0[1]-expectedTan0[1])) > EPSILON {
		t.Errorf("Tangent(0) should equal A.Tangent: got %v, want %v", tan0, expectedTan0)
	}
}

func TestTangentFunctionMatchesMethod(t *testing.T) {
	herm := T{
		A: PointTangent{
			Point:   vec2.T{0.0, 0.0},
			Tangent: vec2.T{1.0, 0.5},
		},
		B: PointTangent{
			Point:   vec2.T{2.0, 1.0},
			Tangent: vec2.T{0.5, 1.0},
		},
	}

	tests := []float32{0.0, 0.25, 0.5, 0.75, 1.0}
	for _, tVal := range tests {
		// Method call
		tan1 := herm.Tangent(tVal)

		// Function call
		tan2 := Tangent(&herm.A.Point, &herm.A.Tangent, &herm.B.Point, &herm.B.Tangent, tVal)

		if math.Abs(float64(tan1[0]-tan2[0])) > EPSILON ||
			math.Abs(float64(tan1[1]-tan2[1])) > EPSILON {
			t.Errorf("Tangent method and function differ at t=%f: method=%v, func=%v", tVal, tan1, tan2)
		}
	}
}

func TestLengthValues(t *testing.T) {
	herm := T{
		A: PointTangent{
			Point:   vec2.T{0.0, 0.0},
			Tangent: vec2.T{1.0, 0.0},
		},
		B: PointTangent{
			Point:   vec2.T{2.0, 0.0},
			Tangent: vec2.T{1.0, 0.0},
		},
	}

	// Test that Length returns a valid positive number
	length0 := herm.Length(0.0)
	if math.IsNaN(float64(length0)) || math.IsInf(float64(length0), 0) {
		t.Errorf("Length(0) should be a valid number: got %f", length0)
	}

	length1 := herm.Length(1.0)
	if math.IsNaN(float64(length1)) || math.IsInf(float64(length1), 0) {
		t.Errorf("Length(1) should be a valid number: got %f", length1)
	}
}

func TestLengthFunctionMatchesMethod(t *testing.T) {
	herm := T{
		A: PointTangent{
			Point:   vec2.T{0.0, 0.0},
			Tangent: vec2.T{1.0, 0.5},
		},
		B: PointTangent{
			Point:   vec2.T{2.0, 1.0},
			Tangent: vec2.T{0.5, 1.0},
		},
	}

	tests := []float32{0.0, 0.25, 0.5, 0.75, 1.0}
	for _, tVal := range tests {
		// Method call
		len1 := herm.Length(tVal)

		// Function call
		len2 := Length(&herm.A.Point, &herm.A.Tangent, &herm.B.Point, &herm.B.Tangent, tVal)

		if math.Abs(float64(len1-len2)) > EPSILON {
			t.Errorf("Length method and function differ at t=%f: method=%f, func=%f", tVal, len1, len2)
		}
	}
}
