package go3d

import "math"

/*
http://graphics.stanford.edu/~seander/bithacks.html
*/

func OddInt(value int64) bool {
	return (value & 1) != 0
}


func EvenInt(value int64) bool {
	return (value & 1) == 0
}


func Odd(value float64) bool {
	return OddInt(int64(value))
}


func Even(value float64) bool {
	return EvenInt(int64(value))
}


func LowerEvenInt(value int64) int64 {
	return value & (^1)
}


func UpperEvenInt(value int64) int64 {
	if OddInt(value) {
		return value + 1
	}
	return value
}


func LowerEven(value float64) float64 {
	return math.Floor(value*0.5) * 2.0
}


func UpperEven(value float64) float64 {
	return math.Ceil(value*0.5) * 2.0
}


func Round(value float64) float64 {
	return math.Floor(value + 0.5)
}


func RoundInt(value float64) int64 {
	return int64(value + 0.5)
}


func IsPowerOfTwo(value uint64) bool {
	return (value & (value - 1)) == 0
}


func LowerPowerOfTwo(value uint64) uint64 {
	for i := uint(63); i >= 0; i-- {
		mask := uint64(1) << i
		if (value & mask) != 0 {
			return mask
		}
	}
	return 0
}


func UpperPowerOfTwo(value uint64) uint64 {
	value--
	value |= value >> 1
	value |= value >> 2
	value |= value >> 4
	value |= value >> 8
	value |= value >> 16
	value |= value >> 32
	value++
	return value
}


func ClampToInt8(value int64) int8 {
	if value < math.MinInt8 {
		return math.MinInt8
	}
	if value > math.MaxInt8 {
		return math.MaxInt8
	}
	return int8(value)
}


func ClampToUint8(value int64) uint8 {
	if value < 0 {
		return 0
	}
	if value > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(value)
}


func ClampToInt16(value int64) int16 {
	if value < math.MinInt16 {
		return math.MinInt16
	}
	if value > math.MaxInt16 {
		return math.MaxInt16
	}
	return int16(value)
}


func ClampToUint16(value int64) uint16 {
	if value < 0 {
		return 0
	}
	if value > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(value)
}


func ClampInt(value int64, min int64, max int64) int64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}


func Clamp(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}


func Frac(value float64) float64 {
	return value - math.Floor(value)
}


func Sign(value float64) float64 {
	return math.Float64frombits(math.Float64bits(value) >> 63)
}


func SignInt(value int64) int64 {
	return value >> 63
}


func IsSpecialFloat(value float64) bool {
	return math.IsInf(value, 0) || math.IsNaN(value)
}


func InRange(value float64, min float64, max float64) bool {
	return value >= min && value <= max
}


func Equal(a float64, b float64, epsilon float64) bool {
	return a >= b-epsilon && a <= b+epsilon
}


func EqualEpsilon(a float64, b float64) bool {
	return Equal(a, b, Epsilon)
}


func EqualSquareEpsilon(a float64, b float64) bool {
	return Equal(a, b, SquareEpsilon)
}


func InvSqrt(value float64) float64 {
	// TODO: Implement http://en.wikipedia.org/wiki/Fast_inverse_square_root
	return 1 / math.Sqrt(value)
}  