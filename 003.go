package euler

import (
	"fmt"

	"github.com/gertjan/euler/primes"
)

func P003() string {
	factors := primes.New().Factors(600851475143)
	max := 0
	for _, f := range factors {
		if f > max {
			max = f
		}
	}
	return fmt.Sprint(max)
}
