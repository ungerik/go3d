# go3d

Performance-oriented 3D math library for Go, optimized for graphics programming, game development, and scientific computing.

[![Go Reference](https://pkg.go.dev/badge/github.com/ungerik/go3d.svg)](https://pkg.go.dev/github.com/ungerik/go3d)
[![Go Report Card](https://goreportcard.com/badge/github.com/ungerik/go3d)](https://goreportcard.com/report/github.com/ungerik/go3d)

## Features

- **Zero-allocation value types**: Stack-allocated arrays for maximum performance
- **Dual precision**: Both float32 and float64 versions of all types
- **Method chaining**: In-place operations return `*T` for fluent API
- **Immutable operations**: Past-tense methods return copies without modifying originals
- **OpenGL compatible**: Column-major matrices, right-handed coordinates
- **Cache-friendly**: Optimized data structures for modern CPU caches
- **Well-tested**: Comprehensive test coverage with benchmarks

## Installation

```bash
go get github.com/ungerik/go3d
```

## Quick Start

```go
package main

import (
    "github.com/ungerik/go3d/vec3"
    "github.com/ungerik/go3d/mat4"
    "github.com/ungerik/go3d/quaternion"
    math "github.com/ungerik/go3d/fmath"
)

func main() {
    // Vector operations
    a := vec3.T{1, 2, 3}
    b := vec3.UnitX

    // In-place modification with chaining
    a.Add(&b).Scale(5)  // a is now {10, 10, 15}

    // Immutable operations
    c := a.Scaled(0.5)  // c = {5, 5, 7.5}, a unchanged

    // Matrix transformations
    var transform mat4.T
    transform.AssignPerspective(math.Pi/4, 16.0/9.0, 0.1, 1000.0)

    // Quaternions for rotations
    q := quaternion.FromYAxisAngle(math.Pi / 4)  // 45° around Y
    rotated := q.RotatedVec3(&a)
}
```

## Package Organization

Every type has its own sub-package and is named `T`. Packages under `float64/` use `float64` instead of `float32`.

### Main Packages (float32)

| Package | Description | Type Size |
|---------|-------------|-----------|
| `vec2` | 2D vectors | 8 bytes |
| `vec3` | 3D vectors | 12 bytes |
| `vec4` | 4D vectors / homogeneous coordinates | 16 bytes |
| `mat2` | 2×2 matrices | 16 bytes |
| `mat3` | 3×3 matrices | 36 bytes |
| `mat4` | 4×4 matrices | 64 bytes (one cache line!) |
| `quaternion` | Quaternions for 3D rotations | 16 bytes |
| `fmath` | Fast math functions (float32) | - |

### Float64 Packages

All types are available in float64 precision under `float64/`:

- `float64/vec2`, `float64/vec3`, `float64/vec4`
- `float64/mat2`, `float64/mat3`, `float64/mat4`
- `float64/quaternion`

### Utility Packages

- `generic` - Generic matrix/vector interfaces
- `hermit2` - 2D Hermite splines
- `hermit3` - 3D Hermite splines

## Core Concepts

### Naming Conventions

**Present tense methods** modify in place and return `*T` for chaining:

```go
v := vec3.T{1, 2, 3}
v.Scale(2)           // v is now {2, 4, 6}
v.Add(&vec3.UnitX)   // v is now {3, 4, 6}
```

**Past tense methods** return a modified copy:

```go
v := vec3.T{1, 2, 3}
v2 := v.Scaled(2)        // v2 = {2, 4, 6}, v unchanged
v3 := v.Added(&vec3.UnitX)  // v3 = {2, 2, 3}, v unchanged
```

### Method Chaining

```go
result := vec3.Zero
result.Add(&vec3.UnitX).Scale(5).Add(&vec3.UnitY)
// result = {5, 1, 0}
```

### Matrix Layout

Matrices use **column-major** layout (OpenGL convention):

```go
mat[column][row]
```

For DirectX (row-major), use `Transpose()`.

### Coordinate Systems

- **Right-handed coordinates**: X × Y = Z
- **Rotation direction**: Counter-clockwise (right-hand rule)
- **Angles**: Radians (use `fmath.Pi` constants)

## Common Operations

### Vectors

```go
// Construction
v := vec3.T{1, 2, 3}

// Length
length := v.Length()         // sqrt(x² + y² + z²)
lengthSq := v.LengthSqr()   // x² + y² + z² (faster, no sqrt)

// Normalization
v.Normalize()               // In-place
unit := v.Normalized()      // Copy

// Arithmetic
v.Add(&other)               // v += other
v.Sub(&other)               // v -= other
v.Scale(2.5)                // v *= 2.5

// Dot and cross products
dot := vec3.Dot(&a, &b)
cross := vec3.Cross(&a, &b)

// Distance
dist := vec3.Distance(&a, &b)
```

### Matrices

```go
// Identity and zero
m := mat4.Ident
z := mat4.Zero

// Translation
m.SetTranslation(&position)

// Rotation
m.AssignXRotation(angle)
m.AssignYRotation(angle)
m.AssignZRotation(angle)
m.AssignQuaternion(&quat)

// Scaling
m.AssignScaling(sx, sy, sz)

// Camera matrix
var view mat4.T
view.AssignLookAt(&eye, &center, &up)

// Projection matrices
var proj mat4.T

// Symmetric perspective (typical 3D game)
proj.AssignPerspective(fovy, aspect, znear, zfar)

// Asymmetric frustum
proj.AssignFrustum(left, right, bottom, top, znear, zfar)

// Orthographic (2D or CAD)
proj.AssignOrthogonalProjection(left, right, bottom, top, znear, zfar)

// Matrix multiplication
mvp := mat4.Ident
mvp.AssignMul(&projection, &view)
mvp.Mul(&model)
```

### Quaternions

```go
// Construction
axis := vec3.T{0, 1, 0}
q := quaternion.FromAxisAngle(&axis, math.Pi/4)

// Convenience constructors
qx := quaternion.FromXAxisAngle(angle)
qy := quaternion.FromYAxisAngle(angle)
qz := quaternion.FromZAxisAngle(angle)

// From Euler angles
q := quaternion.FromEulerAngles(yaw, pitch, roll)

// Rotate vector
v := vec3.T{1, 0, 0}
q.RotateVec3(&v)        // In-place
rotated := q.RotatedVec3(&v)  // Copy

// Combine rotations (order matters!)
q3 := quaternion.Mul(&q1, &q2)  // Apply q2 first, then q1

// Smooth interpolation
halfway := quaternion.Slerp(&start, &end, 0.5)

// Convert to matrix
var m mat4.T
m.AssignQuaternion(&q)
```

## Performance Tips

### 1. Use Appropriate Precision

```go
// For graphics (sufficient for most cases)
import "github.com/ungerik/go3d/vec3"

// For scientific computing (higher precision)
import "github.com/ungerik/go3d/float64/vec3"
```

### 2. Prefer In-Place Operations

```go
// ✓ Good: Minimal allocations
v.Scale(2).Add(&other)

// ✗ Less efficient: Creates copies
result := v.Scaled(2).Added(&other)
```

### 3. Use LengthSqr for Comparisons

```go
// ✓ Faster: No sqrt
thresholdSq := threshold * threshold
if v.LengthSqr() < thresholdSq {
    // ...
}

// ✗ Slower: Uses sqrt
if v.Length() < threshold {
    // ...
}
```

### 4. Pass Pointers to Functions

```go
// ✓ Efficient: No copy
dot := vec3.Dot(&a, &b)

// The API uses pointers for better performance
```

### 5. Reuse Matrix Allocations

```go
// ✓ Good: Single allocation reused
var transform mat4.T
for i := range objects {
    transform.AssignTranslation(&objects[i].Position)
    // Use transform...
}

// ✗ Bad: Allocates every iteration
for i := range objects {
    var transform mat4.T  // New allocation
    // ...
}
```

## Documentation

- **API Reference**: See [API_REFERENCE.md](API_REFERENCE.md) for complete API documentation
- **Code Critique**: See [CODE_CRITIQUE.md](CODE_CRITIQUE.md) for known issues and fixes
- **Go Documentation**: https://pkg.go.dev/github.com/ungerik/go3d

## Documentation Standards

### All exported names must be documented

Every exported function, type, constant, and variable must have a comment following Go conventions:

```go
// Normalize normalizes the vector to unit length.
// It modifies the vector in place and returns a pointer for chaining.
// The zero vector remains zero as it cannot be normalized.
func (vec *T) Normalize() *T
```

### Documentation Requirements

1. **First sentence**: Brief summary (appears in package listings)
2. **Parameters**: Describe non-obvious parameters
3. **Return values**: Explain what is returned
4. **Special cases**: Document edge cases, nil handling, panics
5. **Examples**: Include for complex operations

### Validation

Run the documentation checker:

```bash
./tools/check-docs.sh
```

Generate documentation:

```bash
# View in browser
godoc -http=:6060

# Generate text documentation
go doc -all > docs.txt
```

## Testing

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./vec3
go test ./mat4

# Run with race detector
go test -race ./...

# Run benchmarks
go test -bench=. ./vec3
```

## Examples

### Complete Transform Pipeline

```go
// Object transform
position := vec3.T{10, 5, 0}
rotation := quaternion.FromYAxisAngle(math.Pi / 4)
scale := vec3.T{2, 2, 2}

var model mat4.T
model.SetTranslation(&position)
model.AssignQuaternion(&rotation)
model.ScaleVec3(&scale)

// Camera
eye := vec3.T{0, 10, 20}
center := vec3.T{0, 0, 0}
up := vec3.UnitY

var view mat4.T
view.AssignLookAt(&eye, &center, &up)

// Projection
var projection mat4.T
projection.AssignPerspective(math.Pi/4, 16.0/9.0, 0.1, 1000.0)

// Combined MVP
var mvp mat4.T
mvp.AssignMul(&projection, &view)
mvp.Mul(&model)

// Transform vertex
vertex := vec4.T{1, 0, 0, 1}
transformed := mat4.MulVec4(&mvp, &vertex)
```

### Camera Rotation with Slerp

```go
func SmoothCamera(start, end quaternion.T, t float32) quaternion.T {
    // Clamp t to [0, 1]
    if t < 0 { t = 0 }
    if t > 1 { t = 1 }

    // Smooth interpolation
    return quaternion.Slerp(&start, &end, t)
}
```

## Design Philosophy

The package is designed for **performance over convenience** where necessary:

- **Pointer arguments**: Reduces allocations and copies
- **Stack allocation**: Value types (arrays) avoid GC pressure
- **Column-major matrices**: Direct GPU upload without transpose
- **Method chaining**: Fluent API without sacrificing performance

Performance comparison (see `mat4/mat4_test.go`):

```bash
cd mat4
go test -bench=BenchmarkMulAddVec4_PassBy*
```

## Contributing

Contributions are welcome! Please ensure:

1. All tests pass: `go test ./...`
2. Code is formatted: `go fmt ./...`
3. Documentation is updated
4. Benchmarks show no regressions

## License

MIT - see [LICENSE](LICENSE) file for details.
