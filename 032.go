package euler

import "fmt"

func isPandigital(p int64) bool {
	var digits, count, tmp int

	for p > 0 {
		tmp = digits
		digits = digits | 1<<uint((p%10)-1)
		if tmp == digits {
			return false
		}

		count++
		p /= 10
	}
	return digits == 511
}

func join(a, b int64) int64 {
	c := b
	for c > 0 {
		a *= 10
		c /= 10
	}
	return a + b
}

func P032() string {
	products := make(map[int64]bool)
	var sum, k, p, i, j int64
	var jmin, jmax int64

	for i = 2; i < 100; i++ {
		if i > 9 {
			jmin = 123
		} else {
			jmin = 1234
		}
		jmax = 10000/i + 1

		for j = jmin; j < jmax; j++ {
			k = i * j
			p = join(join(i, j), k)
			if isPandigital(p) {
				products[k] = true
			}
		}
	}

	for p = range products {
		sum += p
	}
	return fmt.Sprint(sum)
}
