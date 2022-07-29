package mmaths

import (
	"math"
)

// CCWEtoCWN converts trig angles (counter-clockwise east) to clockwise north
// from https://math.stackexchange.com/questions/1589793/a-formula-to-convert-a-counter-clockwise-angle-to-clockwise-angle-with-an-offset:
// (−θ+90°) mod 360°: The negative on θ deals with the fact that we are changing from counterclockwise to clockwise.
// The +90° deals with the offset of ninety degrees. And lastly we need to mod by 360° to keep our angle in the desired range [0°,360°]
func CCWEtoCWN(ccwe float64) float64 {
	a := math.Pi/2. - ccwe
	for {
		if a > 0 {
			break
		}
		a += 2. * math.Pi
	}
	return a
}
