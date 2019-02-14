package mmaths

// UniqueInts returns a unique subset of the int slice provided.
// from: https://kylewbanks.com/blog/creating-unique-slices-in-go
func UniqueInts(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
