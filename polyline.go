package mmaths

import "math"

type Polyline struct {
	S [][]float64
}

func (p *Polyline) Chainage() [][]float64 {
	sdist := 0.
	chng := make([][]float64, len(p.S))
	chng[0] = []float64{0., p.S[0][2]}
	for i := 0; i < len(p.S)-1; i++ {
		x, y := p.S[i][0]-p.S[i+1][0], p.S[i][1]-p.S[i+1][1]
		dist := math.Sqrt(x*x + y*y)
		sdist += dist
		chng[i+1] = []float64{sdist, p.S[i+1][2]}
	}
	return chng
}
