package go3d

import "math"


const (
	InversePi = 1 / math.Pi
	HalfPi = math.Pi / 2
	QuarterPi = math.Pi / 4
	TwoPi = math.Pi * 2
	FourPi = math.Pi * 4
	DefaultEpsilon = 0.0001
)

var (
	Epsilon float64 = DefaultEpsilon
	SquareEpsilon float64 = DefaultEpsilon * DefaultEpsilon
)