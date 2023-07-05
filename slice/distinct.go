package slice

// Distinct returns a unique slice
func Distinct(ints []int) []int {
	eval := make(map[int]bool)
	l := []int{}
	for _, v := range ints {
		if eval[v] {
			// do nothing
		} else {
			eval[v] = true
			l = append(l, v)
		}
	}
	// // does not preserve order
	// k := make(map[int]bool)
	// for _, i := range ints {
	// 	if _, ok := k[i]; !ok {
	// 		k[i] = true
	// 	}
	// }
	// l, ii := make([]int, len(k)), 0
	// for i := range k {
	// 	l[ii] = i
	// 	ii++
	// }
	return l
}

func DistinctFloats(floats []float64) []float64 {
	k := make(map[float64]bool)
	for _, v := range floats {
		if _, ok := k[v]; !ok {
			k[v] = true
		}
	}
	l := make([]float64, 0, len(k))
	for v := range k {
		l = append(l, v)
	}
	return l
}
