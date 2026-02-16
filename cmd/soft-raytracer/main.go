package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/novelalex/soft-raytracer/pkg/geom"
	"github.com/novelalex/soft-raytracer/pkg/nmath"
	. "github.com/novelalex/soft-raytracer/pkg/raytracer"
)

func main() {

	floor_shape := geom.DefaultPlane()
	floor_shape.Xf = nmath.NewRotationY(45 * math.Pi / 180.0)
	floor := NewObject(&floor_shape, DefaultMaterial())
	floor_pattern := geom.NewCheckerPattern(nmath.NewColor(0, 0, 0), nmath.NewColor(1, 1, 1))
	floor_pattern.Xf = nmath.NewTranslation(0, -0.001, 0) // pattern had arifacts due to rounding at y=0
	floor.Material.Pattern = &floor_pattern
	floor.Material.Reflective = 0.1

	ceiling_shape := geom.DefaultPlane()
	ceiling_shape.Xf = nmath.NewTranslation(0, 3, 0)
	ceiling := NewObject(&ceiling_shape, DefaultMaterial())

	wall_shape := geom.DefaultPlane()
	wall_shape.SetTransform(wall_shape.Xf.RotateY(45*math.Pi/180.0).Translate(0, 0, 2).RotateX(90 * math.Pi / 180.0))
	wall := NewObject(&wall_shape, DefaultMaterial())
	//wall_pattern := geom.NewRingPattern(nmath.NewColor(0, 0.4, 0), nmath.NewColor(1, 1, 1))
	wall.Material.Color = nmath.NewColor(0.5, 0, 0.5)
	//wall.Material.Pattern = &wall_pattern
	wall.Material.Specular = 0.4
	wall.Material.Shininess = 4

	middle_shape := geom.DefaultSphere()
	middle_shape.Translate(-1.5, 1, 0.5)
	middle := NewObject(&middle_shape, DefaultMaterial())
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3
	middle.Material.Reflective = 0

	right_shape := geom.DefaultSphere()
	right_shape.Translate(1.4, 0.5, -0.4).
		Scale(0.5, 0.5, 0.5)
	right := NewObject(&right_shape, DefaultMaterial())
	right.Material.Color = nmath.NewColor(0.5, 0.5, 0.1)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.9
	right.Material.Reflective = 1

	left_shape := geom.DefaultSphere()
	left_shape.Translate(-0.5, 1, -2.5)
	left := NewObject(&left_shape, DefaultMaterial())
	left.Material.Color = nmath.NewColor(0, 0, 0)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3
	left.Material.Transparency = 1.0
	left.Material.IOR = 1.5

	light := NewPointLight(nmath.NewVec3(-10, 2, -10), nmath.NewColor(1, 1, 1))

	w := NewWorldWith(
		[]PointLight{light},
		[]Object{
			floor, ceiling, wall, right, middle, right, left,
		},
	)

	c := NewCamera(1920, 1080, math.Pi/3.0)
	c.Transform = nmath.NewVec3(0, 1.5, -7).
		LookAt(
			nmath.NewVec3(0, 1, 0),
			nmath.NewVec3(0, 1, 0),
		)

	start_time := time.Now()

	canvas := c.Render(w)

	elapsed_time := time.Since(start_time)
	pixel_count := c.Width * c.Height

	fmt.Println("Rendered", pixel_count, "pixels in", elapsed_time)

	//fmt.Fprint(os.Stdout, canvas.AsPPM())
	err := os.WriteFile("img.ppm", []byte(canvas.AsPPM()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
