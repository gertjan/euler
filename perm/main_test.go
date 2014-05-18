package perm

import "testing"

func TestGen(t *testing.T) {
	c := Gen([]int{0, 1, 2})
	for a := range c {
		t.Logf("%v\n", a)
	}
}
