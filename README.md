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
    math "github.com/chewxy/math32"
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

### Float64 Packages

All types are available in float64 precision under `float64/`:

- `float64/vec2`, `float64/vec3`, `float64/vec4`
- `float64/mat2`, `float64/mat3`, `float64/mat4`
- `float64/quaternion`

### Utility Packages

- `generic` - Generic matrix/vector interfaces
- `hermit2` - 2D Hermite splines
- `hermit3` - 3D Hermite splines

### Float32 Math Functions

This library uses [github.com/chewxy/math32](https://github.com/chewxy/math32) for float32 math functions. Import it as:

```go
import math "github.com/chewxy/math32"
```

The math32 package provides float32 versions of Go's standard `math` library functions, offering better performance for float32-based graphics calculations compared to converting to/from float64.

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
- **Angles**: Radians (use `math32.Pi` constants)

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

## API Reference

### Rectangle (vec2 package)

```go
type Rect struct {
    Min vec2.T
    Max vec2.T
}
```

**Operations:**
```go
r := vec2.Rect{Min: vec2.T{0, 0}, Max: vec2.T{100, 100}}

width := r.Width()
height := r.Height()
center := r.Center()
area := r.Area()

joined := r.Join(&other)        // Bounding rectangle
clamped := r.Clamp(&point)      // Clamp point to rectangle
contains := r.Contains(&point)   // Point-in-rectangle test
```

### Box (vec3 package)

```go
type Box struct {
    Min vec3.T
    Max vec3.T
}
```

**Operations:**
```go
b := vec3.Box{Min: vec3.T{0, 0, 0}, Max: vec3.T{10, 10, 10}}

size := b.Size()
center := b.Center()
volume := b.Volume()

joined := b.Join(&other)        // Bounding box
contains := b.Contains(&point)   // Point-in-box test
```

### Migration Notes

#### Matrix Multiplication Order

Transforms apply **right to left**:

```go
// Mathematical notation: MVP = P × V × M
// In go3d:
var mvp mat4.T
mvp.AssignMul(&projection, &view)  // mvp = P × V
mvp.Mul(&model)                     // mvp = (P × V) × M
```

#### Vec4 Homogeneous Coordinates

```go
// Point in space (affected by translation): w=1
point := vec4.T{x, y, z, 1.0}

// Direction vector (not affected by translation): w=0
direction := vec4.T{x, y, z, 0.0}
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

### Billboard Matrix

Make objects always face the camera:

```go
func CreateBillboard(objectPos, cameraPos, cameraUp vec3.T) mat4.T {
    // Calculate direction to camera
    direction := vec3.Sub(&cameraPos, &objectPos)
    direction.Normalize()

    // Calculate right vector
    right := vec3.Cross(&cameraUp, &direction)
    right.Normalize()

    // Recalculate up vector
    up := vec3.Cross(&direction, &right)

    // Build matrix
    var billboard mat4.T
    billboard[0] = vec4.T{right[0], right[1], right[2], 0}
    billboard[1] = vec4.T{up[0], up[1], up[2], 0}
    billboard[2] = vec4.T{direction[0], direction[1], direction[2], 0}
    billboard[3] = vec4.T{objectPos[0], objectPos[1], objectPos[2], 1}

    return billboard
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
