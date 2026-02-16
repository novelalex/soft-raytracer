package raytracer

import (
	"runtime"
	"sync"

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

	jobs := make(chan renderWorkerJob, c.Width*c.Height)
	results := make(chan renderWorkerResult, c.Width*c.Height)

	num_workers := runtime.NumCPU()
	var wg sync.WaitGroup
	for i := 0; i < num_workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			renderWorker(jobs, results)
		}()
	}

	for y := range c.Height {
		for x := range c.Width {
			jobs <- renderWorkerJob{c, w, x, y}
		}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		image.WritePixel(result.X, result.Y, result.C)
	}

	return image
}

type renderWorkerJob struct {
	C *Camera
	W World
	X uint
	Y uint
}

type renderWorkerResult struct {
	C nmath.Color
	X uint
	Y uint
}

func renderWorker(jobs <-chan renderWorkerJob, results chan<- renderWorkerResult) {
	for job := range jobs {
		ray := job.C.RayForPixel(job.X, job.Y)
		color := job.W.ColorAt(ray, 5)
		results <- renderWorkerResult{
			color, job.X, job.Y,
		}
	}
}
