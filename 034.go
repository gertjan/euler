package euler

import "fmt"

func P034() string {
	fact := map[int]int{
		0: 1,
		1: 1,
		2: 2,
		3: 6,
		4: 24,
		5: 120,
		6: 720,
		7: 5040,
		8: 40320,
		9: 362880,
	}

	var i, tmp, sum, result int
	for i = 3; i < 50000; i++ {
		tmp = i
		sum = 0
		for tmp > 0 {
			sum += fact[tmp%10]
			tmp /= 10
		}

		if sum == i {
			result += i
		}
	}
	return fmt.Sprint(result)
}
