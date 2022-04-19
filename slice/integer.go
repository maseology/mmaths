package slice

import "math"

// Max returns the maximum value of an integer slice
func Max(s []int) int {
	x := -math.MaxInt
	for _, v := range s {
		if v > x {
			x = v
		}
	}
	return x
}
