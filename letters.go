package euler

import (
	"unicode"
)

var numberToLetters = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "one hundred",
	1000: "one thousand",
}

func toLetters(n int) string {
	if l, found := numberToLetters[n]; found {
		return l
	}

	if n > 100 {
		first := toLetters(n/100) + " hundred"
		if rem := n % 100; rem != 0 {
			first += " and " + toLetters(rem)
		}
		return first
	}

	if n > 20 {
		first := toLetters((n / 10) * 10)
		if rem := n % 10; rem != 0 {
			first += "-" + toLetters(rem)
		}
		return first
	}

	return ""
}

func countLetters(s string) int {
	n := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			n++
		}
	}
	return n
}
