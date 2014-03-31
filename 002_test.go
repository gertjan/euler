package euler

import "testing"

func TestP002(t *testing.T) {
	expected := "4613732"
	actual := P002()
	if actual != expected {
		t.Errorf("P002: expected %s, got %s\n", expected, actual)
	}
}
