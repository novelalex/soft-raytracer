package nmath

import (
	"log"
)

type Color struct {
	R, G, B, A float64
}

func NewColor(r, g, b float64) Color {
	return Color{r, g, b, 1}
}

func (c Color) At(i int) float64 {
	switch i {
	case 0:
		return c.R
	case 1:
		return c.G
	case 2:
		return c.B
	case 3:
		return c.A
	default:
		log.Panic("Color.At was given an invalid index")
	}

	return 0
}

func (c Color) AsVec3() Vec3 {
	return Vec3{X: c.R, Y: c.G, Z: c.B}
}

func (c Color) AsVec4() Vec4 {
	return Vec4{X: c.R, Y: c.G, Z: c.B, W: c.A}
}

func (c Color) Add(other Color) Color {
	return Color{
		c.R + other.R,
		c.G + other.G,
		c.B + other.B,
		c.A + other.A,
	}
}

func (c Color) HadamardMult(other Color) Color {
	return Color{
		c.R * other.R,
		c.G * other.G,
		c.B * other.B,
		c.A * other.A,
	}
}
