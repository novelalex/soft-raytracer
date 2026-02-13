package world

import (
	"github.com/novelalex/soft-raytracer/pkg/geom"
	"github.com/novelalex/soft-raytracer/pkg/rendering"
)

type Object struct {
	Shape    geom.Shape
	Material rendering.Material
}

func NewObject(s geom.Shape, m rendering.Material) Object {
	return Object{
		s,
		m,
	}
}
