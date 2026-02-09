package world

import (
	"github.com/novelalex/soft-raytracer/pkg/geom"
	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/rendering"
)

type World struct {
	Lights  []rendering.PointLight
	Objects []geom.Shape
}

func NewWorld() World {
	m1 := rendering.DefaultMaterial()
	m1.Color = nmath.NewColor(0.8, 1.0, 0.6)
	m1.Diffuse = 0.7
	m1.Specular = 0.2
	s1 := geom.NewSphere(nmath.Mat4Identity(), m1)

	t2 := nmath.NewScaling(0.5, 0.5, 0.5)
	s2 := geom.NewSphere(t2, rendering.DefaultMaterial())

	return World{
		[]rendering.PointLight{
			rendering.NewPointLight(nmath.NewVec3(-10, 10, -10), nmath.NewColor(1, 1, 1)),
		},
		[]geom.Shape{
			&s1, &s2,
		},
	}
}

func (w *World) IntersectRay(r geom.Ray) geom.Intersections {
	intersections := geom.Intersections{}
	for _, o := range w.Objects {
		intersections = intersections.Merge(o.IntersectRay(r))
	}
	return intersections
}

func (w *World) ShadeHit(comps geom.IntersectionPrecomputation) nmath.Color {
	out_color := nmath.NewColor(0, 0, 0)
	for _, light := range w.Lights {
		l_color := comps.Object.Material().Lighting(light, comps.Point, comps.EyeV, comps.NormalV)
		out_color = out_color.Add(l_color)
	}

	return out_color
}

func (w *World) ColorAt(r geom.Ray) nmath.Color {
	xs := w.IntersectRay(r)
	hit, ok := xs.Hit()
	if !ok {
		return nmath.NewColor(0, 0, 0)
	}

	precomp := hit.Precompute(r)
	color := w.ShadeHit(precomp)

	return color
}
