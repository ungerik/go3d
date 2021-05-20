Package go3d is a performance oriented vector and matrix math package for 2D and 3D graphics.

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


Note that the package is designed for performance over usability where necessary.
This is the reason why many arguments are passed by pointer reference instead of by value.
Sticking either to passing and returning everything by value or by pointer
would lead to a nicer API, but it would be slower as demonstrated in mat4/mat4_test.go

	cd mat4
	go test -bench=BenchmarkMulAddVec4_PassBy*


Documentation: https://pkg.go.dev/github.com/ungerik/go3d

