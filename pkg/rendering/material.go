package rendering

import (
	"math"

	. "github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Material struct {
	Color                                 Color
	Ambient, Diffuse, Specular, Shininess float64
}

func NewMaterial(color Color, ambient, diffuse, specular, shininess float64) Material {
	return Material{
		color,
		ambient,
		diffuse,
		specular,
		shininess,
	}
}

func DefaultMaterial() Material {
	return Material{
		NewColor(1, 1, 1),
		0.1,
		0.9,
		0.9,
		200.0,
	}
}

func (m Material) Lighting(l PointLight, p, eye, normal Vec3) Color {
	ambient := NewColor(0, 0, 0)
	diffuse := NewColor(0, 0, 0)
	specular := NewColor(0, 0, 0)

	effective_color := m.Color.HadamardMult(l.intensity)
	light_v := l.position.Sub(p).Normalize()
	ambient = effective_color.AsVec3().Mult(m.Ambient).AsColor()

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
			specular = l.intensity.AsVec3().Mult(m.Specular * factor).AsColor()
		}
	}

	return ambient.AsVec3().
		Add(diffuse.AsVec3()).
		Add(specular.AsVec3()).
		AsColor()
}
