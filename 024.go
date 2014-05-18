package euler

import (
	"fmt"

	"github.com/gertjan/euler/perm"
)

func P024() string {
	c := perm.Gen([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	n := 0
	i := 0
	for p := range c {
		i++
		if i == 1000000 {
			for _, i := range p {
				n = n*10 + i
			}
		}
	}
	return fmt.Sprint(n)
}
