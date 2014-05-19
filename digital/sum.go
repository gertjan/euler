package digital

const ZERO = byte('0')

func Sum(n string) uint64 {
	var s uint64
	for _, d := range []byte(n) {
		s += uint64(d - ZERO)
	}
	return s
}
