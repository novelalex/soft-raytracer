package geom

import (
	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/renderer"
)

type Sphere struct {
	transform nmath.Mat4
	material  renderer.Material
}

func NewSphere(t nmath.Mat4, m renderer.Material) Sphere {
	return Sphere{
		t,
		m,
	}
}

func (s Sphere) Transform() nmath.Mat4 {
	return s.transform
}

func (s Sphere) Material() renderer.Material {
	return s.material
}

func (s Sphere) WithTransform(m nmath.Mat4) Sphere {
	return Sphere{
		m.Mult(s.transform),
		s.material,
	}
}

func (s Sphere) WithMaterial(m renderer.Material) Sphere {
	return Sphere{
		s.transform,
		m,
	}
}

func (s Sphere) Translate(x, y, z float64) Sphere {
	return s.WithTransform(nmath.NewTranslation(x, y, z))
}

func (s Sphere) Scale(x, y, z float64) Sphere {
	return s.WithTransform(nmath.NewScaling(x, y, z))
}

func (s Sphere) Rotate(angle float64, axis nmath.Vec3) Sphere {
	return s.WithTransform(nmath.NewRotation(angle, axis))
}

func (s Sphere) RotateX(angle float64) Sphere {
	return s.WithTransform(nmath.NewRotationX(angle))
}

func (s Sphere) RotateY(angle float64) Sphere {
	return s.WithTransform(nmath.NewRotationY(angle))
}

func (s Sphere) RotateZ(angle float64) Sphere {
	return s.WithTransform(nmath.NewRotationZ(angle))
}

func (s Sphere) NormalAt(world_point nmath.Vec3) nmath.Vec3 {
	object_point := s.transform.Inverse().MultV(world_point.AsPoint4())
	object_normal := object_point.Sub(nmath.NewPoint4(0, 0, 0))
	world_normal := s.transform.Inverse().Transpose().MultV(object_normal)
	return world_normal.Normalize().DropW()
}
