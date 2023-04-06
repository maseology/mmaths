package mmaths

import "math"

// Extent is a 2D square spatial extent
type Extent struct{ Xn, Xx, Yn, Yx float64 }

func (ex *Extent) New(s [][]float64) {
	ex.Xn = math.MaxFloat64
	ex.Xx = -math.MaxFloat64
	ex.Yn = math.MaxFloat64
	ex.Yx = -math.MaxFloat64
	for _, c := range s {
		if c[0] < ex.Xn {
			ex.Xn = c[0]
		}
		if c[0] > ex.Xx {
			ex.Xx = c[0]
		}
		if c[1] < ex.Yn {
			ex.Yn = c[1]
		}
		if c[1] > ex.Yx {
			ex.Yx = c[1]
		}
	}
}

func (ex *Extent) Radius() float64 {
	dx, dy := (ex.Xx-ex.Xn)/2., (ex.Yx-ex.Yn)/2.
	return math.Sqrt(dx*dx + dy*dy)
}

func (ex *Extent) Contains(p *Point) bool {
	if p.X >= ex.Xn && p.X <= ex.Xx {
		if p.Y >= ex.Yn && p.Y <= ex.Yx {
			return true
		}
	}
	return false
}
