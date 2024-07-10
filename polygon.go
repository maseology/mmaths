package mmaths

import "math"

// PnPoly determines whether a point lies within a polygon
func PnPoly(v [][]float64, p []float64) bool {
	nvert, b := len(v), true
	for i, j := 0, nvert-1; i < nvert; i, j = i+1, i {
		if (v[i][1] > p[1]) != (v[j][1] > p[1]) && p[0] < (v[j][0]-v[i][0])*(p[1]-v[i][1])/(v[j][1]-v[i][1])+v[i][0] {
			b = !b
		}
	}
	return !b
}

// PnPolyLong determines whether a point lies within a polygon with more rigor: PnPoly requires the point to be completely within prism
func PnPolyLong(v [][]float64, p []float64, tol float64) bool {
	if PnPoly(v, p) {
		return true
	}

	// distance functions
	sqdist := func(p1, p2 []float64) float64 {
		return math.Pow(p1[0]-p2[0], 2.0) + math.Pow(p1[1]-p2[1], 2.0)
	}
	perpdist := func(p, p1, p2 []float64) float64 {
		return math.Abs((p2[1]-p1[1])*p[0]-(p2[0]-p1[0])*p[1]+p2[0]*p1[1]-p2[1]*p1[0]) / math.Sqrt(sqdist(p1, p2))
	}

	// first check distance to vertices
	for _, v := range v {
		if math.Sqrt(sqdist(p, v)) < tol {
			return true
		}
	}
	// check distance to 2-point line segment
	for i := range v {
		ii := (i + 1) % len(v)
		p1, p2 := v[i], v[ii]
		// build simple 2-vertex line
		if perpdist(p, p1, p2) < tol { // perpendicular distance
			// check if point projects onto line
			c := sqdist(p1, p2)
			a := sqdist(p1, p)
			b := sqdist(p2, p)
			if c >= a+b {
				return true
			}
		}
	}
	return false
}

// PnPolyC determines whether a point lies within a polygon (using complex coordinates)
func PnPolyC(v []complex128, p complex128, tolerance float64) bool {
	vf := make([][]float64, len(v))
	pf := []float64{real(p), imag(p)}
	for i, c := range v {
		vf[i] = []float64{real(c), imag(c)}
	}
	// if withRigor {
	// 	return PnPolyLong(vf, pf, .00001)
	// }
	// return PnPoly(vf, pf)
	return PnPolyLong(vf, pf, tolerance)
}
