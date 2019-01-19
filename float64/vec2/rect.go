package vec2

import (
	"fmt"
)

// Rect is a coordinate system aligned rectangle defined by a Min and Max vector.
type Rect struct {
	Min T
	Max T
}

// NewRect creates a Rect from two points.
func NewRect(a, b *T) (rect Rect) {
	rect.Min = Min(a, b)
	rect.Max = Max(a, b)
	return rect
}

// ParseRect parses a Rect from a string. See also String()
func ParseRect(s string) (r Rect, err error) {
	_, err = fmt.Sscan(s, &r.Min[0], &r.Min[1], &r.Max[0], &r.Max[1])
	return r, err
}

// String formats Rect as string. See also ParseRect().
func (rect *Rect) String() string {
	return rect.Min.String() + " " + rect.Max.String()
}

// ContainsPoint returns if a point is contained within the rectangle.
func (rect *Rect) ContainsPoint(p *T) bool {
	return p[0] >= rect.Min[0] && p[0] <= rect.Max[0] &&
		p[1] >= rect.Min[1] && p[1] <= rect.Max[1]
}

// Contains returns if other Rect is contained within the rectangle.
func (rect *Rect) Contains(other *Rect) bool {
	return rect.Min[0] <= other.Min[0] &&
		rect.Min[1] <= other.Min[1] &&
		rect.Max[0] >= other.Max[0] &&
		rect.Max[1] >= other.Max[1]
}

// Area calculates the area of the rectangle.
func (rect *Rect) Area() float64 {
	return (rect.Max[0] - rect.Min[0]) * (rect.Max[1] - rect.Min[1])
}

// func (rect *Rect) Intersects(other *Rect) bool {
// 	panic("not implemented")
// }

// func Intersect(a, b *Rect) Rect {
// 	panic("not implemented")
// }

// Join enlarges this rectangle to contain also the given rectangle.
func (rect *Rect) Join(other *Rect) {
	rect.Min = Min(&rect.Min, &other.Min)
	rect.Max = Max(&rect.Max, &other.Max)
}

// Joined returns the minimal rectangle containing both a and b.
func Joined(a, b *Rect) (rect Rect) {
	rect.Min = Min(&a.Min, &b.Min)
	rect.Max = Max(&a.Max, &b.Max)
	return rect
}
