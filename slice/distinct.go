package slice

// Distinct returns a unique slice
func Distinct(ints []int) []int {
	k := make(map[int]bool)
	for _, i := range ints {
		if _, ok := k[i]; !ok {
			k[i] = true
		}
	}
	l, ii := make([]int, len(k)), 0
	for i := range k {
		l[ii] = i
		ii++
	}
	return l
}
