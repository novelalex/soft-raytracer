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

	//	floor := geom.DefaultSphere()
	//	floor.Scale(10, 0.01, 10)
	//	floor.Mat.Color = nmath.NewColor(1, 0.9, 0.9)
	//	floor.Mat.Specular = 0
	//
	//	left_wall := geom.DefaultSphere()
	//	left_wall.Translate(0, 0, 5).
	//		RotateY(-math.Pi/4.0).
	//		RotateX(math.Pi/2.0).
	//		Scale(10, 0.01, 10)
	//	left_wall.Mat = floor.Mat
	//
	//	right_wall := geom.DefaultSphere()
	//	right_wall.Translate(0, 0, 5).
	//		RotateY(math.Pi/4.0).
	//		RotateX(math.Pi/2.0).
	//		Scale(10, 0.01, 10)
	//	right_wall.Mat = floor.Mat

	floor := geom.DefaultPlane()

	middle := geom.DefaultSphere()
	middle.Translate(-0.5, 1, 0.5)
	middle.Mat.Color = nmath.NewColor(0.1, 1, 0.5)
	middle.Mat.Diffuse = 0.7
	middle.Mat.Specular = 0.3

	right := geom.DefaultSphere()
	right.Translate(1.5, 0.5, -0.5).
		Scale(0.5, 0.5, 0.5)
	right.Mat.Color = nmath.NewColor(0.5, 1, 0.1)
	right.Mat.Diffuse = 0.7
	right.Mat.Specular = 0.9

	left := geom.DefaultSphere()
	left.Translate(-1.5, 0.33, -0.75).
		Scale(0.33, 0.33, 0.33)
	left.Mat.Color = nmath.NewColor(1, 0.8, 0.1)
	left.Mat.Diffuse = 0.7
	left.Mat.Specular = 0.3

	light := rendering.NewPointLight(nmath.NewVec3(-5, 10, -10), nmath.NewColor(1, 1, 1))

	w := world.World{
		Lights: []rendering.PointLight{light},
		Objects: []geom.Shape{
			&floor, &right, &middle, &right, &left,
		},
	}

	c := camera.NewCamera(300, 300, math.Pi/3.0)
	c.Transform = nmath.NewVec3(0, 1.5, -5).
		LookAt(
			nmath.NewVec3(0, 1, 0),
			nmath.NewVec3(0, 1, 0),
		)

	canvas := c.Render(w)

	fmt.Fprint(os.Stdout, canvas.AsPPM())
}
