package mmaths

// RelativeDifference returns the relative difference between to values
func RelativeDifference(f0, f1 float64) float64 {
	return (f1 - f0) / f0
}
