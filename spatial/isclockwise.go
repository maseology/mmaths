package spatial

func IsClockwise(x, y []float64) bool {
	s, n := 0., len(x)
	for j := range x {
		jj := (j + 1) % n
		s += (x[jj] - x[j]) * (y[jj] + y[j])
	}
	return s >= 0.
}
