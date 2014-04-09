package euler

import "fmt"

func P005() string {
	step := 20 * 19
	r := 0
	for i := step; ; i += step {
		ok := true
		for j := 1; j < 21; j++ {
			if i%j != 0 {
				ok = false
				break
			}
		}
		if ok {
			r = i
			break
		}
	}
	return fmt.Sprintf("%v", r)
}
