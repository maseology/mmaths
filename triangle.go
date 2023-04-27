package mmaths

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

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
	return []float64{t.P0.Distance(t.P1), t.P1.Distance(t.P2), t.P2.Distance(t.P0)}
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

func (t *Triangle) Circumcircle() *Circle {
	// find midpoints
	mp1t := t.P0.MidPoint(t.P1)
	mp2t := t.P1.MidPoint(t.P2)
	mp3t := t.P2.MidPoint(t.P0)

	// side slopes
	m1t := (t.P1.Y - t.P0.Y) / (t.P1.X - t.P0.X)
	m2t := (t.P2.Y - t.P1.Y) / (t.P2.X - t.P1.X)
	m3t := (t.P0.Y - t.P2.Y) / (t.P0.X - t.P2.X)

	r, c := 0., Point{}

	// convert to slope of the perpendicular bisectors
	m1, m2 := 0., 0.
	var mp1, mp2 Point
	if !math.IsInf(m1t, 0) && m1t != 0. {
		m1 = -1. / m1t
		mp1 = mp1t
		if !math.IsInf(m2t, 0) && m2t != 0 {
			m2 = -1. / m2t
			mp2 = mp2t
		} else if !math.IsInf(m3t, 0) && m3t != 0 {
			m2 = -1. / m3t
			mp2 = mp3t
		} else {
			goto rightTriangle
		}
	} else {
		if math.IsInf(m2t, 0) || m2t == 0. || math.IsInf(m3t, 0) || m3t == 0. {
			goto rightTriangle
		}
		m1 = -1. / m2t
		m2 = -1. / m3t
		mp1 = mp2t
		mp2 = mp3t
	}

	// 2 equations, 2 unknowns: AX=B --> X=A^-1B
	return func() *Circle {
		mA := mat.NewDense(2, 2, []float64{m1, -1, m2, -1})
		mB := mat.NewDense(2, 1, []float64{m1*mp1.X - mp1.Y, m2*mp2.X - mp2.Y})
		var mM mat.Dense

		if mA.Inverse(mA) != nil {
			goto zerodet
		}
		mM.Mul(mA, mB)
		c = Point{X: mM.At(0, 0), Y: mM.At(1, 0)}
		r = t.P0.Distance(c)
		return &Circle{Centroid: c, Radius: r}

	zerodet: // occurs when 3 points create a straight line (inverse of mA cannot be computed as Det(mA)=0)
		return func() *Circle {
			l0, l1, l2 := t.P0.Distance(t.P1), t.P0.Distance(t.P2), t.P1.Distance(t.P2)
			if l0 > l1 && l0 > l2 {
				r = l0 / 2.0
				c = t.P0.MidPoint(t.P1)
			} else if l1 > l2 {
				r = l1 / 2.0
				c = t.P0.MidPoint(t.P2)
			} else {
				r = l2 / 2.0
				c = t.P1.MidPoint(t.P2)
			}
			return &Circle{Centroid: c, Radius: r}
		}()
	}()

rightTriangle: // right angle triangle
	return func() *Circle {
		l1t := t.P0.Distance(t.P1)
		l2t := t.P1.Distance(t.P2)
		l3t := t.P2.Distance(t.P0)
		if l1t > l2t && l1t > l3t {
			r = l1t / 2.0
			c = mp1t
		} else if l2t > l3t {
			r = l2t / 2.0
			c = mp2t
		} else {
			r = l3t / 2.0
			c = mp3t
		}
		return &Circle{Centroid: c, Radius: r}
	}()
}

func (t *Triangle) MinMaxInteriorAngle() (float64, float64) {
	angle := func(v1, v2, v3 Point) float64 {
		m1 := math.Sqrt(math.Pow(v1.X-v2.X, 2.) + math.Pow(v1.Y-v2.Y, 2.)) // magnitude
		m2 := math.Sqrt(math.Pow(v3.X-v2.X, 2.) + math.Pow(v3.Y-v2.Y, 2.)) // magnitude
		dp := (v1.X-v2.X)*(v3.X-v2.X) + (v1.Y-v2.Y)*(v3.Y-v2.Y)            // dot product
		a := math.Acos(dp / m1 / m2)
		if a > math.Pi || a < 0 {
			panic("Triangle.MinMaxInteriorAngle error")
		}
		return a
	}
	an, ax := math.MaxFloat64, 0.
	v := []Point{t.P0, t.P1, t.P2}
	for n2 := 0; n2 < 3; n2++ {
		n1 := (n2 + 3 - 1) % 3 // note '%' is not a true modulus in Go, it returns the remainder
		n3 := (n2 + 1) % 3
		a := angle(v[n1], v[n2], v[n3])
		if a > ax {
			ax = a
		}
		if a < an {
			an = a
		}
	}
	return an, ax
}
