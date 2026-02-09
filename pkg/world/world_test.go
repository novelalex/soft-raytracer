package world_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/novelalex/soft-raytracer/pkg/geom"
	"github.com/novelalex/soft-raytracer/pkg/nmath"
	"github.com/novelalex/soft-raytracer/pkg/world"
)

var _ = Describe("World", func() {
	Describe("IntersectRay", func() {
		It("should return all intersections of the ray with the world", func() {
			w := world.NewWorld()
			r := geom.NewRay(nmath.NewVec3(0, 0, -5), nmath.NewVec3(0, 0, 1))
			xs := w.IntersectRay(r)

			Expect(len(xs)).To(Equal(4))
			Expect(nmath.ApproxEq(xs[0].T, 4)).To(BeTrue())
			Expect(nmath.ApproxEq(xs[1].T, 4.5)).To(BeTrue())
			Expect(nmath.ApproxEq(xs[2].T, 5.5)).To(BeTrue())
			Expect(nmath.ApproxEq(xs[3].T, 6)).To(BeTrue())
		})
	})
})
