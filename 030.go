package euler

import "fmt"

func P030() string {
	powers := make([]int, 10)
	for i, _ := range powers {
		powers[i] = i * i * i * i * i
	}

	sum := 0
	for i := 2; i < 200000; i++ {
		spd := 0
		for n := i; n > 0; n /= 10 {
			spd += powers[n%10]
		}

		if i == spd {
			sum += i
		}
	}
	return fmt.Sprint(sum)
}
