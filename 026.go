package euler

import "fmt"

func P026() string {
	imax := 0
	max := 0
	for i := 2; i < 1000; i++ {
		rseen := make(map[int]int)
		k := 0
		for j := 10; j != 0; j *= 10 {
			j = j % i
			k++
			if pj, found := rseen[j]; found {
				if nmax := k - pj; nmax > max {
					imax = i
					max = nmax
				}
				break
			} else {
				rseen[j] = k
			}
		}
	}
	return fmt.Sprint(imax)
}
