package slice

import "math"

// OnlyPositive removes all value <= 0.0 and all NaN's
func OnlyPositive(s []float64) []float64 {
	// fmt.Println(len(s))
	// fmt.Println(s)
	var x []int
	for i := range s {
		if s[i] <= 0 || math.IsNaN(s[i]) {
			x = append(x, i)
		}
	}
	Rev(x)
	for _, i := range x {
		s = append(s[:i], s[i+1:]...)
	}
	return s[:len(s)]
	// // s = append([]float64(nil), s[:len(s)-len(x)]...)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// fmt.Println(len(x))
	// fmt.Println(s)
}
