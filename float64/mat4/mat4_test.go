package mat4

import (
	"testing"

	"github.com/ungerik/go3d/float64/vec4"
)

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
