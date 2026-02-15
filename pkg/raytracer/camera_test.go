package raytracer_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/raytracer"
)

var _ = Describe("Camera", func() {
	Describe("Render", func() {
		Context("when rendering the world", func() {
			It("should output the canvis with the world rendered on it", func() {
				w := raytracer.NewWorld()
				c := raytracer.NewCamera(11, 11, math.Pi/2.0)
				from := nmath.NewVec3(0, 0, -5)
				to := nmath.NewVec3(0, 0, 0)
				up := nmath.NewVec3(0, 1, 0)
				c.Transform = from.LookAt(to, up)
				image := c.Render(w)
				Expect(image.PixelAt(5, 5).AsVec3().ApproxEq(nmath.NewVec3(0.38066119308103435, 0.47582649135129296, 0.28549589481077575))).To(BeTrue())
			})
		})
	})
})
