package euler

import (
	"fmt"

	"github.com/gertjan/euler/primes"
)

func P027() string {
	pr := primes.New()
	limit := 1000
	var max, mp int

	for a := -limit; a < limit; a++ {
		for b := -limit; b < limit; b++ {
			np := 0
			for n := 0; ; n++ {
				if p := n*n + a*n + b; pr.Contains(p) {
					np++
				} else {
					break
				}
			}
			if np > max {
				max = np
				mp = a * b
			}
		}
	}
	return fmt.Sprint(mp)
}
