package primes

import "testing"

func TestContains(t *testing.T) {
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
		27:   false,
		35:   false,
		6857: true,
	}
	pr := New()
	for in, expected := range truth {
		actual := pr.Contains(in)
		if actual != expected {
			t.Errorf("Contains(%v): expected %t, got %t\n", in, expected, actual)
		}
	}
}
