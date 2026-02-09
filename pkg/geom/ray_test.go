package geom_test

import (
	. "github.com/novelalex/soft-raytracer/pkg/geom"
	nm "github.com/novelalex/soft-raytracer/pkg/nmath"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ray", func() {
	Describe("At (position)", func() {
		It("works", func() {
			cases := []struct {
				t        float64
				expected nm.Vec3
			}{
				{0, nm.NewVec3(2, 3, 4)},
				{1, nm.NewVec3(3, 3, 4)},
				{-1, nm.NewVec3(1, 3, 4)},
				{2.5, nm.NewVec3(4.5, 3, 4)},
			}

			ray := NewRay(nm.NewVec3(2, 3, 4), nm.NewVec3(1, 0, 0))

			for _, c := range cases {
				result := ray.At(c.t)
				Expect(result.ApproxEq(c.expected)).To(BeTrue())
			}
		})
	})

	Describe("Transform", func() {
		It("can translate", func() {
			r := NewRay(nm.NewVec3(1, 2, 3), nm.NewVec3(0, 1, 0))
			m := nm.NewTranslation(3, 4, 5)
			expected := NewRay(nm.NewVec3(4, 6, 8), nm.NewVec3(0, 1, 0))

			result := r.Transform(m)

			Expect(result.ApproxEq(expected)).To(BeTrue())
		})

		It("can scale", func() {
			r := NewRay(nm.NewVec3(1, 2, 3), nm.NewVec3(0, 1, 0))
			m := nm.NewScaling(2, 3, 4)
			expected := NewRay(nm.NewVec3(2, 6, 12), nm.NewVec3(0, 3, 0))

			result := r.Transform(m)

			Expect(result.ApproxEq(expected)).To(BeTrue())
		})
	})

	Describe("Translate", func() {
		It("translates the ray using the convenience method", func() {
			r := NewRay(nm.NewVec3(1, 2, 3), nm.NewVec3(0, 1, 0))
			expected := NewRay(nm.NewVec3(4, 6, 8), nm.NewVec3(0, 1, 0))

			result := r.Translate(3, 4, 5)

			Expect(result.ApproxEq(expected)).To(BeTrue())
		})
	})

	Describe("Scale", func() {
		It("scales the ray using the convenience method", func() {
			r := NewRay(nm.NewVec3(1, 2, 3), nm.NewVec3(0, 1, 0))
			expected := NewRay(nm.NewVec3(2, 6, 12), nm.NewVec3(0, 3, 0))

			result := r.Scale(2, 3, 4)

			Expect(result.ApproxEq(expected)).To(BeTrue())
		})
	})
})
