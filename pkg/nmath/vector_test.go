package nmath_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/novelalex/soft-raytracer/pkg/nmath"
)

var _ = Describe("Vector", func() {
	Describe("Reflect", func() {
		Context("when reflecting vectors", func() {
			It("should negate the reflected axis", func() {
				v := NewVec3(1, -1, 0)
				n := NewVec3(0, 1, 0)
				result := v.Reflect(n)
				expected := NewVec3(1, 1, 0)
				Expect(result.ApproxEq(expected)).To(BeTrue())
			})

			It("should reflect slanted angles", func() {
				v := NewVec3(1, -1, 0)
				n := NewVec3(0, 1, 0)
				result := v.Reflect(n)
				expected := NewVec3(1, 1, 0)
				Expect(result.ApproxEq(expected)).To(BeTrue())
			})
		})

	})

})
