package mmaths

import "fmt"

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
