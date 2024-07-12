package slice

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

func SliceMed90(s []float64) (_, _, _ float64) {
	sort.Float64s(s)
	ns := float64(len(s))
	return s[int(ns/2.)], s[int(ns*.05)], s[int(ns*.95)]
}

func SliceMinMaxMean(s []float64) (_, _, _ float64) {
	x, n, m := -math.MaxFloat64, math.MaxFloat64, 0.
	for _, v := range s {
		x = math.Max(x, v)
		n = math.Min(n, v)
		m += v
	}
	return n, x, m / float64(len(s))
}

func SliceMeansd(d []float64) (float64, float64) {
	c, s, ss := 0, 0., 0.
	for i := 0; i < len(d); i++ {
		if math.IsNaN(d[i]) {
			continue
		}
		s += d[i]
		c++
	}
	s /= float64(c)
	for i := 0; i < len(d); i++ {
		if math.IsNaN(d[i]) {
			continue
		}
		ss += math.Pow(d[i]-s, 2.)
	}
	return s, math.Sqrt(ss / float64(c-1))
}

func R(x, y []float64) (float64, float64, float64, float64) {
	// see page 367 & 396 of Walpole Meyers Myers
	if len(x) != len(y) {
		panic("Coefficient of determination error: unequal array lengths")
	}
	c, mx, my := 0, 0., 0.
	for i := 0; i < len(x); i++ {
		if math.IsNaN(x[i]) || math.IsNaN(y[i]) {
			continue
		}
		mx += x[i]
		my += y[i]
		c++
	}
	mx /= float64(c)
	my /= float64(c)

	sxx, syy, sxy := 0., 0., 0.
	for i := 0; i < len(x); i++ {
		if math.IsNaN(x[i]) || math.IsNaN(y[i]) {
			continue
		}
		sxx += math.Pow(x[i]-mx, 2.)     // variance in x
		syy += math.Pow(y[i]-my, 2.)     // variance in y
		sxy += (x[i] - mx) * (y[i] - my) // covariance
	}

	if sxx > 0. && syy > 0. {
		return math.Copysign(sxy/math.Sqrt(sxx*syy), sxy), sxx, syy, sxy
	}
	return math.NaN(), sxx, syy, sxy
}
