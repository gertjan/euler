package euler

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func P011(in io.Reader) string {
	grid := readGrid(in)
	var max uint64

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if i < 16 {
				prod := product(grid[i][j], grid[i+1][j], grid[i+2][j], grid[i+3][j])
				if prod > max {
					max = prod
				}
			}
			if j < 16 {
				prod := product(grid[i][j], grid[i][j+1], grid[i][j+2], grid[i][j+3])
				if prod > max {
					max = prod
				}
			}
			if i < 16 && j < 16 {
				prod := product(grid[i][j], grid[i+1][j+1], grid[i+2][j+2], grid[i+3][j+3])
				if prod > max {
					max = prod
				}
			}
			if i > 2 && j < 16 {
				prod := product(grid[i][j], grid[i-1][j+1], grid[i-2][j+2], grid[i-3][j+3])
				if prod > max {
					max = prod
				}
			}
		}
	}
	return fmt.Sprintf("%v", max)
}

func readGrid(in io.Reader) [][]uint64 {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)

	grid := make([][]uint64, 20)
	for i := 0; i < 20; i++ {
		grid[i] = make([]uint64, 20)
		for j := 0; j < 20; j++ {
			sc.Scan()
			val, _ := strconv.ParseUint(sc.Text(), 10, 64)
			grid[i][j] = val
		}
	}

	return grid
}
