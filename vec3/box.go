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

// Intersects returns true if this and the given box intersect. 
// For an explanation of the algorithm, see
// http://rbrundritt.wordpress.com/2009/10/03/determining-if-two-bounding-boxes-overlap/
func (box *Box) Intersects(other *Box) bool {
	d1 := box.Diagonal()
	d2 := other.Diagonal()
	sizes := Add(&d1, &d2)
	c1 := box.Center()
	c2 := other.Center()
	distCenters2 := Sub(&c1, &c2)
	distCenters2.Scale(2)
	distCenters2.Abs()
	return distCenters2[0] <= sizes[0] && distCenters2[1] <= sizes[1] && distCenters2[2] <= sizes[2]
}

// Join enlarges this box to contain also the given box.
func (box *Box) Join(other *Box) {
	box.Min = Min(&box.Min, &other.Min)
	box.Max = Max(&box.Max, &other.Max)
}

// Joined returns the minimal box containing both a and b.
func Joined(a, b *Box) Box {
	var joined Box
	joined.Min = Min(&a.Min, &b.Min)
	joined.Max = Max(&a.Max, &b.Max)
	return joined
}
