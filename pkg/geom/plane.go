package geom

import (
	"math"

	"github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Plane struct {
	Xf nmath.Mat4
	id uint64
}

func DefaultPlane() Plane {
	return Plane{
		nmath.Mat4Identity(),
		newId(),
	}
}
func NewPlane(t nmath.Mat4) Plane {
	return Plane{
		t,
		newId(),
	}
}

func (p Plane) ID() uint64 {
	return p.id
}

func (p Plane) Transform() nmath.Mat4 {
	return p.Xf
}

func (p *Plane) SetTransform(m nmath.Mat4) {
	p.Xf = m
}

func (p *Plane) Translate(x, y, z float64) *Plane {
	p.SetTransform(p.Xf.Mult(nmath.NewTranslation(x, y, z)))
	return p
}

func (p Plane) NormalAt(point nmath.Vec3) nmath.Vec3 {
	object_normal := nmath.NewVec3(0, 1, 0)
	world_normal := p.Xf.Inverse().Transpose().MultV(object_normal.AsVector4()).DropW()
	return world_normal.Normalize()
}

func (p Plane) IntersectRay(r Ray) Intersections {
	ray := r.Transform(p.Xf.Inverse())

	if math.Abs(ray.Dir.Y) < nmath.F64Epsilon {
		return NewIntersections([]Intersection{})
	}
	t := -ray.Origin.Y / ray.Dir.Y
	return NewIntersections([]Intersection{
		NewIntersection(t, &p),
	})
}
