package spatial

type SegmentSearch struct {
	XYextentSearch //  [Xn, Xx, Yn, Yx]
	Segs           [][][]float64
}

func (sgs *SegmentSearch) New(segs [][][]float64, exts [][4]float64) {
	sgs.XYextentSearch.New(exts)
	sgs.Segs = segs
}
