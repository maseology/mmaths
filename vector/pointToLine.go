// modified from https://stackoverflow.com/questions/27161533/find-the-shortest-distance-between-a-point-and-line-segments-not-line
package vector

// Given a line with coordinates 'start' and 'end' and the
// coordinates of a point 'pnt' the proc returns the shortest
// distance from pnt to the line and the coordinates of the
// nearest point on the line.
//
//  1.   Convert the line segment to a vector ('line_vec').
//  2.   Create a vector connecting start to pnt ('pnt_vec').
//  3.   Find the length of the line vector ('line_len').
//  4.   Convert line_vec to a unit vector ('line_unitvec').
//  5.   Scale pnt_vec by line_len ('pnt_vec_scaled').
//  6.   Get the dot product of line_unitvec and pnt_vec_scaled ('t').
//  7.   Ensure t is in the range 0 to 1.
//  8.   Use t to get the nearest location on the line to the end
//       of vector pnt_vec_scaled ('nearest').
//  9.   Calculate the distance from nearest to pnt_vec_scaled.
//  10.  Translate nearest back to the start/end line.
//
// Malcolm Kesson 16 Dec 2012
// http://www.fundza.com/vectors/point2line/index.html

func PointToLine(pnt, start, end [3]float64) (float64, float64, [3]float64) {
	line_vec := Vector(start, end)
	pnt_vec := Vector(start, pnt)
	line_len := Length(line_vec)
	line_unitvec := Unit(line_vec)
	pnt_vec_scaled := Scale(pnt_vec, 1.0/line_len)
	t := Dot(line_unitvec, pnt_vec_scaled)
	nearest := func() [3]float64 {
		if t < 0. {
			// t = 0.
			return Scale(line_vec, 0.)
		} else if t > 1. {
			// t = 1.
			return Scale(line_vec, 1.)
		}
		return Scale(line_vec, t)
	}()
	dist := Distance(nearest, pnt_vec)
	nearest = Add(nearest, start)
	return dist, t, nearest
}
