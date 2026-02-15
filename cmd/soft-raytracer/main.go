package main

import (
	"fmt"
	"math"
	"os"

	"github.com/novelalex/soft-raytracer/pkg/camera"
	"github.com/novelalex/soft-raytracer/pkg/geom"
	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/rendering"
	"github.com/novelalex/soft-raytracer/pkg/world"
)

func main() {

	floor_shape := geom.DefaultPlane()
	floor_shape.Xf = nmath.NewRotationY(45 * math.Pi / 180.0)
	floor := world.NewObject(&floor_shape, rendering.DefaultMaterial())
	floor_pattern := geom.NewCheckerPattern(nmath.NewColor(0, 0, 0), nmath.NewColor(1, 1, 1))
	floor_pattern.Xf = nmath.NewTranslation(0, -0.001, 0) // pattern had arifacts due to rounding at y=0
	floor.Material.Pattern = &floor_pattern
	floor.Material.Reflective = 0.1

	ceiling_shape := geom.DefaultPlane()
	ceiling_shape.Xf = nmath.NewTranslation(0, 3, 0)
	ceiling := world.NewObject(&ceiling_shape, rendering.DefaultMaterial())

	wall_shape := geom.DefaultPlane()
	wall_shape.SetTransform(wall_shape.Xf.RotateY(45*math.Pi/180.0).Translate(0, 0, 2).RotateX(90 * math.Pi / 180.0))
	wall := world.NewObject(&wall_shape, rendering.DefaultMaterial())
	//wall_pattern := geom.NewRingPattern(nmath.NewColor(0, 0.4, 0), nmath.NewColor(1, 1, 1))
	wall.Material.Color = nmath.NewColor(0.5, 0, 0.5)
	//wall.Material.Pattern = &wall_pattern
	wall.Material.Specular = 0.4
	wall.Material.Shininess = 4

	middle_shape := geom.DefaultSphere()
	middle_shape.Translate(-0.5, 1, 0.5)
	middle := world.NewObject(&middle_shape, rendering.DefaultMaterial())
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3
	middle.Material.Reflective = 1

	right_shape := geom.DefaultSphere()
	right_shape.Translate(1.4, 0.5, -0.4).
		Scale(0.5, 0.5, 0.5)
	right := world.NewObject(&right_shape, rendering.DefaultMaterial())
	right.Material.Color = nmath.NewColor(0.5, 0.5, 0.1)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.9
	right.Material.Reflective = 1

	left_shape := geom.DefaultSphere()
	left_shape.Translate(-1.5, 0.33, -0.75).
		Scale(0.33, 0.33, 0.33)
	left := world.NewObject(&left_shape, rendering.DefaultMaterial())
	left.Material.Color = nmath.NewColor(1, 0.8, 0.1)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	light := rendering.NewPointLight(nmath.NewVec3(-5, 2, -10), nmath.NewColor(1, 1, 1))

	w := world.NewWorldWith(
		[]rendering.PointLight{light},
		[]world.Object{
			floor, ceiling, wall, right, middle, right, left,
		},
	)

	c := camera.NewCamera(600, 600, math.Pi/3.0)
	c.Transform = nmath.NewVec3(0, 1.5, -5).
		LookAt(
			nmath.NewVec3(0, 1, 0),
			nmath.NewVec3(0, 1, 0),
		)
	canvas := c.Render(w)

	fmt.Fprint(os.Stdout, canvas.AsPPM())
}
