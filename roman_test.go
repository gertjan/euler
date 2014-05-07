package euler

import "testing"

func TestRoman(t *testing.T) {
	truth := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		10: "X",
		49: "XLIX",
	}
	for d, r := range truth {
		mr := minRoman(d)
		if mr != r {
			t.Errorf("%v in romans should be: %v, not %v\n", d, r, mr)
		}
	}
}

func TestDecimal(t *testing.T) {
	truth := map[string]int{
		"I":    1,
		"XLIX": 49,
	}
	for r, d := range truth {
		a := decimal(r)
		if a != d {
			t.Errorf("%v in decimals should be: %v, not %v\n", r, d, a)
		}
	}
}
