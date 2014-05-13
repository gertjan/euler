package euler

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
)

func P013(in io.Reader) string {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)

	sum := big.NewInt(1)
	numb := new(big.Int)
	for sc.Scan() {
		numb.SetString(sc.Text(), 10)
		sum.Add(sum, numb)
	}

	digits := []byte(sum.String())
	return fmt.Sprintf("%v", string(digits[:10]))
}
