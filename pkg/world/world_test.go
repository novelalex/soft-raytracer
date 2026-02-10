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

	Describe("IsShadowed", func() {
		It("should return false when nothing is collinear with point and light", func() {
			w := world.NewWorld()
			p := nmath.NewVec3(0, 10, 0)
			Expect(w.IsShadowed(p, w.Lights[0])).To(BeFalse())
		})

		It("should return true when an object is between the point and the light", func() {
			w := world.NewWorld()
			p := nmath.NewVec3(10, -10, 10)
			Expect(w.IsShadowed(p, w.Lights[0])).To(BeTrue())
		})

		It("should return false when an object is behind the light", func() {
			w := world.NewWorld()
			p := nmath.NewVec3(-20, 20, -20)
			Expect(w.IsShadowed(p, w.Lights[0])).To(BeFalse())
		})

		It("should return false when an object is behind the point", func() {
			w := world.NewWorld()
			p := nmath.NewVec3(-2, 2, -2)
			Expect(w.IsShadowed(p, w.Lights[0])).To(BeFalse())
		})
	})
})
