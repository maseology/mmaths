package mmaths

import (
	"math"

	"github.com/maseology/mmaths/vector"
)

// Point is a 4D coordinate struct
type Point struct{ X, Y, Z, M float64 }

func (p *Point) ToArray() []float64 {
	return []float64{p.X, p.Y, p.Z, p.M}
}

func (p *Point) Distance(p0 Point) float64 {
	return vector.Distance([3]float64{p.X, p.Y, p.Z}, [3]float64{p0.X, p0.Y, p0.Z})
}

func (p *Point) MidPoint(p0 Point) Point {
	return Point{X: (p.X + p0.X) / 2., Y: (p.Y + p0.Y) / 2.}
}

func (p *Point) Rotate(angle float64, porig Point) Point {
	x := p.X - porig.X
	y := p.Y - porig.Y
	return Point{
		X: x*math.Cos(angle) - y*math.Sin(angle) + porig.X,
		Y: x*math.Sin(angle) + y*math.Cos(angle) + porig.Y,
	}
}
