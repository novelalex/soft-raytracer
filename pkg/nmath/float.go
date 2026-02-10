package nmath

import "math"

const F64Epsilon = 1e-6

const F64EpsilonLoose = 1e-4

func ApproxEq(lhs, rhs float64) bool {
	return math.Abs(lhs-rhs) <= F64Epsilon
}

// This is specificaly for testing calculations that use
// low accuracy math functions (ie. math.pow)
func LooseEq(lhs, rhs float64) bool {
	return math.Abs(lhs-rhs) <= F64EpsilonLoose
}
