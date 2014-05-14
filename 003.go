package euler

import "fmt"

func P003() string {
	factors := factors(600851475143)
	max := 0
	for _, f := range factors {
		if f > max {
			max = f
		}
	}
	return fmt.Sprint(max)
}
