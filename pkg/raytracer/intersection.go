package raytracer

import (
	"sort"

	"github.com/novelalex/soft-raytracer/pkg/geom"
	"github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Intersection struct {
	T      float64
	Object *Object
}

type IntersectionPrecomputation struct {
	T         float64
	Object    *Object
	Shape     geom.Shape
	Point     nmath.Vec3
	EyeV      nmath.Vec3
	NormalV   nmath.Vec3
	ReflectV  nmath.Vec3
	OverPoint nmath.Vec3
	Inside    bool
}

func NewIntersection(t float64, object *Object) Intersection {
	return Intersection{t, object}
}

type Intersections []Intersection

func NewIntersections(xs []Intersection) Intersections {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})
	return xs
}

// TODO: Something is wrong with this function, use append(xs, x) for now
// func (xs Intersections) Append(x Intersection) Intersections {
// 	xs = append(xs, x)
// 	sort.Slice(xs, func(i, j int) bool {
// 		return (xs)[i].T < (xs)[j].T
// 	})
// 	return xs
// }

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

	return NewIntersection(0, &Object{}), false
}

func (x Intersection) Precompute(r geom.Ray) IntersectionPrecomputation {
	point := r.At(x.T)
	normal := x.Object.Shape.NormalAt(point)
	eye := r.Dir.Neg()
	reflect := r.Dir.Reflect(normal)
	var inside bool
	if normal.Dot(eye) < 0 {
		inside = true
		normal = normal.Neg()
	} else {
		inside = false
	}
	over_point := point.Add(normal.Mult(nmath.F64Epsilon))
	return IntersectionPrecomputation{
		T:         x.T,
		Object:    x.Object,
		Shape:     x.Object.Shape,
		Point:     point,
		EyeV:      eye,
		NormalV:   normal,
		ReflectV:  reflect,
		OverPoint: over_point,
		Inside:    inside,
	}
}
