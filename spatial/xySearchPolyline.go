package spatial

import (
	"math"

	"github.com/maseology/mmaths/vector"
)

type XYlineSearch struct {
	sseg  []*XYsearch
	chain [][]float64
}

func (xys *XYlineSearch) New(segs [][][]float64) {
	xys.sseg = make([]*XYsearch, len(segs))
	xys.chain = make([][]float64, len(segs))
	for i, s := range segs {
		var srch XYsearch
		sum := 0.
		srch.New(s)
		xys.sseg[i] = &srch
		xys.chain[i] = make([]float64, len(s))
		for j := 1; j < len(s); j++ {
			sum += vector.Distance(xys.sseg[i].Value3(j), xys.sseg[i].Value3(j-1))
			xys.chain[i][j] = sum
		}
	}
}

func (xys *XYlineSearch) ClosestID(pt []float64) (int, float64, float64, float64, [2]float64) { // brute force
	dn, fchsv, lnsv, isv, opsv := math.MaxFloat64, -1., -1., -1, [2]float64{-1, -1}
	for i, s := range xys.sseg {
		ids, _ := s.ClosestIDs(pt, math.MaxFloat64)
		switch len(ids) {
		case 0:
			continue
		case 1:
			p0 := s.Value3(ids[0])
			d := vector.Distance([3]float64{pt[0], pt[1], 0.}, p0)
			if d < dn {
				dn = d
				fchsv = 0.
				isv = i
			}
		default:
			ip0 := ids[0] // closest vertex to point
			p0, ip1 := s.Value3(ip0), func() []int {
				switch ip0 {
				case 0:
					return []int{1}
				case len(ids) - 1:
					return []int{len(ids) - 2}
				default:
					return []int{ip0 - 1, ip0 + 1}
				}
			}()
			for _, ip := range ip1 {
				d, f, op := func() (d, f float64, op [3]float64) {
					p1 := s.Value3(ip)
					if ip < ip0 {
						d, f, op = vector.PointToLine([3]float64{pt[0], pt[1], 0.}, p1, p0)
					} else {
						d, f, op = vector.PointToLine([3]float64{pt[0], pt[1], 0.}, p0, p1)
					}
					return
				}()
				if d < dn {
					dn = d
					lnsv, fchsv = func(a []float64) (float64, float64) {
						totlen := a[len(a)-1]
						if ip < ip0 {
							return totlen, (f*(a[ip+1]-a[ip]) + a[ip]) / totlen
						} else {
							return totlen, (f*(a[ip0+1]-a[ip0]) + a[ip0]) / totlen
						}
					}(xys.chain[i])
					isv = i
					opsv = [2]float64{op[0], op[1]}
				}
			}

		}
	}
	if isv < 0 {
		return -1, 0., -1., -1., [2]float64{-9999., -9999.}
	}
	return isv, dn, fchsv, lnsv, opsv
}
