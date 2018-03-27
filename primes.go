package maths

// Primes : brute-force finds the first k prime numbers
func Primes(k int) []int {
	if k < 1 {
		panic("primes() input error")
	}
	p := make([]int, 0, k)
	for j := 1; j < k+1; j++ {
		i := 1
		c := 0
	loop:
		i++
		for d := 2; d <= i/2; d++ {
			if i%d == 0 {
				goto loop
			}
		}
		c++
		if c < j {
			goto loop
		}
		p = append(p, i)
	}
	return p
}
