package benchmark

import "math"

func floatABS(x int) int {
	return int(math.Abs(float64(x)))
}
func multABS(x int) int {
	return -x
}
