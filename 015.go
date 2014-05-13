package euler

import "fmt"

func P015() string {
	n := uint64(20)
	result := uint64(1)

	for i := uint64(1); i <= n; i++ {
		result = result * (n + i) / i
	}
	return fmt.Sprintf("%v", result)
}
