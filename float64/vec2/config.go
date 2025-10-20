package vec2

// Epsilon is the tolerance used for numerical stability in floating-point comparisons.
// It is used by Normalize() and other methods to determine if a vector is effectively
// zero or already normalized. Can be adjusted for specific use cases.
// Default: 1e-14 for float64 precision.
var Epsilon float64 = 1e-14
