package euler

import (
	"math"
)

func gcd(a, b int) int {
	var r int
	for b != 0 {
		r = a % b
		a = b
		b = r
	}
	return a
}

func sumOfDivisors(n int) int {
	if n == 1 {
		return 0
	}

	sum := 1
	r := int(math.Sqrt(float64(n)))
	for i := 2; i <= r; i++ {
		if n%i == 0 {
			sum += i
			if other := n / i; i < other {
				sum += other
			}
		}
	}
	return sum
}

func numberOfDivisors(n int) int {
	nb := 2
	r := int(math.Sqrt(float64(n)))
	for i := 2; i <= r; i++ {
		if n%i == 0 {
			nb++
			if i < n/i {
				nb++
			}
		}
	}
	return nb
}
