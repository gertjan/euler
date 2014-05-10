package euler

import (
	"bufio"
	"fmt"
	"io"

	"github.com/gertjan/euler/sudoku"
)

func P096(in io.Reader) string {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanLines)
	sum := 0
	for sc.Scan() {
		grid := ""
		for i := 0; i < 9; i++ {
			sc.Scan()
			grid += sc.Text()
		}
		s := sudoku.From(grid)
		s.Solve()

		sum += s.Values(0, 0)[0] * 100
		sum += s.Values(0, 1)[0] * 10
		sum += s.Values(0, 2)[0]
	}
	return fmt.Sprintf("%v", sum)
}
