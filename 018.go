package euler

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func P018(in io.Reader) string {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)

	max := []int{}
	row := []int{}

	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		row = append(row, n)
		if len(row) > len(max) {
			for i, _ := range row {
				switch {
				case i == 0 && len(max) > 0:
					row[i] += max[i]
				case i > 0 && i < len(max):
					row[i] += maxi(max[i-1], max[i])
				case i == len(max) && len(max) > 0:
					row[i] += max[i-1]
				}
			}
			max = row
			row = []int{}
		}
	}
	return fmt.Sprint(maxi(max...))
}

func maxi(n ...int) int {
	max := 0
	for _, i := range n {
		if i > max {
			max = i
		}
	}
	return max
}
