// modified from https://stackoverflow.com/questions/27161533/find-the-shortest-distance-between-a-point-and-line-segments-not-line
package vector

import "math"

func Dot(v, w [3]float64) float64 {
	x, y, z := v[0], v[1], v[2]
	X, Y, Z := w[0], w[1], w[2]
	return x*X + y*Y + z*Z
}

func Length(v [3]float64) float64 {
	x, y, z := v[0], v[1], v[2]
	return math.Sqrt(x*x + y*y + z*z)
}

func Vector(b, e [3]float64) [3]float64 {
	x, y, z := b[0], b[1], b[2]
	X, Y, Z := e[0], e[1], e[2]
	return [...]float64{X - x, Y - y, Z - z}
}

func Unit(v [3]float64) [3]float64 {
	x, y, z := v[0], v[1], v[2]
	mag := Length(v)
	return [...]float64{x / mag, y / mag, z / mag}
}

func Distance(p0, p1 [3]float64) float64 {
	return Length(Vector(p0, p1))
}

func Scale(v [3]float64, sc float64) [3]float64 {
	x, y, z := v[0], v[1], v[2]
	return [3]float64{x * sc, y * sc, z * sc}
}

func Add(v, w [3]float64) [3]float64 {
	x, y, z := v[0], v[1], v[2]
	X, Y, Z := w[0], w[1], w[2]
	return [3]float64{x + X, y + Y, z + Z}
}

func Angle2d(b, e [3]float64) float64 {
	v := Vector(b, e)
	return math.Atan(v[1] / v[0]) // returns [-pi/2,pi/2]
}
