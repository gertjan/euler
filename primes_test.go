package euler

import "testing"

var truth map[int]bool = map[int]bool{
	2:    true,
	3:    true,
	4:    false,
	5:    true,
	6:    false,
	7:    true,
	8:    false,
	9:    false,
	11:   true,
	6857: true,
}

func TestIsPrime(t *testing.T) {
	for in, expected := range truth {
		actual := isPrime(in)
		if actual != expected {
			t.Errorf("isPrime(%v): expected %t, got %t\n", in, expected, actual)
		}
	}
}
