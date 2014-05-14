package euler

import "fmt"

func P028() string {
	n := 1
	sum := 1
	for c := 2; c <= 1001; c += 2 {
		for i := 0; i < 4; i++ {
			n += c
			sum += n
		}
	}
	return fmt.Sprint(sum)
}
