package go3d

import (
	"os"
	"math"
	"bytes"
	"strings"
	"strconv"
)


type Vec3 struct {
	X, Y, Z float64
}


func (self *Vec3) Field(index int) *float64 {
	switch index {
	case 0:
		return &self.X
	case 1:
		return &self.Y
	case 2:
		return &self.Z
	}
	panic("Out of Vec3 index range [0..2]")
	return nil
}


func (self *Vec3) Slice() []float64 {
	return []float64{self.X, self.Y, self.Z}
}


func (self *Vec3) SetSlice(slice []float64) {
	self.X = slice[0]
	self.Y = slice[1]
	self.Z = slice[2]
}


func (self *Vec3) String() string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Ftoa64(self.X, 'f', -1))
	buf.WriteByte(' ')
	buf.WriteString(strconv.Ftoa64(self.Y, 'f', -1))
	buf.WriteByte(' ')
	buf.WriteString(strconv.Ftoa64(self.Z, 'f', -1))
	return buf.String()
}


func ParseVec3(str string) (vec3 Vec3, err os.Error) {
	for i, s := range strings.Split(str, " ", 3) {
		*vec3.Field(i), err = strconv.Atof64(s)
		if err != nil {
			return
		}
	}
	return
}


func (self *Vec3) Zero() bool {
	return self.X == 0 && self.Y == 0 && self.Z == 0
}


func (self *Vec3) SetZero() *Vec3 {
	self.X = 0
	self.Y = 0
	self.Z = 0
	return self
}


func (self *Vec3) SetMinimum() *Vec3 {
	self.X = -math.MaxFloat64
	self.Y = -math.MaxFloat64
	self.Z = -math.MaxFloat64
	return self
}


func (self *Vec3) SetMaximum() *Vec3 {
	self.X = math.MaxFloat64
	self.Y = math.MaxFloat64
	self.Z = math.MaxFloat64
	return self
}


func (self *Vec3) SetUnitX() *Vec3 {
	self.X = 1
	self.Y = 0
	self.Z = 0
	return self
}


func (self *Vec3) SetUnitY() *Vec3 {
	self.X = 0
	self.Y = 1
	self.Z = 0
	return self
}


func (self *Vec3) SetUnitZ() *Vec3 {
	self.X = 0
	self.Y = 0
	self.Z = 1
	return self
}


func (self *Vec3) SquareLength() float64 {
	return self.X * self.X + self.Y * self.Y + self.Z * self.Z
}


func (self *Vec3) Length() float64 {
	return math.Sqrt(self.SquareLength())
}


func (self *Vec3) UnitLength() bool {
	return EqualSquareEpsilon(self.SquareLength(), 1)
}


func (self *Vec3) Invert() *Vec3 {
	self.X = -self.X
	self.Y = -self.Y
	self.Z = -self.Z
	return self
}


func (self *Vec3) Inverted() Vec3 {
	result := *self
	result.Invert()
	return result
}


func (self *Vec3) Normalize() *Vec3 {
	squareLen := self.SquareLength()
	if squareLen != 0 && squareLen != 1 {
		self.Scale(InvSqrt(squareLen))
	}
	return self
}


func (self *Vec3) Normalized() Vec3 {
	result := *self
	result.Normalize()
	return result
}


func (self *Vec3) Scale(factor float64) *Vec3 {
	self.X *= factor
	self.Y *= factor
	self.Z *= factor
	return self
}


func (self *Vec3) Scaled(factor float64) Vec3 {
	result := *self
	result.Scale(factor)
	return result
}


func (self *Vec3) Add(other *Vec3) *Vec3 {
	self.X += other.X
	self.Y += other.Y
	self.Z += other.Z
	return self
}


func (self *Vec3) Added(other *Vec3) Vec3 {
	result := *self
	result.Add(other)
	return result
}


func (self *Vec3) Subtract(other *Vec3) *Vec3 {
	self.X -= other.X
	self.Y -= other.Y
	self.Z -= other.Z
	return self
}


func (self *Vec3) Subtracted(other *Vec3) Vec3 {
	result := *self
	result.Subtract(other)
	return result
}

