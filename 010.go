package euler

import (
	"fmt"
)

func P010() string {
	_ = isPrime(2000000)
	sum := 5
	for i := 5; i < 2000000; i += 2 {
		if isPrime(i) {
			sum += i
		}
	}
	return fmt.Sprint(sum)
}
