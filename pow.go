package mmaths

// IntPow returns integer i to the power p
func IntPow(i, p int) int {
	if p <= 0 {
		panic("IntPow error, p <= 0")
	}
	t := 1
	for j := 0; j < p; j++ {
		t *= i
	}
	return t
}

// BytePow returns integer i to the power p
func BytePow(p int) byte {
	// if p <= 0 {
	// 	panic("BytePow error, p <= 0")
	// } else if p >= 8 {
	// 	panic("BytePow error, p >= 8")
	// }
	t := byte(1)
	for j := 0; j < p; j++ {
		t *= byte(2)
	}
	return t
}
