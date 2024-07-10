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

func (p *Polyline) Intersections(pln *Polyline) [][]float64 {
	ln0, ex0 := p.breakApart()
	ln1, ex1 := pln.breakApart()
	var ints [][]float64
	for i0, l0 := range ln0 {
		for i1, l1 := range ln1 {
			if ex0[i0].Intersects(ex1[i1], 0.) {
				xy, _ := l0.Intersection2D(&l1)
				if xy != nil {
					ints = append(ints, []float64{xy.X, xy.Y})
				}
			}
		}
	}
	return ints
}

func (p *Polyline) breakApart() ([]LineSegment, []Extent) {
	segs, exts := make([]LineSegment, len(p.S)-1), make([]Extent, len(p.S)-1)
	for i := range len(p.S) - 2 {
		ls := LineSegment{P0: Point{X: p.S[i][0], Y: p.S[i][1]}, P1: Point{X: p.S[i+1][0], Y: p.S[i+1][1]}}
		ls.Build()
		segs[i] = ls
		exts[i].New([][]float64{{p.S[i][0], p.S[i][1]}, {p.S[i+1][0], p.S[i+1][1]}})
	}
	return segs, exts
}

// Dim lstOUT As New List(Of Line)
// For i = 0 To _v.Count - 2
// 	lstOUT.Add(New Line(_v(i), _v(i + 1)))
// Next
// Return lstOUT
// End Function
