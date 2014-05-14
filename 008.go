package euler

import (
	"fmt"
	"io"
	"io/ioutil"
)

func P008(in io.Reader) string {
	buf, _ := ioutil.ReadAll(in)
	digits := make([]int, len(buf))
	zero := byte('0')

	for _, b := range buf {
		d := int(b - zero)
		if d >= 0 && d < 10 {
			digits = append(digits, d)
		}
	}

	maxproduct := 0
	for i := 0; i < (len(digits) - 4); i++ {
		p := prod(digits[i : i+5])
		if p > maxproduct {
			maxproduct = p
		}
	}
	return fmt.Sprint(maxproduct)
}

func prod(digits []int) int {
	prod := 1
	for _, d := range digits {
		prod *= d
	}
	return prod
}
