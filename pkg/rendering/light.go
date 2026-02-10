package rendering

import "github.com/novelalex/soft-raytracer/pkg/nmath"

type PointLight struct {
	Position  nmath.Vec3
	Intensity nmath.Color
}

func NewPointLight(position nmath.Vec3, intensity nmath.Color) PointLight {
	return PointLight{position, intensity}
}
