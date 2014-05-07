package euler

import (
	"fmt"
)

func P014() string {
	length := map[int]int{
		1: 1,
		2: 2,
	}

	max := 0
	maxn := 0

	for n := 3; n < 1000000; n++ {
		chain := []int{n}
		next := nextCollatz(n)

		for _, f := length[next]; !f; _, f = length[next] {
			chain = append(chain, next)
			next = nextCollatz(next)
		}

		l := length[next]
		for i, n := range chain {
			length[n] = l + len(chain) - i
			if length[n] > max {
				maxn = n
				max = length[n]
			}
		}
	}
	return fmt.Sprintf("%v", maxn)
}

func nextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
