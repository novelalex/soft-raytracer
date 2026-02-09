package nmath

import "log"

// Mat4 is a 4D row major Matrix indexed like so:
// ```
// 0  1  2  3
// 4  5  6  7
// 8  9  10 11
// 12 13 14 15
// ```
type Mat4 [16]float64

func Mat4Identity() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func NewMat4() Mat4 {
	return Mat4Identity()
}

func (m Mat4) ApproxEq(n Mat4) bool {
	for i, v := range m {
		if !ApproxEq(v, n[i]) {
			return false
		}
	}

	return true
}

func (m Mat4) LooseEq(n Mat4) bool {
	for i, v := range m {
		if !LooseEq(v, n[i]) {
			return false
		}
	}

	return true
}

func (m *Mat4) At(x, y uint) float64 {
	return m[y*4+x]
}

func (a *Mat4) Mult(b Mat4) Mat4 {
	return Mat4{
		// row 0
		a[0]*b[0] + a[1]*b[4] + a[2]*b[8] + a[3]*b[12],
		a[0]*b[1] + a[1]*b[5] + a[2]*b[9] + a[3]*b[13],
		a[0]*b[2] + a[1]*b[6] + a[2]*b[10] + a[3]*b[14],
		a[0]*b[3] + a[1]*b[7] + a[2]*b[11] + a[3]*b[15],

		// row 1
		a[4]*b[0] + a[5]*b[4] + a[6]*b[8] + a[7]*b[12],
		a[4]*b[1] + a[5]*b[5] + a[6]*b[9] + a[7]*b[13],
		a[4]*b[2] + a[5]*b[6] + a[6]*b[10] + a[7]*b[14],
		a[4]*b[3] + a[5]*b[7] + a[6]*b[11] + a[7]*b[15],

		// row 2
		a[8]*b[0] + a[9]*b[4] + a[10]*b[8] + a[11]*b[12],
		a[8]*b[1] + a[9]*b[5] + a[10]*b[9] + a[11]*b[13],
		a[8]*b[2] + a[9]*b[6] + a[10]*b[10] + a[11]*b[14],
		a[8]*b[3] + a[9]*b[7] + a[10]*b[11] + a[11]*b[15],

		// row 3
		a[12]*b[0] + a[13]*b[4] + a[14]*b[8] + a[15]*b[12],
		a[12]*b[1] + a[13]*b[5] + a[14]*b[9] + a[15]*b[13],
		a[12]*b[2] + a[13]*b[6] + a[14]*b[10] + a[15]*b[14],
		a[12]*b[3] + a[13]*b[7] + a[14]*b[11] + a[15]*b[15],
	}
}

func (m *Mat4) MultV(v Vec4) Vec4 {
	return Vec4{
		m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*v.W,
		m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*v.W,
		m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*v.W,
		m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]*v.W,
	}
}

func (m *Mat4) Transpose() Mat4 {
	return Mat4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	}
}

func (m *Mat4) Inverse() Mat4 {
	adj := Mat4{
		// Row 0
		m[5]*m[10]*m[15] - m[5]*m[11]*m[14] - m[9]*m[6]*m[15] + m[9]*m[7]*m[14] + m[13]*m[6]*m[11] - m[13]*m[7]*m[10],
		-m[1]*m[10]*m[15] + m[1]*m[11]*m[14] + m[9]*m[2]*m[15] - m[9]*m[3]*m[14] - m[13]*m[2]*m[11] + m[13]*m[3]*m[10],
		m[1]*m[6]*m[15] - m[1]*m[7]*m[14] - m[5]*m[2]*m[15] + m[5]*m[3]*m[14] + m[13]*m[2]*m[7] - m[13]*m[3]*m[6],
		-m[1]*m[6]*m[11] + m[1]*m[7]*m[10] + m[5]*m[2]*m[11] - m[5]*m[3]*m[10] - m[9]*m[2]*m[7] + m[9]*m[3]*m[6],
		// Row 1
		-m[4]*m[10]*m[15] + m[4]*m[11]*m[14] + m[8]*m[6]*m[15] - m[8]*m[7]*m[14] - m[12]*m[6]*m[11] + m[12]*m[7]*m[10],
		m[0]*m[10]*m[15] - m[0]*m[11]*m[14] - m[8]*m[2]*m[15] + m[8]*m[3]*m[14] + m[12]*m[2]*m[11] - m[12]*m[3]*m[10],
		-m[0]*m[6]*m[15] + m[0]*m[7]*m[14] + m[4]*m[2]*m[15] - m[4]*m[3]*m[14] - m[12]*m[2]*m[7] + m[12]*m[3]*m[6],
		m[0]*m[6]*m[11] - m[0]*m[7]*m[10] - m[4]*m[2]*m[11] + m[4]*m[3]*m[10] + m[8]*m[2]*m[7] - m[8]*m[3]*m[6],
		// Row 2
		m[4]*m[9]*m[15] - m[4]*m[11]*m[13] - m[8]*m[5]*m[15] + m[8]*m[7]*m[13] + m[12]*m[5]*m[11] - m[12]*m[7]*m[9],
		-m[0]*m[9]*m[15] + m[0]*m[11]*m[13] + m[8]*m[1]*m[15] - m[8]*m[3]*m[13] - m[12]*m[1]*m[11] + m[12]*m[3]*m[9],
		m[0]*m[5]*m[15] - m[0]*m[7]*m[13] - m[4]*m[1]*m[15] + m[4]*m[3]*m[13] + m[12]*m[1]*m[7] - m[12]*m[3]*m[5],
		-m[0]*m[5]*m[11] + m[0]*m[7]*m[9] + m[4]*m[1]*m[11] - m[4]*m[3]*m[9] - m[8]*m[1]*m[7] + m[8]*m[3]*m[5],
		// Row 3
		-m[4]*m[9]*m[14] + m[4]*m[10]*m[13] + m[8]*m[5]*m[14] - m[8]*m[6]*m[13] - m[12]*m[5]*m[10] + m[12]*m[6]*m[9],
		m[0]*m[9]*m[14] - m[0]*m[10]*m[13] - m[8]*m[1]*m[14] + m[8]*m[2]*m[13] + m[12]*m[1]*m[10] - m[12]*m[2]*m[9],
		-m[0]*m[5]*m[14] + m[0]*m[6]*m[13] + m[4]*m[1]*m[14] - m[4]*m[2]*m[13] - m[12]*m[1]*m[6] + m[12]*m[2]*m[5],
		m[0]*m[5]*m[10] - m[0]*m[6]*m[9] - m[4]*m[1]*m[10] + m[4]*m[2]*m[9] + m[8]*m[1]*m[6] - m[8]*m[2]*m[5],
	}

	det := m[0]*adj[0] +
		m[1]*adj[4] +
		m[2]*adj[8] +
		m[3]*adj[12]

	if ApproxEq(det, 0.0) {
		log.Panic("Mat4.Inverse tried to invert a non invertable matrix")
	}

	for i := range adj {
		adj[i] = adj[i] / det
	}

	return adj
}
