package euler

import (
	"encoding/csv"
	"fmt"
	"io"
	//"sort"
	"strconv"
)

func P082(in io.Reader) string {
	r := csv.NewReader(in)
	records, _ := r.ReadAll()
	dim := len(records)

	minCost := 0
	tab := make([][]int, dim)
	for i := 0; i < dim; i++ {
		tab[i] = make([]int, dim)
		for j := 0; j < dim; j++ {
			v, _ := strconv.Atoi(records[i][j])
			tab[i][j] = v
			if v < minCost {
				minCost = v
			}
		}
	}

	current := make([]int, dim)
	for i := 0; i < dim; i++ {
		current[i] = tab[i][0]
	}

	for col := 1; col < dim; col++ {
		//go one col to the right
		for i, path := range current {
			path += tab[i][col]
			current[i] = path
		}
		//optimize
		for changed := true; changed; {
			changed = false
			for i, path := range current {
				if up := i - 1; up >= 0 && current[up] > tab[up][col]+path {
					changed = true
					current[up] = tab[up][col] + path
				}

				if down := i + 1; down < dim && current[down] > tab[down][col]+path {
					changed = true
					current[down] = tab[down][col] + path
				}
			}
		}
	}

	minPath := current[0]
	for _, cp := range current {
		if cp < minPath {
			minPath = cp
		}
	}
	return fmt.Sprint(minPath)
}
