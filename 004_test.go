package euler

import "testing"

func TestP004(t *testing.T) {
	expected := "906609"
	actual := P004()
	if actual != expected {
		t.Errorf("P004: expected %s, got %s\n", expected, actual)
	}
}
