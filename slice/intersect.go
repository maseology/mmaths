package slice

// Intersect returns the intersection of 2 slices
// modified from: https://github.com/juliangruber/go-intersect/blob/master/intersect.go
func Intersect(a, b []int) (set []int) {
	hash := make(map[int]bool, len(a))
	for _, i := range a {
		hash[i] = true
	}
	for _, i := range b {
		if _, ok := hash[i]; ok {
			set = append(set, i)
		}
	}
	return set
}
