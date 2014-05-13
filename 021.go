package euler

import "fmt"

func P021() string {
	n := 10000
	sumPairs := 0
	var s, t int
	for i := 4; i < n; i++ {
		s = sumOfDivisors(i)
		t = sumOfDivisors(s)

		if i != s && s != t && i == t {
			sumPairs += i
		}
	}
	return fmt.Sprintf("%v", sumPairs)
}
