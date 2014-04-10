package euler

import (
	"fmt"
)

func P008(digits []int) string {
	maxproduct := 0
	for i := 0; i < (len(digits) - 4); i++ {
		p := prod(digits[i : i+5])
		if p > maxproduct {
			maxproduct = p
		}
	}
	return fmt.Sprintf("%v", maxproduct)
}

func prod(digits []int) int {
	prod := 1
	for _, d := range digits {
		prod *= d
	}
	return prod
}
