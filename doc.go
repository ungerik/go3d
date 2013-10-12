/*
go3d is a performance oriented vector and matrix math package for 2D and 3D graphics.

Every type has its own sub-package and is named T. So vec3.T is the 3D vector type.
For every vector and matrix type there is a String() method and a Parse() function.
Besides methods of T there are also functions in the packages, like vec3.Dot(a, b).

Packages under the float64 directory are using float64 values instead of float32.

Matrices are organized as arrays of columns which is also the way OpenGL expects matrices.
DirectX expects "arrays of rows" matrices, use the Transpose() to convert.

Methods that don't return a specific value, return a pointer to the struct to allow method call chaining.

Example:

	a := vec3.Zero
	b := vec3.UnitX
	a.Add(&b).Scale(5)

Method names in the past tense return a copy of the struct instead of a pointer to it.

Example:

	a := vec3.UnitX
	b := a.Scaled(5) // a still equals vec3.UnitX

*/
package go3d

// Import all sub-packages for build
import (
	_ "github.com/ungerik/go3d/float64/generic"
	_ "github.com/ungerik/go3d/float64/hermit2"
	_ "github.com/ungerik/go3d/float64/hermit3"
	_ "github.com/ungerik/go3d/float64/mat2x2"
	_ "github.com/ungerik/go3d/float64/mat3x3"
	_ "github.com/ungerik/go3d/float64/mat4x4"
	_ "github.com/ungerik/go3d/float64/quaternion"
	_ "github.com/ungerik/go3d/float64/vec2"
	_ "github.com/ungerik/go3d/float64/vec3"
	_ "github.com/ungerik/go3d/float64/vec4"
	_ "github.com/ungerik/go3d/generic"
	_ "github.com/ungerik/go3d/hermit2"
	_ "github.com/ungerik/go3d/hermit3"
	_ "github.com/ungerik/go3d/mat2x2"
	_ "github.com/ungerik/go3d/mat3x3"
	_ "github.com/ungerik/go3d/mat4x4"
	_ "github.com/ungerik/go3d/quaternion"
	_ "github.com/ungerik/go3d/vec2"
	_ "github.com/ungerik/go3d/vec3"
	_ "github.com/ungerik/go3d/vec4"
)
