package mmaths

import (
	"fmt"
	"math"

	"github.com/maseology/mmaths/vector"
)

// Plane defined as: ax + by + cz + d = 0
// normal vector is (-a,-b,1)
// slope = Sqrt(a ^ 2 + b ^ 2)
// aspect = Atan2(-b, -a) // Asin(-a / Sqrt(a ^ 2 + b ^ 2)) // theta = asin(-a/sqrt(a^2+b^2))
type Plane struct{ a, b, c, d float64 }

func NewPlane(ps []Point) (*Plane, error) {
	if len(ps) != 3 {
		return nil, fmt.Errorf(" planes need to be defined by only 3 points")
	}
	xn, yn, zn := math.MaxFloat64, math.MaxFloat64, math.MaxFloat64
	for _, p := range ps {
		if p.X < xn {
			xn = p.X
		}
		if p.Y < yn {
			yn = p.Y
		}
		if p.Z < zn {
			zn = p.Z
		}
	}
	for _, p := range ps {
		p.X -= xn
		p.Y -= yn
		p.Z -= zn
	}

	n := func() []float64 {
		p0, p1, p2 := [3]float64{ps[0].X, ps[0].Y, ps[0].Z}, [3]float64{ps[0].X, ps[0].Y, ps[0].Z}, [3]float64{ps[0].X, ps[0].Y, ps[0].Z}
		o := vector.Unit(vector.Cross(vector.Subtract(p1, p0), vector.Subtract(p2, p0)))
		return []float64{o[0], o[1], o[2]}
	}()
	if math.Atan2(math.Sqrt(n[0]*n[0]+n[1]*n[1]), n[2]) > math.Pi/2 {
		for _, nn := range n {
			nn *= -1
		}
	}

	soln := n[0]*ps[0].X + n[1]*ps[0].Y + n[2]*ps[0].Z
	return &Plane{n[0], n[1], n[2], -soln}, nil
}
