package euler

import "testing"

func TestP007(t *testing.T) {
	expected := "104743"
	actual := P007()
	if actual != expected {
		t.Errorf("P007: expected %s, got %s\n", expected, actual)
	}
}
