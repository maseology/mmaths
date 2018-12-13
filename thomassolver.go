package mmaths

// ThomasBoundaryCondition is a Thomas algorithm solution
// which solves a tri-diagonal system of equations
// converted from python code in: Bittelli, M., Campbell, G.S., and Tomei, F., 2015. Soil Physics with Python. Oxford University Press.
func ThomasBoundaryCondition(a, b, c, d, x map[int]float64, first, last int) {
	for i := first; i < last; i++ {
		c[i] /= b[i]
		d[i] /= b[i]
		b[i+1] -= a[i+1] * c[i]
		d[i+1] -= a[i+1] * d[i]
	}
	// back substitution
	x[last] = d[last] / b[last]
	for i := last - 1; i > first-1; i-- {
		x[i] = d[i] - c[i]*x[i+1]
	}
}
