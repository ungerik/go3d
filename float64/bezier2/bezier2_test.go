package bezier2

import (
	"testing"

	"github.com/ungerik/go3d/float64/vec2"
)

func TestPoint(t *testing.T) {
	b := T{vec2.T{0, 0}, vec2.T{1, 1}, vec2.T{2, 1}, vec2.T{3, 0}}
	if got, want := b.Point(0), (vec2.T{0, 0}); got != want {
		t.Errorf("cubic bezier point at t=0 failed, got %v, want %v", got, want)
	}
	if got, want := b.Point(1), (vec2.T{3, 0}); got != want {
		t.Errorf("cubic bezier point at t=1 failed, got %v, want %v", got, want)
	}
	if got, want := b.Point(0.5), (vec2.T{1.5, 0.75}); got != want {
		t.Errorf("cubic bezier point at t=0.5 failed, got %v, want %v", got, want)
	}
	if got, want := b.Point(0.25), (vec2.T{0.75, 0.5625}); got != want {
		t.Errorf("cubic bezier point at t=0.25 failed, got %v, want %v", got, want)
	}
	if got, want := b.Point(0.75), (vec2.T{2.25, 0.5625}); got != want {
		t.Errorf("cubic bezier point at t=0.75 failed, got %v, want %v", got, want)
	}
}

func TestTangent(t *testing.T) {
	b := T{vec2.T{0, 0}, vec2.T{1, 1}, vec2.T{2, 1}, vec2.T{3, 0}}
	if got, want := b.Tangent(0), (vec2.T{3, 3}); got != want {
		t.Errorf("cubic bezier tangent at t=0 failed, got %v, want %v", got, want)
	}
	if got, want := b.Tangent(1), (vec2.T{3, -3}); got != want {
		t.Errorf("cubic bezier tangent at t=1 failed, got %v, want %v", got, want)
	}
	if got, want := b.Tangent(0.5), (vec2.T{3, 0}); got != want {
		t.Errorf("cubic bezier tangent at t=0.5 failed, got %v, want %v", got, want)
	}
	if got, want := b.Tangent(0.25), (vec2.T{3, 1.5}); got != want {
		t.Errorf("cubic bezier tangent at t=0.25 failed, got %v, want %v", got, want)
	}
	if got, want := b.Tangent(0.75), (vec2.T{3, -1.5}); got != want {
		t.Errorf("cubic bezier tangent at t=0.75 failed, got %v, want %v", got, want)
	}
}
