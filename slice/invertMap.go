package slice

import "sort"

// InvertMap collect slices of keys that share a particular value
func InvertMap(origMap map[int]int) (newMap map[int][]int, sortedkeys []int) {
	newMap, sortedkeys = make(map[int][]int), make([]int, 0)
	for k, v := range origMap {
		if _, ok := newMap[v]; !ok {
			newMap[v] = []int{k}
			sortedkeys = append(sortedkeys, v)
		} else {
			newMap[v] = append(newMap[v], k)
		}
	}
	sort.Ints(sortedkeys)
	return
}
