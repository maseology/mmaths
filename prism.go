package mmaths

import "math"

const tol = .001

// Prism struct represents a singular model prism
type Prism struct {
	Z           []complex128 // complex coordinates
	Top, Bot, A float64
}

// New prism constructor
func (q *Prism) New(z []complex128, top, bot float64) {
	q.Z = z // complex coordinates
	q.Top = top
	q.Bot = bot
	q.A = func() float64 {
		a := 0.
		nfaces := len(q.Z)
		for j := range q.Z {
			jj := (j + 1) % nfaces
			a += real(q.Z[j])*imag(q.Z[jj]) - real(q.Z[jj])*imag(q.Z[j])
		}
		a /= -2. // negative used here because vertices are entered in clockwise order
		if a <= 0. {
			panic("mmaths.Prism area calculation error, may be given in counter-clockwise order")
		}
		return a
	}()
}

// CentroidXY returns the coordinates of the prism centroid
func (q *Prism) Centroid() complex128 {
	sc, c := 0.+0.i, 0
	for _, v := range q.Z {
		sc += v
		c++
	}
	return sc / complex(float64(c), 0.)
}

// CentroidXY returns the coordinates of the prism centroid
func (q *Prism) CentroidXY() (x, y float64) {
	// sc, c := 0.+0.i, 0
	// for _, v := range q.Z {
	// 	sc += v
	// 	c++
	// }
	// ccxy := sc / complex(float64(c), 0.)
	ccxy := q.Centroid()
	return real(ccxy), imag(ccxy)
}

// getExtentsXY returns the XY-extents of the prism
func (q *Prism) getExtentsXY() (yn, yx, xn, xx float64) {
	yn, yx, xn, xx = math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64
	for _, v := range q.Z {
		yn = math.Min(yn, imag(v))
		yx = math.Max(yx, imag(v))
		xn = math.Min(xn, real(v))
		xx = math.Max(xx, real(v))
	}
	return
}

// ContainsXY returns true if the given (x,y) coordinates are contained by the prism planform bounds
func (q *Prism) ContainsXY(x, y float64) bool {
	return PnPolyC(q.Z, complex(x, y), tol)
}

// Contains returns true if the given particle is contained by the prism bounds
func (q *Prism) Contains(x, y, z float64) bool {
	if !PnPolyC(q.Z, complex(x, y), tol) {
		return false
	}
	return z <= q.Top && z >= q.Bot
}
