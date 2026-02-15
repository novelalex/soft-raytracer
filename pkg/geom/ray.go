package geom

import (
	"github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Ray struct {
	Origin nmath.Vec3
	Dir    nmath.Vec3
}

func NewRay(origin, dir nmath.Vec3) Ray {
	return Ray{origin, dir}
}

func (r Ray) ApproxEq(other Ray) bool {
	return r.Origin.ApproxEq(other.Origin) && r.Dir.ApproxEq(other.Dir)
}

func (r Ray) At(t float64) nmath.Vec3 {
	return r.Origin.Add(r.Dir.Mult(t))
}

func (r Ray) Transform(m nmath.Mat4) Ray {
	return Ray{
		m.MultV(r.Origin.AsPoint4()).DropW(),
		m.MultV(r.Dir.AsVector4()).DropW(),
	}
}

func (r Ray) Translate(x, y, z float64) Ray {
	return r.Transform(nmath.NewTranslation(x, y, z))
}

func (r Ray) Scale(x, y, z float64) Ray {
	return r.Transform(nmath.NewScaling(x, y, z))
}

func (r Ray) Rotate(angle float64, axis nmath.Vec3) Ray {
	return r.Transform(nmath.NewRotation(angle, axis))
}

func (r Ray) RotateX(angle float64) Ray {
	return r.Transform(nmath.NewRotationX(angle))
}

func (r Ray) RotateY(angle float64) Ray {
	return r.Transform(nmath.NewRotationY(angle))
}

func (r Ray) RotateZ(angle float64) Ray {
	return r.Transform(nmath.NewRotationZ(angle))
}
