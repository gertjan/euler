package euler

import (
	"bufio"
	"fmt"
	"io"
)

func P089(in io.Reader) string {
	s := bufio.NewScanner(in)
	s.Split(bufio.ScanWords)

	reduced := 0
	for s.Scan() {
		roman := s.Text()
		d := decimal(roman)
		mr := minRoman(d)
		reduced += (len(roman) - len(mr))
	}
	return fmt.Sprint(reduced)
}
