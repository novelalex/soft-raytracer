package nmath_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/novelalex/soft-raytracer/pkg/nmath"
)

var _ = Describe("Matrix", func() {
	Describe("ApproxEq", func() {
		Context("when matrices are the same", func() {
			It("should return true", func() {
				m1 := Mat4{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 8, 7, 6,
					5, 4, 3, 2,
				}
				m2 := Mat4{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 8, 7, 6,
					5, 4, 3, 2,
				}
				Expect(m1.ApproxEq(m2)).To(BeTrue())
			})
		})

		Context("when matrices are different", func() {
			It("should return false", func() {
				m1 := Mat4{
					1, 2, 3, 4,
					5, 6, 7, 2,
					9, 8, 7, 6,
					5, 4, 3, 2,
				}
				m2 := Mat4{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 58, 7, 6,
					5, 4, 3, 4,
				}
				Expect(m1.ApproxEq(m2)).To(BeFalse())
			})
		})
	})

	Describe("At", func() {
		Context("when matrix is indexed", func() {
			It("should return the element at the index", func() {
				m := Mat4{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 8, 7, 6,
					5, 4, 3, 2,
				}
				result := m.At(1, 2)
				expected := 8.0
				Expect(ApproxEq(result, expected)).To(BeTrue())
			})
		})
	})

	Describe("Mult", func() {
		Context("when matrix is multiplyed by another matrix", func() {
			It("should return the result of the multiplcation", func() {
				m1 := Mat4{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 8, 7, 6,
					5, 4, 3, 2,
				}
				m2 := Mat4{
					-2, 1, 2, 3,
					3, 2, 1, -1,
					4, 3, 6, 5,
					1, 2, 7, 8,
				}
				result := m1.Mult(m2)
				expected := Mat4{
					20, 22, 50, 48,
					44, 54, 114, 108,
					40, 58, 110, 102,
					16, 26, 46, 42,
				}

				Expect(result.ApproxEq(expected)).To(BeTrue())
			})
		})

		Context("when matrix is multiplyed by the identity matrix", func() {
			It("should return the original matrix", func() {
				m1 := Mat4{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 8, 7, 6,
					5, 4, 3, 2,
				}
				m2 := Mat4Identity()
				result := m1.Mult(m2)

				Expect(result.ApproxEq(m1)).To(BeTrue())
			})
		})
	})

	Describe("MultV", func() {
		Context("when matrix is multiplyed by a vector", func() {
			It("should return the vector result of the multiplcation", func() {
				m := Mat4{
					1, 2, 3, 4,
					2, 4, 4, 2,
					8, 6, 4, 1,
					0, 0, 0, 1,
				}
				v := Vec4{1, 2, 3, 1}
				result := m.MultV(v)
				expected := Vec4{18, 24, 33, 1}

				Expect(result.ApproxEq(expected)).To(BeTrue())
			})
		})
	})

	Describe("Transpose", func() {
		Context("when matrix is transposed", func() {
			It("should swap the rows and columns", func() {
				m := Mat4{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 8, 7, 6,
					5, 4, 3, 2,
				}
				result := m.Transpose()
				expected := Mat4{
					1, 5, 9, 5,
					2, 6, 8, 4,
					3, 7, 7, 3,
					4, 8, 6, 2,
				}

				Expect(result.ApproxEq(expected)).To(BeTrue())
			})
		})
	})

	Describe("Inverse", func() {
		Context("when matrix is inverted", func() {
			It("should return the inverse matrix", func() {
				m := Mat4{
					-5, 2, 6, -8,
					1, -5, 1, 8,
					7, 7, -6, -7,
					1, -3, 7, 4,
				}
				result := m.Inverse()
				expected := Mat4{
					0.21805, 0.45113, 0.24060, -0.04511,
					-0.80827, -1.45677, -0.44361, 0.52068,
					-0.07895, -0.22368, -0.05263, 0.19737,
					-0.52256, -0.81391, -0.30075, 0.30639,
				}

				Expect(result.LooseEq(expected)).To(BeTrue())
			})
		})
	})

})
