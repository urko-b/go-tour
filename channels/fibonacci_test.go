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
