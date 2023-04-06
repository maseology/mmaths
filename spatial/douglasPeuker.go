package spatial

import "github.com/maseology/mmaths/vector"

// DouglasPeucker the Ramer–Douglas–Peucker algorithm
// see: https://en.wikipedia.org/wiki/Ramer%E2%80%93Douglas%E2%80%93Peucker_algorithm
func DouglasPeucker(plns [][][]float64, epsilon float64) ([][][]float64, [][]int, int) {
	isnear := func(p0, p1 []float64) bool {
		d2 := 0.
		for i := range p0 {
			diff := p0[i] - p1[i]
			d2 += diff * diff
			if i == 3 {
				break
			}
		}
		return d2 < .001
	}
	remove := func(seg [][]float64, pos []int, s int) ([][]float64, []int) {
		return append(seg[:s], seg[s+1:]...), append(pos[:s], pos[s+1:]...)
	}

	nvert := 0
	posi := make([][]int, len(plns))
	for plid, pln := range plns {
		ipos := make([]int, len(pln))
		for i := range pln {
			ipos[i] = i
		}

		if isnear(pln[0], pln[len(pln)-1]) { // check for polygons
			pln, ipos = remove(pln, ipos, len(pln)-1)
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
				pln, ipos = remove(pln, ipos, len(pln)-1)
			}
		}
		plns[plid] = pln
		posi[plid] = ipos
		nvert += len(pln)
	}
	return plns, posi, nvert
}
