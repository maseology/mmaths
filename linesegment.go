package mmaths

import (
	"math"
)

// LineSegment represents a stright line between two points
type LineSegment struct {
	P0, P1 Point
}

// Intersection2D returns the 2D intersection of two line segments. Returns nil if lines do not intersect.
func (l0 *LineSegment) Intersection2D(l1 *LineSegment) (Point, float64) {
	// first degree Bézier parameter
	d := ((l0.P0.X-l0.P1.X)*(l1.P0.Y-l1.P1.Y) - (l0.P0.Y-l0.P1.Y)*(l1.P0.X-l1.P1.X))
	t := ((l0.P0.X-l1.P0.X)*(l1.P0.Y-l1.P1.Y) - (l0.P0.Y-l1.P0.Y)*(l1.P0.X-l1.P1.X))
	u := ((l0.P0.X-l0.P1.X)*(l0.P0.Y-l1.P0.Y) - (l0.P0.Y-l0.P1.Y)*(l0.P0.X-l1.P0.X))
	t /= d
	u /= -d
	var p Point
	if t >= 0. && t <= 1. && u >= 0. && u <= 1. {
		p.X = l0.P0.X + t*(l0.P1.X-l0.P0.X)
		p.Y = l0.P0.Y + t*(l0.P1.Y-l0.P0.Y)
		return p, t
	}
	return p, math.NaN() //line segments do not intersect
}
