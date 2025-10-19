package vec4

import (
	"testing"
)

func TestIsZeroEps(t *testing.T) {
	tests := []struct {
		name    string
		vec     T
		epsilon float32
		want    bool
	}{
		{"exact zero", T{0, 0, 0, 0}, 0.0001, true},
		{"within epsilon", T{0.00001, -0.00001, 0.00001, -0.00001}, 0.0001, true},
		{"at epsilon boundary", T{0.0001, 0.0001, 0.0001, 0.0001}, 0.0001, true},
		{"outside epsilon", T{0.001, 0, 0, 0}, 0.0001, false},
		{"one component outside", T{0.00001, 0.001, 0.00001, 0.00001}, 0.0001, false},
		{"negative outside epsilon", T{-0.001, 0, 0, 0}, 0.0001, false},
		{"large values", T{1.0, 2.0, 3.0, 4.0}, 0.0001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vec.IsZeroEps(tt.epsilon); got != tt.want {
				t.Errorf("IsZeroEps() = %v, want %v for vec %v with epsilon %v", got, tt.want, tt.vec, tt.epsilon)
			}
		})
	}
}
