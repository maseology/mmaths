package timeslice

import (
	"sort"
	"time"
)

func Sort(tims []time.Time) []time.Time {
	sort.Slice(tims, func(i, j int) bool {
		return tims[i].Before(tims[j])
	})
	return tims
}

// OR Define us a type so we can sort it
// type TimeSlice []time.Time

// // Forward request for length
// func (p TimeSlice) Len() int {
// 	return len(p)
// }

// // Define compare
// func (p TimeSlice) Less(i, j int) bool {
// 	return p[i].Before(p[j])
// }

// // Define swap over an array
// func (p TimeSlice) Swap(i, j int) {
// 	p[i], p[j] = p[j], p[i]
// }

// // THEN Wrap array in type for sorting:   sort.Sort(TimeSlice([]time.Time{}]))
