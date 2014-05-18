package euler

import (
	"fmt"
)

func P027() string {
	limit := 1000
	var max, mp int

	for a := -limit; a < limit; a++ {
		for b := -limit; b < limit; b++ {
			np := 0
			for n := 0; ; n++ {
				if p := n*n + a*n + b; isPrime(p) {
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
