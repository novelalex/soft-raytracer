package rendering

import "github.com/novelalex/soft-raytracer/pkg/nmath"

type Camera struct {
	Width     int
	Height    int
	FOV       float64
	Transform nmath.Mat4
}

func NewCamera(w, h int, fov float64) Camera {
	return Camera{
		w, h, fov, nmath.Mat4Identity(),
	}
}
