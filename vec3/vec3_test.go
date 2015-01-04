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
