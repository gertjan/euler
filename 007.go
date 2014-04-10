package euler

import "fmt"

func P007() string {
	nb := 1
	i := 3
	for ; nb < 10001; i += 2 {
		if isPrime(i) {
			nb++
			if nb == 10001 {
				break
			}
		}
	}
	return fmt.Sprintf("%v", i)
}
