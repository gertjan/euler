package sudoku

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestBackTrack(t *testing.T) {
	f, _ := os.Open("testdata/backtrack.txt")
	b, _ := ioutil.ReadAll(f)
	g := fromString(string(b))
	g.Solve()
	if !g.Solved() {
		t.Fail()
	}
}

func TestGrid3(t *testing.T) {
	f, _ := os.Open("testdata/grid3.txt")
	b, _ := ioutil.ReadAll(f)
	g := fromString(string(b))
	g.Solve()
	if !g.Solved() {
		t.Fail()
	}
}
