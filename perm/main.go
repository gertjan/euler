package perm

import "github.com/cznic/mathutil"

func Gen(a []int) <-chan []int {
	c := make(chan []int)
	go func() {
		mathutil.PermutationFirst(numbers(a))
		send(c, a)
		for mathutil.PermutationNext(numbers(a)) {
			send(c, a)
		}
		close(c)
	}()
	return c
}

func send(c chan []int, a []int) {
	cp := make([]int, len(a))
	copy(cp, a)
	c <- cp
}

type numbers []int

func (n numbers) Len() int           { return len(n) }
func (n numbers) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n numbers) Less(i, j int) bool { return n[i] < n[j] }
