package euler

import "fmt"

func P002() string {
	a, b := 1, 2
	sum := 0
	for b < 4e6 {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	return fmt.Sprintf("%v", sum)
}
