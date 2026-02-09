package nmath

import (
	"math"
)

type Quaternion struct {
	W, I, J, K float64
}

func NewQuaternion(w float64, ijk Vec3) Quaternion {
	return Quaternion{w, ijk.X, ijk.Y, ijk.Z}
}

func (lhs Quaternion) Mult(rhs Quaternion) Quaternion {
	return Quaternion{
		W: (lhs.W * rhs.W) -
			(lhs.I * rhs.I) +
			(lhs.J * rhs.J) +
			(lhs.K * rhs.K),

		I: (lhs.W * rhs.I) +
			(lhs.I * rhs.W) -
			(lhs.K * rhs.J) +
			(lhs.J * rhs.K),

		J: (lhs.W * rhs.J) +
			(lhs.J * rhs.W) -
			(lhs.I * rhs.K) +
			(lhs.K * rhs.I),

		K: (lhs.W * rhs.K) +
			(lhs.K * rhs.W) -
			(lhs.J * rhs.I) +
			(lhs.I * rhs.J),
	}
}

func (lhs Quaternion) MultS(rhs float64) Quaternion {
	return Quaternion{
		lhs.W * rhs,
		lhs.I * rhs,
		lhs.J * rhs,
		lhs.K * rhs,
	}
}

func (lhs Quaternion) DivS(rhs float64) Quaternion {
	return Quaternion{
		lhs.W / rhs,
		lhs.I / rhs,
		lhs.J / rhs,
		lhs.K / rhs,
	}
}

func (q Quaternion) AsVec4() Vec4 {
	return Vec4{
		q.W,
		q.I,
		q.J,
		q.K,
	}
}

func (q Quaternion) Mag() float64 {
	return q.AsVec4().Mag()
}

func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{
		q.W,
		-q.I,
		-q.J,
		-q.K,
	}
}

func (q Quaternion) Inverse() Quaternion {
	mag := q.Mag()
	return q.Conjugate().DivS(mag * mag)
}

func (q Quaternion) Normalize() Quaternion {
	return q.DivS(q.Mag())
}

func AngleAxisRotation(theta float64, axis Vec3) Quaternion {
	rotation_axis := axis.Normalize()
	cos_val := math.Cos(theta / 2.0)
	sin_val := math.Sin(theta / 2.0)
	result := NewQuaternion(cos_val, rotation_axis.Mult(sin_val))
	return result.Normalize()
}

func (q Quaternion) AsMat4() Mat4 {
	return Mat4{
		1.0 - 2.0*q.J*q.J - 2.0*q.K*q.K,
		2.0*q.I*q.J - 2.0*q.K*q.W,
		2.0*q.I*q.K + 2.0*q.J*q.W,
		0.0,

		2.0*q.I*q.J + 2.0*q.K*q.W,
		1.0 - 2.0*q.I*q.I - 2.0*q.K*q.K,
		2.0*q.J*q.K - 2.0*q.I*q.W,
		0.0,

		2.0*q.I*q.K - 2.0*q.J*q.W,
		2.0*q.J*q.K + 2.0*q.I*q.W,
		1.0 - 2.0*q.I*q.I - 2.0*q.J*q.J,
		0.0,

		0.0, 0.0, 0.0, 1.0,
	}
}
