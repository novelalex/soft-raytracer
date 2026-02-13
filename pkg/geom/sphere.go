package geom

import (
	"sync/atomic"

	"github.com/novelalex/soft-raytracer/pkg/nmath"
)

var nextID uint64

type Sphere struct {
	Xf nmath.Mat4
	id uint64
}

func newId() uint64 {
	return atomic.AddUint64(&nextID, 1)
}

func DefaultSphere() Sphere {

	return Sphere{
		nmath.Mat4Identity(),
		newId(),
	}
}
func NewSphere(t nmath.Mat4) Sphere {
	return Sphere{
		t,
		newId(),
	}
}

func (s Sphere) ID() uint64 {
	return s.id
}

func (s Sphere) Transform() nmath.Mat4 {
	return s.Xf
}

func (s *Sphere) SetTransform(m nmath.Mat4) {
	s.Xf = m
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
	object_normal := object_point.Sub(nmath.NewPoint4(0, 0, 0)).DropW()
	world_normal := s.Xf.Inverse().Transpose().MultV(object_normal.AsVector4())
	return world_normal.DropW().Normalize()
}

func (s Sphere) IntersectRay(r Ray) Intersections {
	return r.IntersectSphere(s)
}
