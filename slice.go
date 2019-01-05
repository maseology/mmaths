// just a set of slice manipulation tools

package mmaths

import (
	"math"
)

// Rev is quick function used to reverse order of a slice
func Rev(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// RevF is quick function used to reverse order of a float64 slice
func RevF(s []float64) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Sequential builds a n+1 length slice indexed from 0 to n
func Sequential(n int) []int {
	iout := make([]int, n+1)
	for i := 0; i <= n; i++ {
		iout[i] = i
	}
	return iout
}

// SliceMax returns the maximum value of a slice
func SliceMax(s []float64) float64 {
	x := -math.MaxFloat64
	for _, v := range s {
		x = math.Max(x, v)
	}
	return x
}

// SliceMin returns the minimum value of a slice
func SliceMin(s []float64) float64 {
	x := math.MaxFloat64
	for _, v := range s {
		x = math.Min(x, v)
	}
	return x
}
