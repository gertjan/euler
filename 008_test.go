package euler

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestP008(t *testing.T) {
	expected := "40824"

	in, _ := os.Open("testdata/008.txt")
	buf, _ := ioutil.ReadAll(in)
	digits := make([]int, len(buf))
	zero := byte('0')

	for _, b := range buf {
		d := int(b - zero)
		if d >= 0 && d < 10 {
			digits = append(digits, d)
		}
	}

	actual := P008(digits)
	if actual != expected {
		t.Errorf("P008: expected %s, got %s\n", expected, actual)
	}
}
