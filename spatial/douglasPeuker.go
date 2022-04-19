package spatial

import "github.com/maseology/mmaths/vector"

// DouglasPeucker the Ramer–Douglas–Peucker algorithm
// see: https://en.wikipedia.org/wiki/Ramer%E2%80%93Douglas%E2%80%93Peucker_algorithm
func DouglasPeucker(plns [][][]float64, epsilon float64) [][][]float64 {
	panic("untested...")
	isnear := func(p0, p1 []float64) bool {
		d2 := 0.
		for i := range p0 {
			diff := p0[i] - p1[i]
			d2 += diff * diff
		}
		return d2 < .001
	}
	remove := func(seg [][]float64, s int) [][]float64 {
		return append(seg[:s], seg[s+1:]...)
	}
	for _, pln := range plns {
		if isnear(pln[0], pln[len(pln)-1]) {
			pln = remove(pln, len(pln)-1)
		}

		rm, crm := make([]bool, len(pln)), 0
		var recurse func(int, int)
		recurse = func(i0, i1 int) {
			dx, ix := 0., -1
			for i := i0 + 1; i < i1; i++ {
				d, _, _ := vector.PointToLine([3]float64{pln[i][0], pln[i][1], 0.},
					[3]float64{pln[i0][0], pln[i0][1], 0.},
					[3]float64{pln[i1][0], pln[i1][1], 0.})
				if d > dx {
					dx = d
					ix = i
				}
			}
			if dx > epsilon {
				recurse(i0, ix)
				recurse(ix, i1)
			} else {
				for i := i0 + 1; i < i1; i++ {
					if rm[i] {
						panic("should not happen")
					}
					rm[i] = true
					crm += 1
				}
			}
		}
		recurse(0, len(pln)-1)

		for i := len(pln) - 1; i >= 0; i-- {
			if !rm[i] {
				pln = remove(pln, len(pln)-1)
			}
		}
		// pln = func() [][]float64 {
		// 	o, ii := make([][]float64, len(pln)-crm), 0
		// 	for i := 0; i < len(pln); i++ {
		// 		if !rm[i] {
		// 			o[ii] = pln[i]
		// 			ii++
		// 		}
		// 	}
		// 	return o
		// }()
	}
	return plns
}
