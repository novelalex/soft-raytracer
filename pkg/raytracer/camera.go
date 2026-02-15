package raytracer

import (
	"github.com/novelalex/soft-raytracer/pkg/geom"
	"github.com/novelalex/soft-raytracer/pkg/gfx"
	"github.com/novelalex/soft-raytracer/pkg/nmath"
)

type Camera struct {
	Width      uint
	Height     uint
	FOV        float64
	Transform  nmath.Mat4
	HalfWidth  float64
	HalfHeight float64
	PixelSize  float64
}

func NewCamera(w, h uint, fov float64) Camera {
	c := Camera{
		w, h, fov, nmath.Mat4Identity(), 0, 0, 0,
	}
	c.ComputePixelSize()
	return c
}

func (c *Camera) ComputePixelSize() {
	half_view := c.FOV / 2.0
	aspect := float64(c.Width) / float64(c.Height)
	if aspect >= 1 {
		c.HalfWidth = half_view
		c.HalfHeight = half_view / aspect
	} else {
		c.HalfWidth = half_view * aspect
		c.HalfHeight = half_view
	}
	c.PixelSize = (c.HalfWidth * 2.0) / float64(c.Width)
}

func (c Camera) RayForPixel(px, py uint) geom.Ray {
	x_offset := (float64(px) + 0.5) * c.PixelSize
	y_offset := (float64(py) + 0.5) * c.PixelSize

	world_x := c.HalfWidth - x_offset
	world_y := c.HalfHeight - y_offset

	pixel := c.Transform.Inverse().MultV(nmath.NewPoint4(world_x, world_y, -1)).DropW()
	origin := c.Transform.Inverse().MultV(nmath.NewPoint4(0, 0, 0)).DropW()
	direction := pixel.Sub(origin).Normalize()

	return geom.NewRay(origin, direction)
}

func (c *Camera) Render(w World) gfx.Canvas {
	image := gfx.NewCanvas(c.Width, c.Height)
	for y := range c.Height - 1 {
		for x := range c.Width - 1 {
			ray := c.RayForPixel(x, y)
			color := w.ColorAt(ray, 5)
			image.WritePixel(x, y, color)
		}
	}

	return image
}
