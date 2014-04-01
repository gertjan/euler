package euler

import "testing"

func TestP003(t *testing.T) {
	expected := "6857"
	actual := P003()
	if actual != expected {
		t.Errorf("P003: expected %s, got %s\n", expected, actual)
	}
}
