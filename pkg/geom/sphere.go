package geom

import (
	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/rendering"
)

type Sphere struct {
	Xf  nmath.Mat4
	Mat rendering.Material
}

func DefaultSphere() Sphere {
	return Sphere{
		nmath.Mat4Identity(),
		rendering.DefaultMaterial(),
	}
}
func NewSphere(t nmath.Mat4, m rendering.Material) Sphere {
	return Sphere{
		t,
		m,
	}
}

func (s Sphere) Transform() nmath.Mat4 {
	return s.Xf
}

func (s Sphere) Material() rendering.Material {
	return s.Mat
}

func (s *Sphere) SetTransform(m nmath.Mat4) {
	s.Xf = m
}

func (s *Sphere) SetMaterial(m rendering.Material) {
	s.Mat = m
}

func (s *Sphere) Translate(x, y, z float64) *Sphere {
	s.SetTransform(s.Xf.Mult(nmath.NewTranslation(x, y, z)))
	return s
}

func (s *Sphere) Scale(x, y, z float64) *Sphere {
	s.SetTransform(s.Xf.Mult(nmath.NewScaling(x, y, z)))
	return s
}

func (s *Sphere) Rotate(angle float64, axis nmath.Vec3) *Sphere {
	s.SetTransform(s.Xf.Mult(nmath.NewRotation(angle, axis)))
	return s
}

func (s *Sphere) RotateX(angle float64) *Sphere {
	s.SetTransform(s.Xf.Mult(nmath.NewRotationX(angle)))
	return s
}

func (s *Sphere) RotateY(angle float64) *Sphere {
	s.SetTransform(s.Xf.Mult(nmath.NewRotationY(angle)))
	return s
}

func (s *Sphere) RotateZ(angle float64) *Sphere {
	s.SetTransform(s.Xf.Mult(nmath.NewRotationZ(angle)))
	return s
}

func (s Sphere) NormalAt(world_point nmath.Vec3) nmath.Vec3 {
	object_point := s.Xf.Inverse().MultV(world_point.AsPoint4())
	object_normal := object_point.Sub(nmath.NewPoint4(0, 0, 0))
	world_normal := s.Xf.Inverse().Transpose().MultV(object_normal)
	return world_normal.Normalize().DropW()
}

func (s Sphere) IntersectRay(r Ray) Intersections {
	return r.IntersectSphere(s)
}
