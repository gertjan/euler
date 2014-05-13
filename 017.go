package euler

import "fmt"

func P017() string {
	n := 0
	for i := 1; i <= 1000; i++ {
		s := toLetters(i)
		n += countLetters(s)
	}
	return fmt.Sprintf("%v", n)
}
