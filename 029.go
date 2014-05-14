package euler

import (
	"fmt"
	"math/big"
)

func P029() string {
	powers := make(map[string]struct{})
	one := big.NewInt(1)
	max := big.NewInt(100)

	for a := big.NewInt(2); a.Cmp(max) <= 0; a.Add(a, one) {
		for b := big.NewInt(2); b.Cmp(max) <= 0; b.Add(b, one) {
			var c big.Int
			c.Exp(a, b, nil)
			powers[c.String()] = struct{}{}
		}
	}

	return fmt.Sprint(len(powers))
}
