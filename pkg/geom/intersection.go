package geom

import (
	"sort"

	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/renderer"
)

type Shape interface {
	Transform() nmath.Mat4
	SetTransform(nmath.Mat4)
	Material() renderer.Material
	SetMaterial(renderer.Material)
	NormalAt(world_point nmath.Vec3) nmath.Vec3
}

type Intersection struct {
	T      float64
	Object Shape
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

func (xs *Intersections) Append(x Intersection) {
	*xs = append(*xs, x)
	sort.Slice(*xs, func(i, j int) bool {
		return (*xs)[i].T < (*xs)[j].T
	})
}

func (xs Intersections) Hit() (x Intersection, ok bool) {
	for i := range len(xs) {
		if xs[i].T >= 0 {
			return xs[i], true
		}
	}

	return NewIntersection(0, &Sphere{}), false
}
