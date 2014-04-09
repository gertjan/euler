package euler

import "testing"

func TestP005(t *testing.T) {
	expected := "232792560"
	actual := P005()
	if actual != expected {
		t.Errorf("P005: expected %s, got %s\n", expected, actual)
	}
}
