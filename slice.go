package mmaths

import (
	"math"
	"sort"
)

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

// SliceMean returns the mean value of a slice
func SliceMean(s []float64) float64 {
	x := 0.
	for _, v := range s {
		x += v
	}
	return x / float64(len(s))
}

// SliceMedian returns the median value of a slice
func SliceMedian(s []float64) float64 {
	sort.Float64s(s)
	return s[int(float64(len(s)/2.))]
}
