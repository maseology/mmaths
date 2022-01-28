package spatial

import (
	"sort"

	"github.com/maseology/mmaths"
)

type XYsearch struct {
	bn [2]*mmaths.BinaryNode
	is [2]mmaths.IndexedSlice
	xr [2][]int // indx to order
}

func (xys *XYsearch) New(pts [][2]float64) {

	xs, ys := make([]float64, len(pts)), make([]float64, len(pts))
	for i, c := range pts {
		xs[i] = c[0]
		ys[i] = c[1]
	}

	var xi, yi mmaths.IndexedSlice
	xi.New(xs)
	yi.New(ys)
	sort.Sort(xi)
	sort.Sort(yi)

	xys.xr = [...][]int{make([]int, len(xs)), make([]int, len(ys))}
	xys.is = [...]mmaths.IndexedSlice{xi, yi}
	for i, v := range xi.Indx {
		xys.xr[0][v] = i
	}
	for i, v := range yi.Indx {
		xys.xr[1][v] = i
	}

	var addNode func(*mmaths.IndexedSlice, *mmaths.BinaryNode, int, int)
	addNode = func(is *mmaths.IndexedSlice, node *mmaths.BinaryNode, first, last int) {
		// from a uniform distribution, picks a "balanced" tree
		if first <= last {
			nid := first + (last-first)/2 // median node. same as (first+last)/2, avoids overflow
			node.Indx = is.Indx[nid]
			node.Val = is.Val[nid]

			if first <= nid-1 {
				node.Left = &mmaths.BinaryNode{}
				addNode(is, node.Left, first, nid-1)
			}
			if nid+1 <= last {
				node.Right = &mmaths.BinaryNode{}
				addNode(is, node.Right, nid+1, last)
			}
		}
	}

	xys.bn = [...]*mmaths.BinaryNode{{}, {}}
	addNode(&xi, xys.bn[0], 0, len(xs)-1)
	addNode(&yi, xys.bn[1], 0, len(ys)-1)
}

// Value return array value at i
func (xys *XYsearch) Value(i int) [2]float64 {
	return [2]float64{
		xys.is[0].Val[xys.xr[0][i]],
		xys.is[1].Val[xys.xr[1][i]],
	}
}

// Value return array value at i
func (xys *XYsearch) Value3(i int) [3]float64 {
	return [3]float64{
		xys.is[0].Val[xys.xr[0][i]],
		xys.is[1].Val[xys.xr[1][i]],
		0.,
	}
}

func (xys *XYsearch) ClosestIDs(pt [2]float64, searchRadius float64) ([]int, []float64) {
	// collect closest points
	chc := make(chan map[int]bool)
	closest1D := func(dim int) {
		coll := map[int]bool{}
		isp := xys.bn[dim].IndexOf(pt[dim]) // closest feature ID
		switch isp {
		case -1: // closest to first
			// coll[xys.is[dim].Indx[0]] = true
			for i := 0; i < len(xys.is[dim].Indx); i++ {
				if xys.is[dim].Val[i]-pt[dim] > searchRadius {
					break
				}
				coll[xys.is[dim].Indx[i]] = true
			}
		case len(xys.is[dim].Val): // closest to last
			// coll[xys.is[dim].Indx[len(xys.is[dim].Val)-1]] = true
			for i := len(xys.is[dim].Indx) - 1; i >= 0; i-- {
				if xys.is[dim].Val[i]-pt[dim] > searchRadius {
					break
				}
				coll[xys.is[dim].Indx[i]] = true
			}
		default:
			xi := xys.xr[dim][isp] // feature order
			for pt[dim]-xys.is[dim].Val[xi] < searchRadius {
				coll[xys.is[dim].Indx[xi]] = true
				xi--
				if xi < 0 {
					break
				}
			}
			ix := len(xys.xr[0])
			if isp < ix {
				xi := xys.xr[dim][isp] + 1 // feature order
				// fmt.Println(dim, xi)
				if xi < ix {
					for xys.is[dim].Val[xi]-pt[dim] < searchRadius {
						coll[xys.is[dim].Indx[xi]] = true
						xi++
						if xi >= ix {
							break
						}
					}
				}
			}
		}
		chc <- coll
	}
	go closest1D(0)
	go closest1D(1)
	c0 := <-chc
	c1 := <-chc
	close(chc)
	// collect closest points
	// c0, c1 := func() (map[int]bool, map[int]bool) {
	// 	closest1D := func(dim int) map[int]bool {
	// 		coll := map[int]bool{}
	// 		isp := xys.bn[dim].IndexOf(pt[dim]) // closest feature ID
	// 		switch isp {
	// 		case -1: // before first
	// 			// coll[xys.is[dim].Indx[0]] = false
	// 			for i := 0; i < len(xys.is[dim].Indx); i++ {
	// 				if xys.is[dim].Val[i]-pt[dim] > searchRadius {
	// 					break
	// 				}
	// 				coll[xys.is[dim].Indx[i]] = true
	// 			}
	// 		case len(xys.is[dim].Val): // after last
	// 			// coll[xys.is[dim].Indx[len(xys.is[dim].Val)-1]] = false
	// 			for i := len(xys.is[dim].Indx) - 1; i >= 0; i-- {
	// 				if xys.is[dim].Val[i]-pt[dim] > searchRadius {
	// 					break
	// 				}
	// 				coll[xys.is[dim].Indx[i]] = true
	// 			}
	// 		default:
	// 			xi := xys.xr[dim][isp] // feature order
	// 			for pt[dim]-xys.is[dim].Val[xi] < searchRadius {
	// 				coll[xys.is[dim].Indx[xi]] = true
	// 				xi--
	// 				if xi < 0 {
	// 					break
	// 				}
	// 			}
	// 			ix := len(xys.xr[0])
	// 			if isp < ix {
	// 				xi := xys.xr[dim][isp] + 1 // feature order
	// 				// fmt.Println(dim, xi)
	// 				if xi < ix {
	// 					for xys.is[dim].Val[xi]-pt[dim] < searchRadius {
	// 						coll[xys.is[dim].Indx[xi]] = true
	// 						xi++
	// 						if xi >= ix {
	// 							break
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 		return coll
	// 	}
	// 	return closest1D(0), closest1D(1)
	// }()

	sr2 := searchRadius * searchRadius
	cocoll := map[int]float64{}
	dist2 := func(p, q [2]float64) float64 {
		s2 := 0.
		for i := 0; i < 2; i++ {
			ds := q[i] - p[i]
			s2 += ds * ds
		}
		return s2
	}
	for i := range c0 {
		if _, ok := c1[i]; ok {
			xi, yi := xys.xr[0][i], xys.xr[1][i] // feature order
			d := dist2(pt, [...]float64{xys.is[0].Val[xi], xys.is[1].Val[yi]})
			if d < sr2 {
				cocoll[i] = d
			}
		}
	}
	return mmaths.SortMapFloat(cocoll) // point IDs, distances -- sorted by distance
}
