package geom

import (
	"sort"

	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/rendering"
)

type Shape interface {
	Transform() nmath.Mat4
	SetTransform(nmath.Mat4)
	Material() rendering.Material
	SetMaterial(rendering.Material)
	IntersectRay(Ray) Intersections
	NormalAt(world_point nmath.Vec3) nmath.Vec3
}

type Intersection struct {
	T      float64
	Object Shape
}

type IntersectionPrecomputation struct {
	T       float64
	Object  Shape
	Point   nmath.Vec3
	EyeV    nmath.Vec3
	NormalV nmath.Vec3
	Inside  bool
}

func NewIntersection(t float64, object Shape) Intersection {
	return Intersection{t, object}
}

type Intersections []Intersection

func NewIntersections(xs []Intersection) Intersections {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})
	return xs
}

func (xs Intersections) Append(x Intersection) Intersections {
	xs = append(xs, x)
	sort.Slice(xs, func(i, j int) bool {
		return (xs)[i].T < (xs)[j].T
	})
	return xs
}

func (xs Intersections) Merge(ys Intersections) Intersections {
	for i := range ys {
		xs = append(xs, ys[i])
	}
	sort.Slice(xs, func(i, j int) bool {
		return (xs)[i].T < (xs)[j].T
	})

	return xs
}

func (xs Intersections) Hit() (x Intersection, ok bool) {
	for i := range len(xs) {
		if xs[i].T >= 0 {
			return xs[i], true
		}
	}

	return NewIntersection(0, &Sphere{}), false
}

func (x Intersection) Precompute(r Ray) IntersectionPrecomputation {
	point := r.At(x.T)
	normal := x.Object.NormalAt(point)
	eye := r.Dir.Neg()
	var inside bool
	if normal.Dot(eye) < 0 {
		inside = true
		normal = normal.Neg()
	} else {
		inside = false
	}

	return IntersectionPrecomputation{
		T:       x.T,
		Object:  x.Object,
		Point:   point,
		EyeV:    eye,
		NormalV: normal,
		Inside:  inside,
	}
}
