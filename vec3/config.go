package vec3

// Epsilon is the tolerance used for numerical stability in floating-point comparisons.
// It is used by Normalize(), Normal(), and other methods to determine if a vector is
// effectively zero or already normalized. Can be adjusted for specific use cases.
// Default: 1e-8 for float32 precision.
var Epsilon float32 = 1e-8
