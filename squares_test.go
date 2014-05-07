package euler

import "testing"

func TestSquares(t *testing.T) {
	sq := squares(100)
	if len(sq) != 11 {
		t.Fail()
	}
}
