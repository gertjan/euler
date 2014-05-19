package euler

import "fmt"

func P031() string {
	target := 200
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	ways := make([]int, target+1)
	ways[0] = 1
	for i := range coins {
		for j := coins[i]; j <= target; j++ {
			ways[j] += ways[j - coins[i]]
		}
	}
	return fmt.Sprint(ways[target])
}
