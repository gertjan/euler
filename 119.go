package euler

import "fmt"
import "sort"

type Ascending []uint64

func (a Ascending) Len() int {
	return len(a)
}

func (a Ascending) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Ascending) Less(i, j int) bool {
	return a[i] < a[j]
}

func P119() string {
	a := make([]uint64, 0)
	n := 50

	for b := uint64(2); len(a) < n; b++ {
		p := b
		for e := uint64(2); e < 50 && len(a) < n; e++ {
			p *= b
			if sum(digits(p)) == b {
				a = append(a, p)
				//fmt.Println(a)
			}
		}
	}
	sort.Sort(Ascending(a))
	return fmt.Sprintf("%v", a[29])
}

func product(a []uint64) uint64 {
	prod := uint64(1)
	for _, i := range a {
		prod *= i
	}
	return prod
}

func sum(a []uint64) uint64 {
	sum := uint64(0)
	for _, i := range a {
		sum += i
	}
	return sum
}

func digits(i uint64) []uint64 {
	digits := fmt.Sprintf("%v", i)
	result := make([]uint64, len(digits))

	for i, d := range digits {
		result[i] = uint64(d - '0')
	}
	return result
}
