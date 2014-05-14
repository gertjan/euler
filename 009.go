package euler

import (
	"fmt"
)

func P009() string {
	var p int
	sq := squares(1000 * 1000)
	for c, csq := range sq {
		for b, bsq := range sq {
			if b >= c {
				break
			}
			for a, asq := range sq {
				if a >= b {
					break
				}
				if a+b+c == 1000 && asq+bsq == csq {
					p = a * b * c
				}
			}
		}
	}
	return fmt.Sprint(p)
}
