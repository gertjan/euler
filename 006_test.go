package euler

import "testing"

func TestP006(t *testing.T) {
	expected := "25164150"
	actual := P006()
	if actual != expected {
		t.Errorf("P006: expected %s, got %s\n", expected, actual)
	}
}
