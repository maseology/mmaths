package mmaths

import (
	"fmt"
	"math"
)

// Matrix alias for a 2d slice
type Matrix [][]float64

// Print matrix
func (mt Matrix) Print() {
	n := len(mt)
	m := len(mt[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%5f", mt[i][j])
			fmt.Print("   ")
		}
		println()
	}
}

// T transpose matrix
func (mt Matrix) T() Matrix {
	r, c := len(mt), len(mt[0])
	var b Matrix = make([][]float64, c)
	for j := 0; j < c; j++ {
		b[j] = make([]float64, r)
		for i := 0; i < r; i++ {
			b[j][i] = mt[i][j]
		}
	}
	return b
}

// Multiply matrices
func (mt Matrix) Multiply(mp Matrix) Matrix {
	nL := len(mt)
	mL := len(mt[0])
	nR := len(mp)
	mR := len(mp[0])
	if nR != mL {
		panic("incompatible matrices for multiplication")
	}
	var b Matrix = make([][]float64, nL)
	for n := 0; n < nL; n++ {
		b[n] = make([]float64, mR)
		for m := 0; m < mR; m++ {
			for k := 0; k < nR; k++ {
				b[n][m] += mt[n][k] * mp[k][m]
			}
		}
	}
	return b
}

// GaussJordanElimination is an algorithm for solving systems of linear equations.
// It is usually understood as a sequence of operations performed on the associated
// matrix of coefficients. This method can also be used to find the rank of a matrix,
// to calculate the determinant of a matrix, and to calculate the inverse of an
// invertible square matrix.
// from: http://www.sanfoundry.com/cpp-program-implement-gauss-jordan-elimination/
func (mt Matrix) GaussJordanElimination() Matrix {
	n := len(mt)
	m := len(mt[0])
	a := make([][]float64, n*2)
	for i := 0; i < n*2; i++ {
		a[i] = make([]float64, m*2)
	}

	// copy matrix values for solving
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a[i][j] = mt[i][j]
		}
		for j := m; j < 2*m; j++ {
			if j == i+m {
				a[i][j] = 1.
			}
		}
	}

	// partial pivoting
	for i := n - 1; i > 0; i-- {
		if a[i-1][1] < a[i][1] {
			for j := 0; j < m*2; j++ {
				d := a[i][j]
				a[i][j] = a[i-1][j]
				a[i-1][j] = d
			}
		}
	}

	// reducing to diagonal  matrix
	for i := 0; i < n; i++ {
		for j := 0; j < m*2; j++ {
			if j != i {
				d := a[j][i] / a[i][i]
				for k := 0; k < m*2; k++ {
					a[j][k] -= a[i][k] * d
				}
			}
		}
	}

	// reducing to unit matrix
	for i := 0; i < n; i++ {
		d := a[i][i]
		for j := 0; j < m*2; j++ {
			a[i][j] = a[i][j] / d
		}
	}

	var b Matrix = make([][]float64, n)
	for i := 0; i < n; i++ {
		b[i] = a[i][m:]
	}
	return b
}

var t bool

// GradientDescent is a first-order iterative optimization algorithm for finding the minimum of a function.
// used to solve for x in Ax=B
// from: https://stackoverflow.com/questions/16422287/linear-regression-library-for-go-language
func GradientDescent(A Matrix, B []float64, n int, alpha float64) []float64 {
	at := A.T()
	x := make([]float64, len(at))
	t = true

	for i := 0; i < n; i++ {
		diffs := calcDiff(x, B, at)
		grad := calcGradient(diffs, at)

		for j := 0; j < len(grad); j++ {
			x[j] += alpha * grad[j]
		}
	}
	return x
}

func calcDiff(x []float64, b []float64, a [][]float64) []float64 {
	diffs := make([]float64, len(b))
	for i := 0; i < len(b); i++ {
		prediction := 0.0
		for j := 0; j < len(x); j++ {
			prediction += x[j] * a[j][i]
		}
		diffs[i] = b[i] - prediction
		if t && math.IsNaN(diffs[i]) {
			fmt.Println("diffs Nan")
			fmt.Println(b[i], prediction)
			fmt.Println(x)
			t = false
		}
		if t && math.IsInf(diffs[i], 0) {
			fmt.Println("diffs Inf")
			fmt.Println(b[i], prediction)
			fmt.Println(x)
			t = false
		}
	}
	return diffs
}

func calcGradient(diffs []float64, a [][]float64) []float64 {
	gradient := make([]float64, len(a))
	for i := 0; i < len(diffs); i++ {
		for j := 0; j < len(a); j++ {
			gradient[j] += diffs[i] * a[j][i]
			if t && math.IsNaN(gradient[j]) {
				fmt.Println("grad1")
				fmt.Println(diffs[i], a[j][i], gradient[j])
				t = false
			}
		}
	}
	for i := 0; i < len(a); i++ {
		g1 := gradient[i]
		gradient[i] = gradient[i] / float64(len(diffs))
		if t && math.IsNaN(gradient[i]) {
			fmt.Println("grad2")
			fmt.Println(g1, gradient[i])
			t = false
		}
	}
	return gradient
}
