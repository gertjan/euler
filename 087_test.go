package euler

import "testing"

func TestP087(t *testing.T) {
	test(t, P087, "1097343")
}

func BenchmarkP087(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P087()
	}
}
