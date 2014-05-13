package euler

import (
	"fmt"
	"math/big"

	"github.com/gertjan/euler/digital"
)

func P016() string {
	n := new(big.Int)
	n.SetBit(n, 1000, 1)
	return fmt.Sprintf("%v", digital.Sum(n.String()))
}
