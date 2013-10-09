package vec3

import (
	"fmt"
)

type Box struct {
	Min T
	Max T
}

// ParseBox parses a Box from a string. See also String()
func ParseBox(s string) (r Box, err error) {
	_, err = fmt.Sscanf(s, "%f %f %f %f %f %f", &r.Min[0], &r.Min[1], &r.Min[2], &r.Max[0], &r.Max[1], &r.Max[2])
	return r, err
}

// String formats Box as string. See also ParseBox().
func (self *Box) String() string {
	return self.Min.String() + " " + self.Max.String()
}

func (self *Box) ContainsPoint(p *T) bool {
	return p[0] >= self.Min[0] && p[0] <= self.Max[0] &&
		p[1] >= self.Min[1] && p[1] <= self.Max[1] &&
		p[2] >= self.Min[2] && p[2] <= self.Max[2]
}

func (self *Box) Contains(other *Box) bool {
	panic("not implemented")
}

func (self *Box) Intersects(other *Box) bool {
	panic("not implemented")
}

func Intersect(a, b *Box) Box {
	panic("not implemented")
}

func Join(a, b *Box) Box {
	panic("not implemented")
}
