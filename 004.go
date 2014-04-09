package euler

import (
	"fmt"
)

func P004() string {
	max := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			p := i * j
			if isPalindrome(fmt.Sprintf("%v", p)) {
				if p > max {
					max = p
				}
			}
		}
	}
	return fmt.Sprintf("%v", max)
}
