package hermit

import (
	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/vec3"
)

func Point2D(pointA, tangentA, pointB, tangentB *vec2.T, t float32) vec2.T {
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

func Point3D(pointA, tangentA, pointB, tangentB *vec3.T, t float32) vec3.T {
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

func Tangent2D(pointA, tangentA, pointB, tangentB *vec2.T, t float32) vec2.T {
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

func Tangent3D(pointA, tangentA, pointB, tangentB *vec3.T, t float32) vec3.T {
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

func Length2D(pointA, tangentA, pointB, tangentB *vec2.T, t float32) float32 {
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

func Length3D(pointA, tangentA, pointB, tangentB *vec3.T, t float32) float32 {
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
