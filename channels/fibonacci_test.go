package channels

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	c := make(chan int64, 100)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func Test_fibonacci_select(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci_select(c, quit)
}
