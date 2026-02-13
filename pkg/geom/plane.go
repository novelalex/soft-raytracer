package geom

import (
	"math"

	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/rendering"
)

type Plane struct {
	Xf  nmath.Mat4
	Mat rendering.Material
}

func DefaultPlane() Plane {
	return Plane{
		nmath.Mat4Identity(),
		rendering.DefaultMaterial(),
	}
}
func NewPlane(t nmath.Mat4, m rendering.Material) Plane {
	return Plane{
		t,
		m,
	}
}

func (p Plane) Transform() nmath.Mat4 {
	return p.Xf
}

func (p Plane) Material() rendering.Material {
	return p.Mat
}

func (p *Plane) SetTransform(m nmath.Mat4) {
	p.Xf = m
}

func (p *Plane) SetMaterial(m rendering.Material) {
	p.Mat = m
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
