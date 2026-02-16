package geom

import (
	"math"

	"github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Cube struct {
	Xf nmath.Mat4
}

func DefaultCube() Cube {

	return Cube{
		nmath.Mat4Identity(),
	}
}
func NewCube(t nmath.Mat4) Cube {
	return Cube{
		t,
	}
}

func (c Cube) Transform() nmath.Mat4 {
	return c.Xf
}

func (c *Cube) SetTransform(m nmath.Mat4) {
	c.Xf = m
}

func (c *Cube) Translate(x, y, z float64) *Cube {
	c.SetTransform(c.Xf.Mult(nmath.NewTranslation(x, y, z)))
	return c
}

func (c *Cube) Scale(x, y, z float64) *Cube {
	c.SetTransform(c.Xf.Mult(nmath.NewScaling(x, y, z)))
	return c
}

func (c *Cube) Rotate(angle float64, axis nmath.Vec3) *Cube {
	c.SetTransform(c.Xf.Mult(nmath.NewRotation(angle, axis)))
	return c
}

func (c *Cube) RotateX(angle float64) *Cube {
	c.SetTransform(c.Xf.Mult(nmath.NewRotationX(angle)))
	return c
}

func (c *Cube) RotateY(angle float64) *Cube {
	c.SetTransform(c.Xf.Mult(nmath.NewRotationY(angle)))
	return c
}

func (c *Cube) RotateZ(angle float64) *Cube {
	c.SetTransform(c.Xf.Mult(nmath.NewRotationZ(angle)))
	return c
}

func (c Cube) localNormalAt(p nmath.Vec3) nmath.Vec3 {
	maxc := max(math.Abs(p.X), math.Abs(p.Y), math.Abs(p.Z))
	if nmath.ApproxEq(maxc, math.Abs(p.X)) {
		return nmath.NewVec3(p.X, 0, 0)
	} else if nmath.ApproxEq(maxc, math.Abs(p.Y)) {
		return nmath.NewVec3(0, p.Y, 0)
	}
	return nmath.NewVec3(0, 0, p.Z)

}

func (c Cube) NormalAt(world_point nmath.Vec3) nmath.Vec3 {
	object_point := c.Xf.Inverse().MultV(world_point.AsPoint4())
	object_normal := c.localNormalAt(object_point.DropW())
	world_normal := c.Xf.Inverse().Transpose().MultV(object_normal.AsVector4())
	return world_normal.DropW().Normalize()
}

func axisAlignedBoundingBoxCheckAxis(origin, direction float64) (float64, float64) {
	var tmin float64
	var tmax float64
	tmin_numerator := (-1 - origin)
	tmax_numerator := (1 - origin)

	if math.Abs(direction) >= nmath.F64Epsilon {
		tmin = tmin_numerator / direction
		tmax = tmax_numerator / direction
	} else {
		tmin = tmin_numerator * math.Inf(1)
		tmax = tmax_numerator * math.Inf(1)
	}
	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}

	return tmin, tmax
}

func (c Cube) localIntersectRay(r Ray) []float64 {
	xtmin, xtmax := axisAlignedBoundingBoxCheckAxis(r.Origin.X, r.Dir.X)
	ytmin, ytmax := axisAlignedBoundingBoxCheckAxis(r.Origin.Y, r.Dir.Y)
	ztmin, ztmax := axisAlignedBoundingBoxCheckAxis(r.Origin.Z, r.Dir.Z)

	tmin := max(xtmin, ytmin, ztmin)
	tmax := min(xtmax, ytmax, ztmax)

	if tmin > tmax {
		return []float64{}
	}

	return []float64{tmin, tmax}
}

func (c Cube) IntersectRay(r Ray) []float64 {
	ray := r.Transform(c.Xf.Inverse())
	return c.localIntersectRay(ray)
}
