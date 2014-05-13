package euler

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func P022(in io.Reader) string {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)

	names := make([]string, 0)
	for sc.Scan() {
		names = append(names, sc.Text())
	}

	sort.Strings(names)

	result := 0
	for i, n := range names {
		result += (i + 1) * alphaValue(n)
	}
	return fmt.Sprintf("%v", result)
}

func alphaValue(s string) int {
	min := byte('A')
	max := byte('Z')
	val := len(s)
	for _, r := range []byte(s) {
		if min <= r && r <= max {
			val += int(r - min)
		}
	}
	return val
}
