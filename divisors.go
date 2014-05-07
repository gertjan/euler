package euler

import (
	"math"
)

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
