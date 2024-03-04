package slice

// Difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []int) []int {
	mb := make(map[int]bool, len(b))
	for _, x := range b {
		mb[x] = true
	}
	o := make([]int, 0, len(a)-len(b))
	for _, x := range a {
		if _, ok := mb[x]; !ok {
			o = append(o, x)
		}
	}
	return o
}
