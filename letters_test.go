package euler

import "testing"

func TestToLetters(t *testing.T) {
	truth := map[int]string{
		1:   "one",
		2:   "two",
		3:   "three",
		4:   "four",
		5:   "five",
		6:   "six",
		7:   "seven",
		8:   "eight",
		9:   "nine",
		10:  "ten",
		11:  "eleven",
		12:  "twelve",
		13:  "thirteen",
		14:  "fourteen",
		15:  "fifteen",
		16:  "sixteen",
		17:  "seventeen",
		18:  "eighteen",
		19:  "nineteen",
		20:  "twenty",
		30:  "thirty",
		40:  "forty",
		50:  "fifty",
		60:  "sixty",
		70:  "seventy",
		80:  "eighty",
		90:  "ninety",
		100: "one hundred",
		856: "eight hundred and fifty-six",
	}

	for i, n := range truth {
		if a := toLetters(i); a != n {
			t.Errorf("toLetters(%v) != %v, got %v", i, n, a)
		}
	}
}
