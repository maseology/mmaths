package mmaths

import (
	"math"
)

// RoundTo previously "Round" but now included in math after Go v1.10
func RoundTo(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}

// RoundToValue rounds to the nearest multiple of the given roundTo value
func RoundToValue(f, roundTo float64, roundUpDwn int) float64 {
	if roundTo == 0.0 {
		panic("RoundToValue must have a non-zone roundTo value")
	}
	f1 := f / roundTo
	if f1 == float64(int64(f1)) {
		return f
	}
	if roundUpDwn < 0 {
		f1 -= 0.5
	} else if roundUpDwn > 0 {
		f1 += 0.5
	}
	f2 := RoundTo(f1-float64(int64(f1)), 0) + float64(int64(f1))
	return f2 * roundTo
}

// SignificantFigure converts the number to n significant figures
// ex: 1234567 to 1230000 at significant 3 figures
func SignificantFigure(f float64, sigFig int) float64 {
	if f <= 0. {
		return 0.
	}
	f2 := math.Pow10(int(math.Log10(math.Abs(f))))
	return RoundTo(f/f2, sigFig-1) * f2
}

// RoundRange determines the likely range that encompasses the given values
func RoundRange(minVal, maxVal float64) (float64, float64) {
	sig := int(math.Log10(maxVal-minVal)) - 1
	return RoundToValue(minVal, math.Pow10(sig), -1), RoundToValue(maxVal, math.Pow10(sig), 1)
}
