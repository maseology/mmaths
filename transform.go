package mmaths

import (
	"log"
	"math"
)

// LinearTransform linearly transforms U[0.0,1.0] to [l,h] space
func LinearTransform(l, h, u float64) float64 {
	if u < 0.0 || u > 1.0 {
		log.Fatalf("linear transform error, passing u = %v", u)
	}
	return (h-l)*u + l
}

// LogLinearTransform linearly transforms U[0.0,1.0] to 10^[l,h] space
func LogLinearTransform(l, h, u float64) float64 {
	if l <= 0.0 || h <= 0.0 {
		log.Fatalf("Log linear transform range error, [l,h] = [%v,%v]", l, h)
	}
	if u < 0.0 || u > 1.0 {
		log.Fatalf("Log linear transform error, passing u = %v", u)
	}
	// return math.Pow(10.0, math.Log10(l*math.Pow(h/l, u)))
	return l * math.Pow(h/l, u)
}
