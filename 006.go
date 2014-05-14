package euler

import "fmt"

func P006() string {
	limit := 100
	sum := limit * (limit + 1) / 2
	ssq := (2*limit + 1) * (limit + 1) * limit / 6
	return fmt.Sprint(sum*sum - ssq)
}
