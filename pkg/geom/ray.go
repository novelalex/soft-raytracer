package geom

import (
	"math"

	"github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Ray struct {
	Origin nmath.Vec3
	Dir    nmath.Vec3
}

func NewRay(origin, dir nmath.Vec3) Ray {
	return Ray{origin, dir}
}

func (r Ray) ApproxEq(other Ray) bool {
	return r.Origin.ApproxEq(other.Origin) && r.Dir.ApproxEq(other.Dir)
}

func (r Ray) At(t float64) nmath.Vec3 {
	return r.Origin.Add(r.Dir.Mult(t))
}

func (r Ray) Transform(m nmath.Mat4) Ray {
	return Ray{
		m.MultV(r.Origin.AsPoint4()).DropW(),
		m.MultV(r.Dir.AsVector4()).DropW(),
	}
}

func (r Ray) Translate(x, y, z float64) Ray {
	return r.Transform(nmath.NewTranslation(x, y, z))
}

func (r Ray) Scale(x, y, z float64) Ray {
	return r.Transform(nmath.NewScaling(x, y, z))
}

func (r Ray) Rotate(angle float64, axis nmath.Vec3) Ray {
	return r.Transform(nmath.NewRotation(angle, axis))
}

func (r Ray) RotateX(angle float64) Ray {
	return r.Transform(nmath.NewRotationX(angle))
}

func (r Ray) RotateY(angle float64) Ray {
	return r.Transform(nmath.NewRotationY(angle))
}

func (r Ray) RotateZ(angle float64) Ray {
	return r.Transform(nmath.NewRotationZ(angle))
}

func (ray Ray) IntersectSphere(s Sphere) Intersections {
	D := s.transform.Inverse().MultV(ray.Dir.AsVector4())
	S := s.transform.Inverse().MultV(ray.Origin.AsPoint4())
	C := nmath.NewPoint4(0, 0, 0)
	r := 1.0

	// this does some extra calculations that are not necessary for unit spheres with ray transformations
	// but it is a useful refrence if i need real ray sphere intersection

	a := D.Dot(D)
	b := 2.0*S.Dot(D) - 2.0*D.Dot(C)
	c := S.Dot(S) - 2.0*S.Dot(C) + C.Dot(C) - r*r

	sol := nmath.Solve(a, b, c)

	if len(sol) == 0 {
		return NewIntersections([]Intersection{})
	}

	if len(sol) == 1 {
		return NewIntersections([]Intersection{
			NewIntersection(sol[0], s),
		})
	}

	t1 := math.Min(sol[0], sol[1])
	t2 := math.Min(sol[0], sol[1])
	return NewIntersections([]Intersection{
		NewIntersection(t1, s),
		NewIntersection(t2, s),
	})

}
