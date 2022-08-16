package mmaths

import "math"

// LineSegment represents a stright line between two points
type LineSegment struct {
	P0, P1        Point
	b, m, xintcpt float64
}

func (ls *LineSegment) Build() {
	d := ls.P0.X - ls.P1.X
	if d == 0. { // vertical line
		ls.b = -9999.
		ls.m = -9999.
		ls.xintcpt = ls.P0.X
	} else {
		ls.m = (ls.P0.Y - ls.P1.Y) / d
		ls.b = ls.P0.Y - ls.m*ls.P0.X
		ls.xintcpt = (ls.P0.Y - ls.b) / ls.m
	}
}

func (ls *LineSegment) IntersectionX(x float64) *Point {
	if ls.m == -9999. { // vertical line
		return nil
	} else if math.Min(ls.P0.X, ls.P1.X) <= x && math.Max(ls.P0.X, ls.P1.X) >= x {
		return &Point{X: x, Y: ls.m*x + ls.b}
	}
	return nil
}

func (ls *LineSegment) IntersectionY(y float64) *Point {
	if ls.m == -9999. { // vertical line
		return &Point{X: ls.P0.X, Y: y}
	} else if math.Min(ls.P0.Y, ls.P1.Y) <= y && math.Max(ls.P0.Y, ls.P1.Y) >= y && ls.m != 0. {
		return &Point{X: (y - ls.b) / ls.m, Y: y}
	}
	return nil
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
