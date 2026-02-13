package geom

import (
	"math"
	. "github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Pattern interface {
	At(Vec3) Color
	AtObject(Shape, Vec3) Color
	Transform() Mat4
	SetTransform(Mat4)
}

type StripePattern struct {
	A  Color
	B  Color
	Xf Mat4
}

func NewStripePattern(a, b Color) StripePattern {
	return StripePattern{a, b, Mat4Identity()}
}

func (s StripePattern) Transform() Mat4 {
	return s.Xf
}

func (s *StripePattern) SetTransform(m Mat4) {
	s.Xf = m
}

func (s StripePattern) AtObject(obj Shape, p Vec3) Color {
	object_point := obj.Transform().Inverse().MultV(p.AsPoint4())
	pattern_point := s.Transform().Inverse().MultV(object_point).DropW()
	return s.At(pattern_point)
}

func (s StripePattern) At(point Vec3) Color {
	if ApproxEq(math.Mod(math.Floor(point.X), 2), 0) {
		return s.A
	}
	return s.B
}
