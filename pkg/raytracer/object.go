package raytracer

import (
	"github.com/novelalex/soft-raytracer/pkg/geom"
)

type Object struct {
	Shape    geom.Shape
	Material Material
}

func NewObject(s geom.Shape, m Material) Object {
	return Object{
		s,
		m,
	}
}

func NewGlassSphere() Object {
	shape := geom.DefaultSphere()
	glass_sphere := Object{
		&shape,
		DefaultMaterial(),
	}
	glass_sphere.Material.Transparency = 1.0
	glass_sphere.Material.IOR = 1.5
	return glass_sphere
}

func (o *Object) IntersectRay(r geom.Ray) Intersections {
	xs := Intersections{}
	geom_intersects := o.Shape.IntersectRay(r)
	for _, intersect := range geom_intersects {
		xs = append(xs, NewIntersection(intersect, o))
	}
	return xs
}
