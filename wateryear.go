package mmaths

import "time"

func Wateryear(dt time.Time) int {
	if dt.Month() < 10 {
		return dt.Year()
	}
	return dt.Year() + 1
}
