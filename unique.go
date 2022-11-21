package mmaths

// UniqueInts returns a unique subset of the int slice provided.
// from: https://kylewbanks.com/blog/creating-unique-slices-in-go
func UniqueInts(input []int) []int {
	m := make(map[int]int)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val]++
		}
	}
	u := make([]int, 0, len(m))
	for c := range m {
		u = append(u, c)
	}
	return u
}
