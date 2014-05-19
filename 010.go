package euler

import (
	"fmt"

	"github.com/gertjan/euler/primes"
)

func P010() string {
	sum := 0
	for _, p := range primes.New().Slice(2000000) {
		sum += p
	}
	return fmt.Sprint(sum)
}
