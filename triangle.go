package mmaths

import "math"

type Triangle struct{ P0, P1, P2 Point }

func (t *Triangle) New(p0, p1, p2 []float64) {
	t.P0 = Point{p0[0], p0[1], 0., 0.}
	t.P1 = Point{p1[0], p1[1], 0., 0.}
	t.P2 = Point{p2[0], p2[1], 0., 0.}
}

func (t *Triangle) Area() float64 {
	return math.Abs(t.P0.X*(t.P1.Y-t.P2.Y)+t.P1.X*(t.P2.Y-t.P0.Y)+t.P2.X*(t.P0.Y-t.P1.Y)) / 2.
}

func (t *Triangle) EdgeLengths() []float64 {
	return []float64{t.P0.Distance(&t.P1), t.P1.Distance(&t.P2), t.P2.Distance(&t.P0)}
}

func (t *Triangle) Contains(x, y float64) bool {
	pnt := Point{x, y, 0., 0.}
	vector := func(b, e Point) Point {
		x, y := b.X, b.Y
		X, Y := e.X, e.Y
		return Point{X - x, Y - y, 0., 0.}
	}
	dot := func(v, w Point) float64 {
		x, y := v.X, v.Y
		X, Y := w.X, w.Y
		return x*X + y*Y
	}

	v0, v1, v2 := vector(t.P0, t.P2), vector(t.P0, t.P1), vector(t.P0, pnt)
	dot00, dot01, dot02, dot11, dot12 := dot(v0, v0), dot(v0, v1), dot(v0, v2), dot(v1, v1), dot(v1, v2)
	dbDenom := dot00*dot11 - dot01*dot01
	u, v := (dot11*dot02-dot01*dot12)/dbDenom, (dot00*dot12-dot01*dot02)/dbDenom
	return (u >= 0) && (v >= 0) && (u+v < 1)
}

func (t *Triangle) BarycentricWeights(x, y float64) []float64 {
	denom := (t.P1.Y-t.P2.Y)*(t.P0.X-t.P2.X) + (t.P2.X-t.P1.X)*(t.P0.Y-t.P2.Y)
	w0 := ((t.P1.Y-t.P2.Y)*(x-t.P2.X) + (t.P2.X-t.P1.X)*(y-t.P2.Y))
	w1 := ((t.P2.Y-t.P0.Y)*(x-t.P2.X) + (t.P0.X-t.P2.X)*(y-t.P2.Y))
	w0 /= denom
	w1 /= denom
	w2 := 1.0 - w1 - w0
	if w0 < 0 || w1 < 0 || w2 < 0 {
		return nil // point lies outside of triangle
	}
	return []float64{w0, w1, w2}
}
