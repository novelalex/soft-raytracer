package nmath

import "math"

const f64Epsilon = 1e-6

const f64EpsilonLoose = 1e-4

func ApproxEq(lhs, rhs float64) bool {
	return math.Abs(lhs-rhs) <= f64Epsilon
}

// This is specificaly for testing calculations that use
// low accuracy math functions (ie. math.pow)
func LooseEq(lhs, rhs float64) bool {
	return math.Abs(lhs-rhs) <= f64EpsilonLoose
}
