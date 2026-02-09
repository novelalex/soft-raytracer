package nmath_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/novelalex/soft-raytracer/pkg/nmath"
)

var _ = Describe("Quadratics", func() {
	Describe("Solve", func() {
		Context("when given a, b, and c values", func() {
			It("should apply the quadratic formula", func() {
				result := Solve(1, -4, -8)
				expected := Solution{
					2.0 + 2.0*math.Sqrt(3.0),
					2.0 - 2.0*math.Sqrt(3.0),
				}
				Expect(result.ApproxEq(expected)).To(BeTrue())
			})
		})

	})

})
