package renderer

import "github.com/novelalex/soft-raytracer/pkg/nmath"

type PointLight struct {
	position  nmath.Vec3
	intensity nmath.Color
}

func NewPointLight(position nmath.Vec3, intensity nmath.Color) PointLight {
	return PointLight{position, intensity}
}
