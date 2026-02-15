package geom

import (
	"math"

	"github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Sphere struct {
	Xf nmath.Mat4
}

func DefaultSphere() Sphere {

	return Sphere{
		nmath.Mat4Identity(),
	}
}
func NewSphere(t nmath.Mat4) Sphere {
	return Sphere{
		t,
	}
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

func (s Sphere) IntersectRay(ray Ray) []float64 {
	D := s.Xf.Inverse().MultV(ray.Dir.AsVector4())
	S := s.Xf.Inverse().MultV(ray.Origin.AsPoint4())
	C := nmath.NewPoint4(0, 0, 0)
	r := 1.0

	// this does some extra calculations that are not necessary for unit spheres with ray transformations
	// but it is a useful refrence if i need real ray sphere intersection

	a := D.Dot(D)
	b := 2.0*S.Dot(D) - 2.0*D.Dot(C)
	c := S.Dot(S) - 2.0*S.Dot(C) + C.Dot(C) - r*r

	sol := nmath.Solve(a, b, c)

	if len(sol) == 0 {
		return []float64{}
	}

	if len(sol) == 1 {
		return []float64{sol[0]}
	}

	t1 := math.Min(sol[0], sol[1])
	t2 := math.Max(sol[0], sol[1])
	return []float64{t1, t2}

}
