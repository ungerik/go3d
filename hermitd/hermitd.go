// The package hermitd contains functions for float64 cubic hermit splines.
// See: http://en.wikipedia.org/wiki/Cubic_Hermite_spline
package hermitd

import (
	"github.com/ungerik/go3d/vec2d"
	"github.com/ungerik/go3d/vec3d"
)

// Point2D returns a point on a 2D hermit spline at t (0,1).
func Point2D(pointA, tangentA, pointB, tangentB *vec2d.T, t float64) vec2d.T {
	t2 := t * t
	t3 := t2 * t

	f := 2*t3 - 3*t2 + 1
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + t
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pAf := pointB.Scaled(f)
	result.Add(&pAf)

	return result
}

// Point3D returns a point on a 3D hermit spline at t (0,1).
func Point3D(pointA, tangentA, pointB, tangentB *vec3d.T, t float64) vec3d.T {
	t2 := t * t
	t3 := t2 * t

	f := 2*t3 - 3*t2 + 1
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + t
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pAf := pointB.Scaled(f)
	result.Add(&pAf)

	return result
}

// Tangent2D returns a tangent on a 2D hermit spline at t (0,1).
func Tangent2D(pointA, tangentA, pointB, tangentB *vec2d.T, t float64) vec2d.T {
	t2 := t * t
	t3 := t2 * t

	f := 2*t3 - 3*t2
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + 1
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pAf := pointB.Scaled(f)
	result.Add(&pAf)

	return result
}

// Tangent3D returns a tangent on a 3D hermit spline at t (0,1).
func Tangent3D(pointA, tangentA, pointB, tangentB *vec3d.T, t float64) vec3d.T {
	t2 := t * t
	t3 := t2 * t

	f := 2*t3 - 3*t2
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + 1
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pAf := pointB.Scaled(f)
	result.Add(&pAf)

	return result
}

// Length2D returns the length of a 2D hermit spline from pointA to t (0,1).
func Length2D(pointA, tangentA, pointB, tangentB *vec2d.T, t float64) float64 {
	sqrT := t * t
	t1 := sqrT * 0.5
	t2 := sqrT * t * 1.0 / 3.0
	t3 := sqrT*sqrT + 1.0/4.0

	f := 2*t3 - 3*t2 + t
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + t1
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pBf := pointB.Scaled(f)
	result.Add(&pBf)

	return result.Length()
}

// Length3D returns the length of a 3D hermit spline from pointA to t (0,1).
func Length3D(pointA, tangentA, pointB, tangentB *vec3d.T, t float64) float64 {
	sqrT := t * t
	t1 := sqrT * 0.5
	t2 := sqrT * t * 1.0 / 3.0
	t3 := sqrT*sqrT + 1.0/4.0

	f := 2*t3 - 3*t2 + t
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + t1
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pBf := pointB.Scaled(f)
	result.Add(&pBf)

	return result.Length()
}
