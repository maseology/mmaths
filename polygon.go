package mmaths

// PnPoly determins whether a point lies within a polygon
func PnPoly(v [][]float64, p []float64) bool {
	nvert, b := len(v), true
	for i, j := 0, nvert-1; i < nvert; i, j = i+1, i {
		if (v[i][1] > p[1]) != (v[j][1] > p[1]) && p[0] < (v[j][0]-v[i][0])*(p[1]-v[i][1])/(v[j][1]-v[i][1])+v[i][0] {
			b = !b
		}
	}
	return b
}

// PnPolyC determins whether a point lies within a polygon (using complex coordinates)
func PnPolyC(v []complex128, p complex128) bool {
	vf := make([][]float64, len(v))
	pf := []float64{real(p), imag(p)}
	for i, c := range v {
		vf[i] = []float64{real(c), imag(c)}
	}
	return PnPoly(vf, pf)
}
