package sudoku

import (
	"bytes"
	"sort"
	"strconv"
)

const dimension = 9

type Sudoku interface {
	Values(x, y int) []int
	Solve()
	Solved() bool
}

func (g grid) Solved() bool {
	solved := true
	for x := 0; x < dimension; x++ {
		for y := 0; y < dimension; y++ {
			if len(g[x][y]) != 1 {
				solved = false
			}
		}
	}
	return solved
}

type grid [][][]int

func (g grid) Values(x, y int) []int {
	return g[x][y]
}

func (g grid) String() string {
	var str bytes.Buffer
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			v := g[i][j]
			if len(v) == 1 {
				str.WriteString(strconv.Itoa(v[0]))
			} else {
				str.WriteRune(rune('0'))
			}
		}
		str.WriteRune('\n')
	}
	return str.String()
}

func (g grid) search() bool {
	g.pruneAll()
	switch {
	case g.Solved():
		return true
	case !g.valid():
		return false
	}

	original := g.String()
	guess := g.makeGuess()

	for _, n := range g[guess.x][guess.y] {
		g.set(original)
		g[guess.x][guess.y] = []int{n}
		found := g.search()
		if found {
			return found
		}
	}

	return false
}

func (g grid) Solve() {
	g.search()
}

func (g grid) valid() bool {
	g.pruneAll()
	for x := 0; x < dimension; x++ {
		for y := 0; y < dimension; y++ {
			if len(g[x][y]) == 0 {
				return false
			}
		}
	}
	return true
}

func (g grid) makeGuess() guess {
	guesses := []guess{}
	for x := 0; x < dimension; x++ {
		for y := 0; y < dimension; y++ {
			if l := len(g[x][y]); l > 1 {
				guesses = append(guesses, guess{x: x, y: y, l: l})
			}
		}
	}

	sort.Sort(bestFirst(guesses))
	return guesses[0]
}

type state struct {
	grid string
	try  int
}

func (g grid) pruneAll() {
	changed := true
	for changed {
		changed = false
		for i := 0; i < dimension; i++ {
			for j := 0; j < dimension; j++ {
				if len(g[i][j]) == 1 {
					changed = changed || g.prune(i, j)
				}
			}
		}
	}
}

func (g grid) prune(i, j int) bool {
	val := g[i][j][0]
	pruned := false
	pr := false

	//row and column
	for x := 0; x < dimension; x++ {
		if x != j {
			g[i][x], pr = prune(g[i][x], val)
			pruned = pruned || pr
		}

		if x != i {
			g[x][j], pr = prune(g[x][j], val)
			pruned = pruned || pr
		}
	}

	//box
	bi := i / 3
	bj := j / 3
	for x := 0; x < dimension; x++ {
		for y := 0; y < dimension; y++ {
			bx := x / 3
			by := y / 3
			if bx == bi && by == bj && !(i == x && j == y) {
				g[x][y], pr = prune(g[x][y], val)
				pruned = pruned || pr
			}
		}
	}

	return pruned
}

func From(s string) Sudoku {
	return fromString(s)
}

func fromString(s string) grid {
	g := make(grid, dimension)
	for i := 0; i < dimension; i++ {
		g[i] = make([][]int, dimension)
	}
	g.set(s)
	return g
}

func (g grid) set(s string) {
	k := 0
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			for s[k] < '0' || s[k] > '9' {
				k++
			}
			d := int(s[k] - '0')
			k++
			if d == 0 {
				g[i][j] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			} else {
				g[i][j] = []int{d}
			}
		}
	}
}

func prune(a []int, n int) ([]int, bool) {
	pruned := false
	for i, x := range a {
		if x == n {
			pruned = true
			a = append(a[:i], a[i+1:]...)
		}
	}
	return a, pruned
}

type guess struct {
	x int
	y int
	l int
}

type bestFirst []guess

func (b bestFirst) Len() int {
	return len(b)
}

func (b bestFirst) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b bestFirst) Less(i, j int) bool {
	return b[i].l < b[j].l
}
