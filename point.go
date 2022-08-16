package mmaths

import "github.com/maseology/mmaths/vector"

// Point is a 4D coordinate struct
type Point struct{ X, Y, Z, M float64 }

func (p *Point) Distance(p0 *Point) float64 {
	return vector.Distance([3]float64{p.X, p.Y, p.Z}, [3]float64{p0.X, p0.Y, p0.Z})
}
