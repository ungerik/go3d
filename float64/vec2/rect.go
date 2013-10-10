package vec2

import (
	"fmt"
)

// Rect is a rectangle defined by Min and Max vector corners.
type Rect struct {
	Min T
	Max T
}

// ParseRect parses a Rect from a string. See also String()
func ParseRect(s string) (r Rect, err error) {
	_, err = fmt.Sscanf(s, "%f %f %f %f", &r.Min[0], &r.Min[1], &r.Max[0], &r.Max[1])
	return r, err
}

// String formats Rect as string. See also ParseRect().
func (self *Rect) String() string {
	return self.Min.String() + " " + self.Max.String()
}

// ContainsPoint returns if a point is within the rectangle.
func (self *Rect) ContainsPoint(p *T) bool {
	return p[0] >= self.Min[0] && p[0] <= self.Max[0] &&
		p[1] >= self.Min[1] && p[1] <= self.Max[1]
}

func (self *Rect) Contains(other *Rect) bool {
	panic("not implemented")
}

func (self *Rect) Intersects(other *Rect) bool {
	panic("not implemented")
}

func Intersect(a, b *Rect) Rect {
	panic("not implemented")
}

func Join(a, b *Rect) Rect {
	panic("not implemented")
}
