package world

import (
	"github.com/novelalex/soft-raytracer/pkg/geom"
	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/rendering"
)

type World struct {
	Lights      []rendering.PointLight
	Objects     []Object
	shapeLookup map[uint64]*Object
}

func NewWorld() World {
	m1 := rendering.DefaultMaterial()
	m1.Color = nmath.NewColor(0.8, 1.0, 0.6)
	m1.Diffuse = 0.7
	m1.Specular = 0.2
	s1 := geom.NewSphere(nmath.Mat4Identity())

	t2 := nmath.NewScaling(0.5, 0.5, 0.5)
	s2 := geom.NewSphere(t2)

	o1 := NewObject(&s1, m1)
	o2 := NewObject(&s2, rendering.DefaultMaterial())

	return World{
		[]rendering.PointLight{
			rendering.NewPointLight(nmath.NewVec3(-10, 10, -10), nmath.NewColor(1, 1, 1)),
		},
		[]Object{
			o1, o2,
		},
		map[uint64]*Object{},
	}
}

func NewWorldWith(lights []rendering.PointLight, objects []Object) World {
	w := World{
		lights,
		objects,
		map[uint64]*Object{},
	}
	w.CreateShapeLookup()

	return w
}

func (w *World) CreateShapeLookup() {
	for _, o := range w.Objects {
		w.shapeLookup[uint64(o.Shape.ID())] = &o
	}
}

func (w *World) LookupShape(s geom.Shape) *Object {
	return w.shapeLookup[s.ID()]
}

func (w *World) IntersectRay(r geom.Ray) geom.Intersections {
	intersections := geom.Intersections{}
	for _, o := range w.Objects {
		intersections = intersections.Merge(o.Shape.IntersectRay(r))
	}
	return intersections
}

func (w *World) ShadeHit(comps geom.IntersectionPrecomputation) nmath.Color {
	out_color := nmath.NewColor(0, 0, 0)
	for _, light := range w.Lights {
		in_shadow := w.IsShadowed(comps.OverPoint, light)
		shape := comps.Object
		object := w.LookupShape(shape)
		l_color := object.Material.Lighting(shape, light, comps.Point, comps.EyeV, comps.NormalV, in_shadow)
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

func (w *World) IsShadowed(p nmath.Vec3, l rendering.PointLight) bool {
	v := l.Position.Sub(p)
	dist := v.Mag()
	dir := v.Normalize()

	r := geom.NewRay(p, dir)
	xs := w.IntersectRay(r)

	h, ok := xs.Hit()
	if ok && h.T < dist {
		return true
	}

	return false
}
