package euler

import "testing"

func TestIsPandigital(t *testing.T) {
	r := isPandigital(123456789)
	if !r {
		t.Fail()
	}
}

func TestP032(t *testing.T) {
	test(t, P032, "45228")
}
