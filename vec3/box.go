package vec3

import (
	"fmt"
)

// Box is a coordinate system aligned 3D box defined by a Min and Max vector.
type Box struct {
	Min T
	Max T
}

var (
	// MaxBox holds a box that contains the entire R3 space that can be represented as vec3
	MaxBox = Box{MinVal, MaxVal}
) 

// ParseBox parses a Box from a string. See also String()
func ParseBox(s string) (r Box, err error) {
	_, err = fmt.Sscan(s, &r.Min[0], &r.Min[1], &r.Min[2], &r.Max[0], &r.Max[1], &r.Max[2])
	return r, err
}

// String formats Box as string. See also ParseBox().
func (box *Box) String() string {
	return box.Min.String() + " " + box.Max.String()
}

// ContainsPoint returns if a point is contained within the box.
func (box *Box) ContainsPoint(p *T) bool {
	return p[0] >= box.Min[0] && p[0] <= box.Max[0] &&
		p[1] >= box.Min[1] && p[1] <= box.Max[1] &&
		p[2] >= box.Min[2] && p[2] <= box.Max[2]
}

func (box *Box) Center() T {
	c := Add(&box.Min, &box.Max)
	c.Scale(0.5)
	return c
}

func (box *Box) Diagonal() T {
	return Sub(&box.Max, &box.Min)
}

// func (box *Box) Contains(other *Box) bool {
// 	panic("not implemented")
// }

// func (box *Box) Intersects(other *Box) bool {
// 	panic("not implemented")
// }

// func Intersect(a, b *Box) Box {
// 	panic("not implemented")
// }

// func Join(a, b *Box) Box {
// 	panic("not implemented")
// }
