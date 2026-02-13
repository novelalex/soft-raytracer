package rendering

import (
	"math"

	"github.com/novelalex/soft-raytracer/pkg/geom"
	. "github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Material struct {
	Color                                 Color
	Ambient, Diffuse, Specular, Shininess float64
	Pattern                               geom.Pattern
}

func NewMaterial(color Color, ambient, diffuse, specular, shininess float64) Material {
	return Material{
		color,
		ambient,
		diffuse,
		specular,
		shininess,
		nil,
	}
}

func DefaultMaterial() Material {
	return Material{
		NewColor(1, 1, 1),
		0.1,
		0.9,
		0.9,
		200.0,
		nil,
	}
}

func (m Material) Lighting(l PointLight, p, eye, normal Vec3, in_shadow bool) Color {
	ambient := NewColor(0, 0, 0)
	diffuse := NewColor(0, 0, 0)
	specular := NewColor(0, 0, 0)

	var color Color
	if m.Pattern != nil {
		color = m.Pattern.At(p)
	} else {
		color = m.Color
	}

	effective_color := color.HadamardMult(l.Intensity)
	
	light_v := l.Position.Sub(p).Normalize()
	ambient = effective_color.AsVec3().Mult(m.Ambient).AsColor()

	if in_shadow {
		return ambient
	}

	light_dot_normal := light_v.Dot(normal)
	if light_dot_normal < 0 {
		diffuse = NewColor(0, 0, 0)
		specular = NewColor(0, 0, 0)
	} else {
		diffuse = effective_color.AsVec3().Mult(m.Diffuse * light_dot_normal).AsColor()
		reflect_v := light_v.Neg().Reflect(normal)
		reflect_dot_eye := reflect_v.Dot(eye)
		if reflect_dot_eye <= 0 {
			specular = NewColor(0, 0, 0)
		} else {
			factor := math.Pow(reflect_dot_eye, m.Shininess)
			specular = l.Intensity.AsVec3().Mult(m.Specular * factor).AsColor()
		}
	}

	return ambient.AsVec3().
		Add(diffuse.AsVec3()).
		Add(specular.AsVec3()).
		AsColor()
}
