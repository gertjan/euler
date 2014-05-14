package euler

import (
	"fmt"
)

func P012() string {
	c, kill := triangles()
	var n int
	for n = range c {
		nd := numberOfDivisors(n)
		if nd > 500 {
			break
		}
	}
	kill <- true
	return fmt.Sprint(n)
}

func triangles() (<-chan int, chan<- bool) {
	c, kill := make(chan int), make(chan bool)
	go func() {
		i := 1
		n := 1
		for {
			select {
			case <-kill:
				close(c)
				return
			case c <- n:

				n, i = n+i+1, i+1
			}
		}
	}()
	return c, kill
}
