package mmaths

import "math"

type Histogram struct {
	Levels []float64
	Bins   []int
	Width  float64
}

func NewHistogram(a []float64, nbins int) *Histogram {
	l, b := make([]float64, nbins), make([]int, nbins)

	n, x := func(a []float64) (min, max float64) {
		min, max = math.MaxFloat64, -math.MaxFloat64
		for _, v := range a {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
		return
	}(a)

	w := (x*1.000001 - n) / float64(nbins) // increasing maximum value to keep within the highest bin
	for i := range nbins {
		l[i] = (2*n + w*float64(2*i+1)) / 2. // midpoint
	}
	for _, v := range a {
		b[int((v-n)/w)]++
	}

	return &Histogram{l, b, w}
}

func (h *Histogram) BinFromLevel(v float64) int {
	for i := range len(h.Bins) {
		if v >= h.Levels[i]-h.Width/2 && v < h.Levels[i]+h.Width/2 {
			return i
		}
	}
	return -1
}
