package euler

import (
	"fmt"
	"math"
)

func P087() string {
	limit := 50000000
	root := int(math.Sqrt(float64(limit)))
	p := primesUntil(root)
	pow := make([][]int, len(p))
	for i, prime := range p {
		p2 := prime * prime
		p3 := p2 * prime
		p4 := p3 * prime
		pow[i] = []int{p2, p3, p4}
	}

	empty := struct{}{}
	numbers := make(map[int]struct{})

	for a := range p {
		for b := range p {
			for c := range p {
				value := pow[a][0] + pow[b][1] + pow[c][2]
				if value < limit {
					numbers[value] = empty
				}
			}
		}
	}
	return fmt.Sprintf("%v", len(numbers))
}
