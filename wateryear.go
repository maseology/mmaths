package mmaths

import "time"

func Wateryear(dt time.Time) int {
	if dt.Month() < 10 {
		return dt.Year()
	}
	return dt.Year() + 1
}

func Dateday(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
