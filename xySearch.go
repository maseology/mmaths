package mmaths

import (
	"sort"
)

type XYsearch struct {
	bn [2]*binaryNode
	is [2]IndexedSlice
	xr [2][]int // indx to order
}

func (xys *XYsearch) New(pts [][2]float64) {

	xs, ys := make([]float64, len(pts)), make([]float64, len(pts))
	for i, c := range pts {
		xs[i] = c[0]
		ys[i] = c[1]
	}

	var xi, yi IndexedSlice
	xi.New(xs)
	yi.New(ys)
	sort.Sort(xi)
	sort.Sort(yi)

	xys.xr = [...][]int{make([]int, len(xs)), make([]int, len(ys))}
	xys.is = [...]IndexedSlice{xi, yi}
	for i, v := range xi.Indx {
		xys.xr[0][v] = i
	}
	for i, v := range yi.Indx {
		xys.xr[1][v] = i
	}

	var addNode func(*IndexedSlice, *binaryNode, int, int)
	addNode = func(is *IndexedSlice, node *binaryNode, first, last int) {
		// from a uniform distribution, picks a "balanced" tree
		if first <= last {
			nid := first + (last-first)/2 // median node. same as (first+last)/2, avoids overflow
			node.indx = is.Indx[nid]
			node.val = is.Val[nid]

			if first <= nid-1 {
				node.left = &binaryNode{}
				addNode(is, node.left, first, nid-1)
			}
			if nid+1 <= last {
				node.right = &binaryNode{}
				addNode(is, node.right, nid+1, last)
			}
		}
	}

	xys.bn = [...]*binaryNode{{}, {}}
	addNode(&xi, xys.bn[0], 0, len(xs)-1)
	addNode(&yi, xys.bn[1], 0, len(ys)-1)
}

func (xys *XYsearch) ClosestIDs(pt [2]float64, searchRadius float64) ([]int, []float64) {
	// collect closest points
	chc := make(chan map[int]bool)
	closest1D := func(dim int) {
		isp := xys.bn[dim].IndexOf(pt[dim]) // closest feature ID
		coll := map[int]bool{}
		if isp > -1 {
			xi := xys.xr[dim][isp] // feature order
			for pt[dim]-xys.is[dim].Val[xi] < searchRadius {
				coll[xys.is[dim].Indx[xi]] = true
				xi--
				if xi < 0 {
					break
				}
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
		chc <- coll
	}
	go closest1D(0)
	go closest1D(1)
	c0 := <-chc
	c1 := <-chc
	close(chc)

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
	return SortMapFloat(cocoll) // point IDs, distances -- sorted by distance
}
