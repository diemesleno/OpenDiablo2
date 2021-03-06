package d2math

import "math"

const (
	// Epsilon is used as the threshold for 'almost equal' operations.
	Epsilon float64 = 0.0001

	// RadToDeg is used to convert anges in radians to degrees by multiplying the radians by RadToDeg. Similarly,degrees
	// are converted to radians when dividing by RadToDeg.
	RadToDeg float64 = 57.29578

	// RadFull is the radian equivalent of 360 degrees.
	RadFull float64 = 6.283185253783088
)

// EqualsApprox returns true if the difference between a and b is less than Epsilon.
func EqualsApprox(a, b float64) bool {
	return math.Abs(a-b) < Epsilon
}

// CompareFloat64Fuzzy returns an integer between -1 and 1 describing
// the comparison of floats a and b. 0 will be returned if the
// absolute difference between a and b is less than Epsilon.
func CompareFloat64Fuzzy(a, b float64) int {
	delta := a - b
	if math.Abs(delta) < Epsilon {
		return 0
	}

	if delta > 0 {
		return 1
	}

	return -1
}

// ClampFloat64 returns a clamped to min and max.
func ClampFloat64(a, min, max float64) float64 {
	if a > max {
		return max
	} else if a < min {
		return min
	}

	return a
}

// Sign returns the sign of a.
func Sign(a float64) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}

	return 0
}

// Lerp returns the linear interpolation from a to b using interpolator x.
func Lerp(a, b, x float64) float64 {
	return a + x*(b-a)
}

// Unlerp returns the intepolator Lerp would require to return x when given
// a and b. The x argument of this function can be thought of as the return
// value of lerp. The return value of this function can be used as x in
// Lerp.
func Unlerp(a, b, x float64) float64 {
	return (x - a) / (b - a)
}

// WrapInt wraps x to between 0 and max. For example WrapInt(450, 360) would return 90.
func WrapInt(x, max int) int {
	wrapped := x % max

	if wrapped < 0 {
		return max + wrapped
	}

	return wrapped
}
