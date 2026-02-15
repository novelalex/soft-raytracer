package raytracer

import (
	"github.com/novelalex/soft-raytracer/pkg/geom"
	"github.com/novelalex/soft-raytracer/pkg/nmath"
)

type World struct {
	Lights  []PointLight
	Objects []Object
}

func NewWorld() World {
	m1 := DefaultMaterial()
	m1.Color = nmath.NewColor(0.8, 1.0, 0.6)
	m1.Diffuse = 0.7
	m1.Specular = 0.2
	s1 := geom.NewSphere(nmath.Mat4Identity())

	t2 := nmath.NewScaling(0.5, 0.5, 0.5)
	s2 := geom.NewSphere(t2)

	o1 := NewObject(&s1, m1)
	o2 := NewObject(&s2, DefaultMaterial())

	return World{
		[]PointLight{
			NewPointLight(nmath.NewVec3(-10, 10, -10), nmath.NewColor(1, 1, 1)),
		},
		[]Object{
			o1, o2,
		},
	}
}

func NewWorldWith(lights []PointLight, objects []Object) World {
	w := World{
		lights,
		objects,
	}

	return w
}

func (w *World) IntersectRay(r geom.Ray) Intersections {
	intersections := Intersections{}
	for _, o := range w.Objects {
		intersections = intersections.Merge(o.IntersectRay(r))
	}
	return intersections
}

func (w *World) ShadeHit(comps IntersectionPrecomputation, remaining uint32) nmath.Color {
	out_color := nmath.NewColor(0, 0, 0)
	for _, light := range w.Lights {
		in_shadow := w.IsShadowed(comps.OverPoint, light)
		object := comps.Object
		shape := comps.Shape
		l_color := object.Material.Lighting(shape, light, comps.Point, comps.EyeV, comps.NormalV, in_shadow)
		reflect_color := w.ReflectedColor(comps, remaining)
		out_color = out_color.Add(l_color).Add(reflect_color)
	}

	return out_color
}

func (w *World) ReflectedColor(comps IntersectionPrecomputation, remaining uint32) nmath.Color {
	out_color := nmath.NewColor(0, 0, 0)
	if remaining <= 0 {
		return out_color // BLACK
	}
	object := comps.Object
	if nmath.ApproxEq(object.Material.Reflective, 0.0) {
		return out_color // BLACK
	}
	reflect_ray := geom.NewRay(comps.OverPoint, comps.ReflectV)
	out_color = w.ColorAt(reflect_ray, remaining-1)

	return out_color.AsVec3().Mult(object.Material.Reflective).AsColor()
}

func (w *World) ColorAt(r geom.Ray, remaining uint32) nmath.Color {
	xs := w.IntersectRay(r)
	hit, ok := xs.Hit()
	if !ok {
		return nmath.NewColor(0, 0, 0)
	}

	precomp := hit.Precompute(r)
	color := w.ShadeHit(precomp, remaining)

	return color
}

func (w *World) IsShadowed(p nmath.Vec3, l PointLight) bool {
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
