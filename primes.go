package euler

import "math"
import "sort"

var primes []int = []int{2, 3, 5}

//isPrime returns whether n is prime.
func isPrime(n int) bool {
	i := sort.SearchInts(primes, n)
	if i < len(primes) {
		return primes[i] == n
	}

	for np := primes[len(primes)-1] + 2; np <= n; np += 2 {
		ip := true
		root := int(math.Sqrt(float64(np)))
		for _, p := range primes {
			if np%p == 0 {
				ip = false
			}

			if p > root {
				break
			}
		}

		if ip {
			primes = append(primes, np)
		}
	}

	return primes[len(primes)-1] == n
}

//factors returns the prime factors of n.
func factors(n int) []int {
	root := int(math.Sqrt(float64(n)))
	_ = isPrime(root)
	factors := []int{}

	rem := n
	for _, p := range primes {
		for rem%p == 0 {
			rem /= p
			factors = append(factors, p)
		}

		if rem == 1 {
			break
		}
	}

	return factors
}

func primesUntil(limit int) []int {
	isPrime(limit)
	idx := sort.SearchInts(primes, limit)
	return primes[:idx]
}
