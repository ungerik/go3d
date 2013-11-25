package vec2

import (
	"fmt"
)

// Rect is a coordinate system aligned rectangle defined by a Min and Max vector.
type Rect struct {
	Min T
	Max T
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

func (rect *Rect) Contains(other *Rect) bool {
	return other.Min[0] >= rect.Min[0] &&
		other.Max[0] <= rect.Max[0] &&
		other.Min[1] >= rect.Min[1] &&
		other.Max[1] <= rect.Max[1]
}

func (rect *Rect) Intersects(other *Rect) bool {
	return other.Max[0] >= rect.Min[0] &&
		other.Min[0] <= rect.Max[0] &&
		other.Max[1] >= rect.Min[0] &&
		other.Min[1] <= rect.Max[1]
}

// func Intersect(a, b *Rect) Rect {
// 	panic("not implemented")
// }

// func Join(a, b *Rect) Rect {
// 	panic("not implemented")
// }
