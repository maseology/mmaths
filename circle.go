package mmaths

import "math"

type Circle struct {
	Centroid Point
	Radius   float64
}

func (c *Circle) Contains(p []float64) bool {
	return math.Sqrt(math.Pow(c.Centroid.X-p[0], 2.)+math.Pow(c.Centroid.Y-p[1], 2.)) <= c.Radius
}
