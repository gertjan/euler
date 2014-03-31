package euler

import "testing"

func TestP001(t *testing.T) {
	expected := "233168"
	actual := P001()
	if actual != expected {
		t.Errorf("P001: expected %s, got %s\n", expected, actual)
	}
}
