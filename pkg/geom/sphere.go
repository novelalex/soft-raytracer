package geom

import (
	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/rendering"
)

type Sphere struct {
	transform nmath.Mat4
	material  rendering.Material
}

func NewSphere(t nmath.Mat4, m rendering.Material) Sphere {
	return Sphere{
		t,
		m,
	}
}

func (s Sphere) Transform() nmath.Mat4 {
	return s.transform
}

func (s Sphere) Material() rendering.Material {
	return s.material
}

func (s *Sphere) SetTransform(m nmath.Mat4) {
	s.transform = m
}

func (s *Sphere) SetMaterial(m rendering.Material) {
	s.material = m
}

func (s *Sphere) Translate(x, y, z float64) *Sphere {
	s.SetTransform(s.transform.Mult(nmath.NewTranslation(x, y, z)))
	return s
}

func (s *Sphere) Scale(x, y, z float64) *Sphere {
	s.SetTransform(s.transform.Mult(nmath.NewScaling(x, y, z)))
	return s
}

func (s *Sphere) Rotate(angle float64, axis nmath.Vec3) *Sphere {
	s.SetTransform(s.transform.Mult(nmath.NewRotation(angle, axis)))
	return s
}

func (s *Sphere) RotateX(angle float64) *Sphere {
	s.SetTransform(s.transform.Mult(nmath.NewRotationX(angle)))
	return s
}

func (s *Sphere) RotateY(angle float64) *Sphere {
	s.SetTransform(s.transform.Mult(nmath.NewRotationY(angle)))
	return s
}

func (s *Sphere) RotateZ(angle float64) *Sphere {
	s.SetTransform(s.transform.Mult(nmath.NewRotationZ(angle)))
	return s
}

func (s Sphere) NormalAt(world_point nmath.Vec3) nmath.Vec3 {
	object_point := s.transform.Inverse().MultV(world_point.AsPoint4())
	object_normal := object_point.Sub(nmath.NewPoint4(0, 0, 0))
	world_normal := s.transform.Inverse().Transpose().MultV(object_normal)
	return world_normal.Normalize().DropW()
}
