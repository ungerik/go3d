// Package qbezier2 contains a float32 type T and functions for 2D quadratic Bezier splines.
// See: http://en.wikipedia.org/wiki/B%C3%A9zier_curve
package qbezier2

import (
	"fmt"

	math "github.com/chewxy/math32"
	"github.com/ungerik/go3d/vec2"
)

// T holds the data to define a quadratic bezier spline.
type T struct {
	P0, P1, P2 vec2.T
}

// Parse parses T from a string. See also String()
func Parse(s string) (r T, err error) {
	_, err = fmt.Sscan(s,
		&r.P0[0], &r.P0[1],
		&r.P1[0], &r.P1[1],
		&r.P2[0], &r.P2[1],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (bez *T) String() string {
	return fmt.Sprintf("%s %s %s",
		bez.P0.String(), bez.P1.String(), bez.P2.String(),
	)
}

// Point returns a point on a quadratic bezier spline at t (0,1).
func (bez *T) Point(t float32) vec2.T {
	return Point(&bez.P0, &bez.P1, &bez.P2, t)
}

// Tangent returns a tangent on a quadratic bezier spline at t (0,1).
func (bez *T) Tangent(t float32) vec2.T {
	return Tangent(&bez.P0, &bez.P1, &bez.P2, t)
}

// Length returns the length of a quadratic bezier spline from P0 to t (0,1).
func (bez *T) Length(t float32) float32 {
	return Length(&bez.P0, &bez.P1, &bez.P2, t)
}

// Point returns a point on a quadratic bezier spline at t (0,1).
func Point(p0, p1, p2 *vec2.T, t float32) vec2.T {
	t1 := 1.0 - t

	f := t1 * t1
	result := p0.Scaled(f)

	f = 2.0 * t1 * t
	p1f := p1.Scaled(f)
	result.Add(&p1f)

	f = t * t
	p2f := p2.Scaled(f)
	result.Add(&p2f)

	return result
}

// Tangent returns a tangent on a quadratic bezier spline at t (0,1).
func Tangent(p0, p1, p2 *vec2.T, t float32) vec2.T {
	t1 := 1.0 - t

	f := 2.0 * t1
	p1f := vec2.Sub(p1, p0)
	result := p1f.Scaled(f)

	f = 2.0 * t
	p2f := vec2.Sub(p2, p1)
	p2f.Scale(f)
	result.Add(&p2f)

	if result[0] == 0 && result[1] == 0 {
		fmt.Printf("zero tangent!  p0=%v, p1=%v, p2=%v, t=%v\n", p0, p1, p2, t)
		panic("zero tangent of qbezier2")
	}

	return result
}

// Length returns the length of a quadratic bezier spline from p0 to t (0,1).
//
// Note that although this calculation is accurate at t=0, 0.5, and 1 due
// to the nature of quadratic curves, it is an approximation for other values of t.
func Length(p0, p1, p2 *vec2.T, t float32) float32 {
	ax := p0[0] - 2*p1[0] + p2[0]
	ay := p0[1] - 2*p1[1] + p2[1]
	bx := 2*p1[0] - 2*p0[0]
	by := 2*p1[1] - 2*p0[1]

	a := 4 * (ax*ax + ay*ay)
	b := 4 * (ax*bx + ay*by)
	c := bx*bx + by*by

	abc := 2 * math.Sqrt(a+b+c)
	a2 := math.Sqrt(a)
	a32 := 2 * a * a2
	c2 := 2 * math.Sqrt(c)
	ba := b / a2

	return t * (a32*abc + a2*b*(abc-c2) + (4*c*a-b*b)*math.Log((2*a2+ba+abc)/(ba+c2))) / (4 * a32)
}
