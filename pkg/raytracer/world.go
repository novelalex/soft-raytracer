package raytracer

import (
	"math"

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

func (w *World) ShadeHit(comps IntersectionPrecomputation, remaining int) nmath.Color {
	out_color := nmath.NewColor(0, 0, 0)
	for _, light := range w.Lights {
		in_shadow, shadow_strength := w.IsShadowed(comps.OverPoint, light)
		object := comps.Object
		shape := comps.Shape
		l_color := object.Material.Lighting(shape, light, comps.Point, comps.EyeV, comps.NormalV, in_shadow, shadow_strength)
		reflect_color := w.ReflectedColor(comps, remaining)
		refract_color := w.RefractedColor(comps, remaining)

		material := object.Material
		if material.Reflective > 0 && material.Transparency > 0 {
			reflectance := Schlick(comps)
			out_color = out_color.Add(l_color).Add(reflect_color.AsVec3().Mult(reflectance).AsColor()).Add(refract_color.AsVec3().Mult(1 - reflectance).AsColor())
		} else {
			out_color = out_color.Add(l_color).Add(reflect_color).Add(refract_color)
		}

	}

	return out_color
}

func (w *World) ReflectedColor(comps IntersectionPrecomputation, remaining int) nmath.Color {
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

func (w *World) RefractedColor(comps IntersectionPrecomputation, remaining int) nmath.Color {
	out_color := nmath.NewColor(0, 0, 0)
	if remaining <= 0 {
		return out_color // BLACK
	}
	object := comps.Object

	if nmath.ApproxEq(object.Material.Transparency, 0.0) {
		return out_color // BLACK
	}

	n_ratio := comps.N1 / comps.N2
	cos_i := comps.EyeV.Dot(comps.NormalV)
	sin2_t := n_ratio * n_ratio * (1 - cos_i*cos_i)
	if sin2_t > 1.0 {
		return out_color // BLACK
	}

	cos_t := math.Sqrt(1.0 - sin2_t)
	direction := comps.NormalV.Mult(n_ratio*cos_i - cos_t).Sub(comps.EyeV.Mult(n_ratio))
	refract_ray := geom.NewRay(comps.UnderPoint, direction)
	out_color = w.ColorAt(refract_ray, remaining-1).AsVec3().Mult(comps.Object.Material.Transparency).AsColor()

	return out_color
}

func (w *World) ColorAt(r geom.Ray, remaining int) nmath.Color {
	xs := w.IntersectRay(r)
	hit, ok := xs.Hit()
	if !ok {
		return nmath.NewColor(0, 0, 0)
	}

	precomp := hit.Precompute(r, xs)
	color := w.ShadeHit(precomp, remaining)

	return color
}

func (w *World) IsShadowed(p nmath.Vec3, l PointLight) (bool, float64) {
	v := l.Position.Sub(p)
	dist := v.Mag()
	dir := v.Normalize()

	r := geom.NewRay(p, dir)
	xs := w.IntersectRay(r)

	transmission := 1.0

	for _, x := range xs {
		if x.T >= 0 && x.T < dist {
			// skip fully transparent objects
			if x.Object.Material.Transparency >= 1.0-nmath.F64Epsilon {
				continue
			}
			transmission *= x.Object.Material.Transparency

			// bail if fully opaque
			if transmission <= nmath.F64EpsilonLoose {
				return true, 1.0
			}
		}
	}

	shadow_strength := 1.0 - transmission

	if shadow_strength > nmath.F64EpsilonLoose {
		return true, shadow_strength
	}

	return false, 0.0
}

func Schlick(c IntersectionPrecomputation) float64 {
	cos := c.EyeV.Dot(c.NormalV)
	if c.N1 > c.N2 {
		n := c.N1 / c.N2
		sin2_t := n * n * (1.0 - cos*cos)
		if sin2_t > 1.0 {
			return 1.0
		}
		cos_t := math.Sqrt(1.0 - sin2_t)
		cos = cos_t
	}

	r0 := ((c.N1 - c.N2) / (c.N1 + c.N2)) * ((c.N1 - c.N2) / (c.N1 + c.N2))
	cos_sub_1 := 1 - cos
	return r0 + (1+r0)*cos_sub_1*cos_sub_1*cos_sub_1*cos_sub_1*cos_sub_1
}
