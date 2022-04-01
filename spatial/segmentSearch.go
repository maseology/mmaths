package spatial

type SegmentSearch struct {
	XYextentSearch //  [Xn, Xx, Yn, Yx]
	Segs           [][][2]float64
}

func (sgs *SegmentSearch) New(segs [][][2]float64, exts [][4]float64) {
	sgs.XYextentSearch.New(exts)
	sgs.Segs = segs
}
