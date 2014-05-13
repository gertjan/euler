package euler

import "fmt"

func P023() string {
	n := 28123
	sieve := make([]bool, n+1)
	abundantNumbers := []int{}

	for i := 12; i < n; i++ {
		if sumOfDivisors(i) > i {
			abundantNumbers = append(abundantNumbers, i)
			for _, o := range abundantNumbers {
				sum := i + o
				if sum < len(sieve) {
					sieve[sum] = true
				}
			}
		}
	}

	sum := 0
	for i, b := range sieve {
		if !b {
			sum += i
		}
	}
	return fmt.Sprintf("%v", sum)
}
