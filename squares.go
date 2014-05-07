package euler

func squares(max int) []int {
	sq := []int{1}
	i := 1
	s := 3
	for i <= max {
		sq = append(sq, i)
		i = i + s
		s += 2
	}
	return sq
}
