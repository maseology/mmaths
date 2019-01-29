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
