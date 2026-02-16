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
	T          float64
	Object     *Object
	Shape      geom.Shape
	Point      nmath.Vec3
	EyeV       nmath.Vec3
	NormalV    nmath.Vec3
	ReflectV   nmath.Vec3
	OverPoint  nmath.Vec3
	UnderPoint nmath.Vec3
	N1         float64
	N2         float64
	Inside     bool
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

func (x Intersection) IsHit(xs Intersections) bool {
	hit, ok := xs.Hit()
	if !ok {
		return false
	}

	if hit.ApproxEq(x) {
		return true
	}

	return false
}

func (x Intersection) ApproxEq(y Intersection) bool {
	return nmath.ApproxEq(x.T, y.T) && x.Object == y.Object
}

func objectSliceContainsObject(objs []*Object, obj *Object) bool {
	for _, o := range objs {
		if o == obj {
			return true
		}
	}
	return false
}

func objectSliceFindObjectIndex(objs []*Object, obj *Object) int {
	for i, o := range objs {
		if o == obj {
			return i
		}
	}
	return -1
}

func (x Intersection) Precompute(r geom.Ray, xs Intersections) IntersectionPrecomputation {
	point := r.At(x.T)
	normal := x.Object.Shape.NormalAt(point)
	eye := r.Dir.Neg()
	reflect := r.Dir.Reflect(normal)

	n1 := 0.0
	n2 := 0.0

	containers := []*Object{}
	for _, intersect := range xs {
		if intersect.IsHit(xs) {
			if len(containers) == 0 {
				n1 = 1.0
			} else {
				n1 = containers[len(containers)-1].Material.IOR
			}
		}

		if objectSliceContainsObject(containers, intersect.Object) {
			idx := objectSliceFindObjectIndex(containers, intersect.Object)
			containers = append(containers[:idx], containers[idx+1:]...)
		} else {
			containers = append(containers, intersect.Object)
		}
		if intersect.IsHit(xs) {
			if len(containers) == 0 {
				n2 = 1.0
			} else {
				n2 = containers[len(containers)-1].Material.IOR
			}
			break
		}
	}

	var inside bool
	if normal.Dot(eye) < 0 {
		inside = true
		normal = normal.Neg()
	} else {
		inside = false
	}
	over_point := point.Add(normal.Mult(nmath.F64Epsilon))
	under_point := point.Sub(normal.Mult(nmath.F64Epsilon))
	return IntersectionPrecomputation{
		T:          x.T,
		Object:     x.Object,
		Shape:      x.Object.Shape,
		Point:      point,
		EyeV:       eye,
		NormalV:    normal,
		ReflectV:   reflect,
		OverPoint:  over_point,
		UnderPoint: under_point,
		N1:         n1,
		N2:         n2,
		Inside:     inside,
	}
}
