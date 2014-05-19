package primes

import (
	"math"
	"sort"
)

type Primes interface {
	Contains(n int) bool
	Factors(n int) []int
	Slice(n int) []int
}

type cachedPrimes struct {
	nextToCheck int
	primes      []int
}

func New() Primes {
	return &cachedPrimes{
		nextToCheck: 11,
		primes:      []int{2, 3, 5, 7},
	}
}

//Contains returns whether the collection of primes contains n.
func (cp *cachedPrimes) Contains(n int) bool {
	var next, root int
	var prime bool

	for n >= cp.nextToCheck {
		next = cp.nextToCheck
		prime = true
		root = int(math.Sqrt(float64(next)))

		for _, p := range cp.primes {
			if next%p == 0 {
				prime = false
				break
			}

			if p > root {
				break
			}
		}

		if prime {
			cp.primes = append(cp.primes, next)
		}
		cp.nextToCheck += 2
	}

	idx := sort.SearchInts(cp.primes, n)
	return idx < len(cp.primes) && cp.primes[idx] == n
}

//Factors returns the prime factors of n.
func (cp *cachedPrimes) Factors(n int) []int {
	root := int(math.Sqrt(float64(n)))
	cp.Contains(root)
	factors := []int{}

	rem := n
	for _, p := range cp.primes {
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

//Slice returns all prime number less than the limit.
func (cp *cachedPrimes) Slice(limit int) []int {
	cp.Contains(limit)
	idx := sort.SearchInts(cp.primes, limit)
	return cp.primes[:idx]
}
