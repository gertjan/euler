package euler

import (
	"fmt"

	"github.com/gertjan/euler/perm"
)

// 5 sets of 3
// set[i][2] == set[i+1][1]
// result:
// sum(r[0:3]) == sum(r[3:6]) == sum(r[6:9]) == sum(r[9:12]) == sum(r[12:15])
// r[2] == r[4] && r[5] == r[7] && r[8] == r[10] && r[11] == r[13] && r[14] == r[1]
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4
// 0 1 2 3 2 4 5 4 6 7 6 8 9 8 1
//6531031914842725
//0 1 2 3  4 5 6 7 8 9
//6 5 3 10 1 9 4 8 2 7
func P068() string {
	max := 0
	c := perm.Gen([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	for a := range c {
		if s := sumInt(a[:3]); s == sumInt(a[2:5]) && s == sumInt(a[4:7]) && s == sumInt(a[6:9]) && s == sumInt(append(a[8:], a[1])) {
			if a[0] < a[3] && a[0] < a[5] && a[0] < a[7] && a[0] < a[9] {
				b := make([]int, 15)
				for i, _ := range b {
					switch i {
					case 0, 1, 2, 3:
						b[i] = a[i]
					case 5, 6:
						b[i] = a[i-1]
					case 4, 8, 9:
						b[i] = a[i-2]
					case 7, 11, 12:
						b[i] = a[i-3]
					case 10:
						b[i] = a[i-4]
					case 13:
						b[i] = a[i-5]
					case 14:
						b[i] = a[1]
					}
				}

				number := 0
				for _, x := range b {
					if x == 10 {
						number = number*100 + x
					} else {
						number = number*10 + x
					}
				}
				if max < number && len(fmt.Sprint(number)) == 16 {
					max = number
				}
			}
		}
	}
	return fmt.Sprint(max)
}

func asString(a []int) string {
	i := 0
	for _, b := range a {
		i = i*10 + b
	}
	return fmt.Sprint(i)
}

func sumInt(a []int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}
