package euler

import (
	"fmt"
	"math/big"

	"github.com/gertjan/euler/digital"
)

func P020() string {
	n := big.NewInt(1)
	for i := 2; i <= 100; i++ {
		n.Mul(n, big.NewInt(int64(i)))
	}
	return fmt.Sprintf("%v", digital.Sum(n.String()))
}
