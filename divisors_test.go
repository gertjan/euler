package euler

import "testing"

func TestNumberOfDivisors(t *testing.T) {
	var truth map[int]int = map[int]int{
		2:    2,
		3:    2,
		4:    3,
		5:    2,
		6:    4,
		7:    2,
		8:    4,
		9:    3,
		11:   2,
		28:   6,
		6857: 2,
	}

	for in, expected := range truth {
		actual := numberOfDivisors(in)
		if actual != expected {
			t.Errorf("numberOfDivisors(%v): expected %v, got %v\n", in, expected, actual)
		}
	}
}
