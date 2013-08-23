package vec2d

type Rect struct {
	Min T
	Max T
}

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
