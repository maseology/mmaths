package timeslice

import "time"

// Distinct returns a unique slice of time
func Distinct(tims []time.Time) []time.Time {
	k := make(map[time.Time]bool)
	for _, i := range tims {
		if _, ok := k[i]; !ok {
			k[i] = true
		}
	}
	l, ii := make([]time.Time, len(k)), 0
	for i := range k {
		l[ii] = i
		ii++
	}
	return l
}
