package main

import (
	"fmt"
	"os"

	"github.com/novelalex/soft-raytracer/pkg/geom"
	"github.com/novelalex/soft-raytracer/pkg/gfx"
	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/renderer"
)

func main() {
	canvas := gfx.NewCanvas(100, 100)
	eye_pos := nmath.NewVec3(0, 0, -5)
	var canvas_pixels float64 = float64(canvas.Width())
	wall_z := 10.0
	wall_size := 7.0
	pixel_size := wall_size / canvas_pixels
	half_wall_size := wall_size / 2.0

	shape := geom.NewSphere(nmath.Mat4Identity(), renderer.DefaultMaterial())
	light := renderer.NewPointLight(nmath.NewVec3(-10, 10, -10), nmath.NewColor(1, 1, 1))

	for y := range canvas.Height() {
		world_y := half_wall_size - pixel_size*float64(y)
		for x := range canvas.Width() {
			world_x := -half_wall_size + pixel_size*float64(x)
			position := nmath.NewVec3(world_x, world_y, wall_z)
			r := geom.NewRay(eye_pos, position.Sub(eye_pos).Normalize())
			xs := r.IntersectSphere(shape)
			hit, ok := xs.Hit()
			if ok {
				point := r.At(hit.T)
				normal := hit.Object.NormalAt(point)
				eye_v := r.Dir.Neg()
				color := hit.Object.Material().Lighting(light, point, eye_v, normal)
				canvas.WritePixel(x, y, color)
			}
		}
	}

	fmt.Fprint(os.Stdout, canvas.AsPPM())
}
