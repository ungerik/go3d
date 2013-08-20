package vec3

type Box struct {
	Min T
	Max T
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
