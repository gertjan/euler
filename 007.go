package euler

import (
	"fmt"

	"github.com/gertjan/euler/primes"
)

func P007() string {
	nb := 1
	i := 3
	pr := primes.New()
	for ; nb < 10001; i += 2 {
		if pr.Contains(i) {
			nb++
			if nb == 10001 {
				break
			}
		}
	}
	return fmt.Sprint(i)
}
