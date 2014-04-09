package euler

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var truth map[string]bool = map[string]bool{
		"aba":      true,
		"aa":       true,
		"abc":      false,
		"abcddcba": true,
		"abcdcba":  true,
		"abcdecba": false,
	}

	for in, expected := range truth {
		actual := isPalindrome(in)
		if actual != expected {
			t.Errorf("isPalindrome(%v): expected %t, got %t\n", in, expected, actual)
		}
	}
}
