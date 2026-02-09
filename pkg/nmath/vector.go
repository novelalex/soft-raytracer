package nmath

import "math"

type Vec4 struct {
	X, Y, Z, W float64
}

func NewVec4(x, y, z, w float64) Vec4 {
	return Vec4{x, y, z, w}
}

func NewPoint4(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 1}
}
func NewVector4(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 0}
}

func (v Vec4) DropW() Vec3 {
	return Vec3{v.X, v.Y, v.Z}
}

func (lhs Vec4) ApproxEq(rhs Vec4) bool {
	return ApproxEq(lhs.X, rhs.X) &&
		ApproxEq(lhs.Y, rhs.Y) &&
		ApproxEq(lhs.Z, rhs.Z) &&
		ApproxEq(lhs.W, rhs.W)
}

func (lhs Vec4) Add(rhs Vec4) Vec4 {
	return Vec4{
		lhs.X + rhs.X,
		lhs.Y + rhs.Y,
		lhs.Z + rhs.Z,
		lhs.W + rhs.W,
	}
}

func (lhs Vec4) Sub(rhs Vec4) Vec4 {
	return Vec4{
		lhs.X - rhs.X,
		lhs.Y - rhs.Y,
		lhs.Z - rhs.Z,
		lhs.W - rhs.W,
	}
}

func (lhs Vec4) Mult(rhs float64) Vec4 {
	return Vec4{
		lhs.X * rhs,
		lhs.Y * rhs,
		lhs.Z * rhs,
		lhs.W * rhs,
	}
}

func (lhs Vec4) Div(rhs float64) Vec4 {
	return Vec4{
		lhs.X / rhs,
		lhs.Y / rhs,
		lhs.Z / rhs,
		lhs.W / rhs,
	}
}

func (lhs Vec4) Dot(rhs Vec4) float64 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z + lhs.W*rhs.W
}

// 3D cross product, w is dropped
func (lhs Vec4) Cross(rhs Vec4) Vec4 {
	return Vec4{
		lhs.Y*rhs.Z - lhs.Z*rhs.Y,
		lhs.Z*rhs.X - lhs.X*rhs.Z,
		lhs.X*rhs.Y - lhs.Y*rhs.X,
		0,
	}
}

// Magnitude squared
func (v Vec4) Mag2() float64 {
	return v.Dot(v)
}

func (v Vec4) Mag() float64 {
	return math.Sqrt(v.Mag2())
}

func (v Vec4) Normalize() Vec4 {
	mag := v.Mag()
	return Vec4{
		v.X / mag,
		v.Y / mag,
		v.Z / mag,
		v.W / mag,
	}
}

// -----------------------------------------------------------------------
// Vec3
// -----------------------------------------------------------------------

type Vec3 struct {
	X, Y, Z float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func (lhs Vec3) ApproxEq(rhs Vec3) bool {
	return ApproxEq(lhs.X, rhs.X) &&
		ApproxEq(lhs.Y, rhs.Y) &&
		ApproxEq(lhs.Z, rhs.Z)
}

func (v Vec3) AsPoint4() Vec4 {
	return Vec4{v.X, v.Y, v.Z, 1}
}

func (v Vec3) AsVector4() Vec4 {
	return Vec4{v.X, v.Y, v.Z, 0}
}

func (v Vec3) WithW(w float64) Vec4 {
	return Vec4{v.X, v.Y, v.Z, w}
}

func (v Vec3) Neg() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func (lhs Vec3) Add(rhs Vec3) Vec3 {
	return Vec3{lhs.X + rhs.X, lhs.Y + rhs.Y, lhs.Z + rhs.Z}
}

func (lhs Vec3) Sub(rhs Vec3) Vec3 {
	return Vec3{lhs.X - rhs.X, lhs.Y - rhs.Y, lhs.Z - rhs.Z}
}

func (lhs Vec3) Mult(rhs float64) Vec3 {
	return Vec3{
		lhs.X * rhs,
		lhs.Y * rhs,
		lhs.Z * rhs,
	}
}

func (lhs Vec3) Div(rhs float64) Vec3 {
	return Vec3{
		lhs.X / rhs,
		lhs.Y / rhs,
		lhs.Z / rhs,
	}
}

func (lhs Vec3) Dot(rhs Vec3) float64 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}

func (lhs Vec3) Cross(rhs Vec3) Vec3 {
	return Vec3{
		lhs.Y*rhs.Z - lhs.Z*rhs.Y,
		lhs.Z*rhs.X - lhs.X*rhs.Z,
		lhs.X*rhs.Y - lhs.Y*rhs.X,
	}
}

func (in Vec3) Reflect(normal Vec3) Vec3 {
	return in.Sub(normal.Mult(2).Mult(in.Dot(normal)))
}

func (v Vec3) Mag2() float64 {
	return v.Dot(v)
}

func (v Vec3) Mag() float64 {
	return math.Sqrt(v.Mag2())
}

func (v Vec3) Normalize() Vec3 {
	m := v.Mag()
	if m == 0 {
		return v
	}
	return Vec3{v.X / m, v.Y / m, v.Z / m}
}

func (v Vec3) AsColor() Color {
	return Color{v.X, v.Y, v.Z, 1}
}

func (from Vec3) LookAt(to, up Vec3) Mat4 {
	forward := to.Sub(from).Normalize()
	left := forward.Cross(up.Normalize())
	true_up := left.Cross(forward)
	orientation := Mat4{
		left.X, left.Y, left.Z, 0,
		true_up.X, true_up.Y, true_up.Z, 0,
		-forward.X, -forward.Y, -forward.Z, 0,
		0, 0, 0, 1,
	}

	return orientation.Translate(-from.X, -from.Y, -from.Z)
}
