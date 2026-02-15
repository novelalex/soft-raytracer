package geom

import (
	"github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Shape interface {
	Transform() nmath.Mat4
	SetTransform(nmath.Mat4)
	IntersectRay(Ray) []float64
	NormalAt(world_point nmath.Vec3) nmath.Vec3
}
