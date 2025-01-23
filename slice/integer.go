package slice

import "math"

// MaxInt returns the maximum value of an integer slice
func MaxInt(s []int) int {
	x := -math.MaxInt
	for _, v := range s {
		if v > x {
			x = v
		}
	}
	return x
}
