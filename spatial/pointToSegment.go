package spatial

type PointPolylineConnection struct {
	Seg                int
	Dist, Frac, Length float64
	OrthPnt            [2]float64
}

func (sgs *SegmentSearch) PointToSegment(x, y, searchRadius float64) (PointPolylineConnection, []int) {

	isegs, bl1 := sgs.XYextentSearch.Contains(x, y), true // trial 1: (quickly) search by extent
	// isegs, bl1 := srchext.Contains(p.X, p.Y), true
	iclose, dist, fchain, totlen, opt := -1, -1., -1., -1., [2]float64{}
	// if i == -1519547369 {
	// 	print("") // for testing
	// }
retry:
	switch len(isegs) {
	case 0:
		if bl1 { // 1 interation of adding a buffer, retry
			isegs = sgs.XYextentSearch.Intersect([...]float64{
				x - searchRadius,
				x + searchRadius,
				y - searchRadius,
				y + searchRadius,
			})
			bl1 = false
			goto retry
		}
		// fmt.Println("none")
		isegs = nil
	// case 1:
	// 	fmt.Println(isegs[0]) // commented out to obtain dsv and fsv
	default:
		// fmt.Print(isegs)
		var srchsegs XYlineSearch
		segs := make([][][]float64, len(isegs))
		for i, id := range isegs {
			segs[i] = sgs.Segs[id]
		}
		srchsegs.New(segs)
		// iclose, dist := srchsegs.ClosestID([2]float64{p.X, p.Y})
		iclose, dist, fchain, totlen, opt = srchsegs.ClosestID([]float64{x, y}) // trial 2: search by distance to line segment
		fchain = 1. - fchain                                                    // chain is returned in an upstream direction
		if iclose < 0 {
			// fmt.Println("  none")
			isegs = nil
		} else {
			// fmt.Printf("  chosen: %d %.3f %.3f\n", isegs[iclose], dist, fchain)
			isegs = []int{isegs[iclose]}
		}
	}
	if len(isegs) == 0 {
		return PointPolylineConnection{}, nil
	}
	return PointPolylineConnection{
		Seg:     isegs[0],
		Dist:    dist,
		Frac:    fchain,
		Length:  totlen,
		OrthPnt: opt,
	}, isegs
}
