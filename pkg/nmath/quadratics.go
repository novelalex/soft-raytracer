package nmath

import "math"

type Solution []float64

func Solve(a, b, c float64) Solution {
	discriminant := b*b - 4*a*c
	if discriminant < -f64Epsilon {
		return Solution{}
	} else if ApproxEq(discriminant, 0.0) {
		return Solution{(-b) / (2.0 * a)}
	} else {
		return Solution{
			((-b) + math.Sqrt(discriminant)) / (2.0 * a),
			((-b) - math.Sqrt(discriminant)) / (2.0 * a),
		}
	}
}

func (a Solution) ApproxEq(b Solution) bool {
	if len(a) != len(b) {
		return false
	} else if len(a) == 0 {
		return true
	} else if len(a) == 1 {
		return ApproxEq(a[0], b[0])
	}
	return (ApproxEq(a[0], b[0]) && ApproxEq(a[1], b[1])) ||
		(ApproxEq(a[0], b[1]) && ApproxEq(a[1], b[0]))

}
