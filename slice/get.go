package slice

import "sort"

// GetIndxFloat64 performs a Go-native binary tree search on a sorted slice of floats to determine the position x would fit.
// modified from: https://flaviocopes.com/golang-algorithms-binary-search/
func GetIndxFloat64(sorted []float64, x float64) int {
	return sort.Search(len(sorted), func(i int) bool { return sorted[i] >= x })
}
