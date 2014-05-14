package euler

import "fmt"

func P072() string {
	//init
	limit := 1000001
	phi := make([]int, limit)
	for i, _ := range phi {
		phi[i] = i
	}

	result := 0
	for i := 2; i < limit; i++ {
		if phi[i] == i {
			for j := i; j < limit; j += i {
				phi[j] = phi[j] / i * (i - 1)
			}
		}
		result += phi[i]
	}
	return fmt.Sprint(result)
}
