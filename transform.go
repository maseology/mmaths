package maths

import "log"

// LinearTransform linearly transforms U[0.0,1.0] to [l,h] space
func LinearTransform(l, h, u float64) float64 {
	if u < 0.0 || u > 1.0 {
		log.Panicf("linear transform error, passing u = %v", u)
	}
	// if p.Log {
	// 	return math.Pow(10.0, math.Log10(p.Low*math.Pow(p.High/p.Low, u)))
	// }
	return (h-l)*u + l
}
