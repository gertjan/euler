package euler

import (
	"fmt"
	"math/big"
)

func P025() string {
	a, b := big.NewInt(1), big.NewInt(1)
	i := 1
	for len(a.String()) < 1000 {
		a, b = b, a.Add(a, b)
		i++
	}
	return fmt.Sprint(i)
}
