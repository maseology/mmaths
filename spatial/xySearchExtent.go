package spatial

import (
	"sort"

	"github.com/maseology/mmaths"
	"github.com/maseology/mmaths/slice"
)

type XYextentSearch struct{ m [4]mmaths.IndexedSlice }

// New [Xn, Xx, Yn, Yx]
func (xye *XYextentSearch) New(exts [][4]float64) {
	// build search tree
	txn, txx, tyn, tyx := make([]float64, len(exts)), make([]float64, len(exts)), make([]float64, len(exts)), make([]float64, len(exts))
	for i, v := range exts {
		txn[i] = v[0]
		txx[i] = v[1]
		tyn[i] = v[2]
		tyx[i] = v[3]
	}

	var ixn, ixx, iyn, iyx mmaths.IndexedSlice
	ixn.New(txn)
	ixx.New(txx)
	iyn.New(tyn)
	iyx.New(tyx)
	sort.Sort(ixn)
	sort.Sort(ixx)
	sort.Sort(iyn)
	sort.Sort(iyx)

	xye.m = [...]mmaths.IndexedSlice{ixn, ixx, iyn, iyx}
}

func (xye *XYextentSearch) Contains(x, y float64) []int {
	x00 := slice.GetIndxFloat64(xye.m[0].Val, x)
	x11 := slice.GetIndxFloat64(xye.m[1].Val, x)
	y02 := slice.GetIndxFloat64(xye.m[2].Val, y)
	y13 := slice.GetIndxFloat64(xye.m[3].Val, y)

	set := slice.Intersect(xye.m[0].Indx[:x00], xye.m[1].Indx[x11:])
	set = slice.Intersect(set, xye.m[2].Indx[:y02])
	set = slice.Intersect(set, xye.m[3].Indx[y13:])
	return set
}

// Intersect Xn, Xx, Yn, Yx
func (xye *XYextentSearch) Intersect(ext [4]float64) []int {
	x0 := slice.GetIndxFloat64(xye.m[0].Val, ext[0])
	x1 := slice.GetIndxFloat64(xye.m[0].Val, ext[1])
	y0 := slice.GetIndxFloat64(xye.m[2].Val, ext[2])
	y1 := slice.GetIndxFloat64(xye.m[2].Val, ext[3])
	return slice.Intersect(xye.m[0].Indx[x0:x1], xye.m[2].Indx[y0:y1])
}
