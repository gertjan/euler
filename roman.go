package euler

const (
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)

func add(s string, c string, n int) string {
	for i := 0; i < n; i++ {
		s += c
	}
	return s
}

func minRoman(n int) string {
	roman := ""
	rem := n

	roman = add(roman, "M", rem/1000)
	rem %= 1000
	i := rem / 100
	switch i {
	case 4:
		roman += "CD"
	case 5:
		roman += "D"
	case 9:
		roman += "CM"
	case 1, 2, 3:
		roman = add(roman, "C", i)
	case 6, 7, 8:
		roman += "D"
		roman = add(roman, "C", i-5)
	}

	rem %= 100
	i = rem / 10
	switch i {
	case 4:
		roman += "XL"
	case 5:
		roman += "L"
	case 9:
		roman += "XD"
	case 1, 2, 3:
		roman = add(roman, "X", i)
	case 6, 7, 8:
		roman += "L"
		roman = add(roman, "X", i-5)
	}

	i = rem % 10
	switch i {
	case 4:
		roman += "IV"
	case 5:
		roman += "V"
	case 9:
		roman += "IX"
	case 1, 2, 3:
		roman = add(roman, "I", i)
	case 6, 7, 8:
		roman += "V"
		roman = add(roman, "I", i-5)
	}

	return roman
}

func decimal(roman string) int {
	d := 0
	for i, c := range roman {
		switch c {
		case 'I':
			if i+1 < len(roman) && (roman[i+1] == 'V' || roman[i+1] == 'X') {
				d -= 1
			} else {
				d += 1
			}
		case 'X':
			if i+1 < len(roman) && (roman[i+1] == 'L' || roman[i+1] == 'C') {
				d -= 10
			} else {
				d += 10
			}
		case 'C':
			if i+1 < len(roman) && (roman[i+1] == 'D' || roman[i+1] == 'M') {
				d -= 100
			} else {
				d += 100
			}
		case 'V':
			d += V
		case 'L':
			d += L
		case 'D':
			d += D
		case 'M':
			d += M
		}
	}
	return d
}
