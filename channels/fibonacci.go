package channels

import "fmt"

// Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
// Channels aren't like files; you don't usually need to close them.
// Closing is only necessary when the receiver must be told there are no more values coming,
// such as to terminate a range loop.
func fibonacci(n int, c chan int64) {
	x, y := int64(0), int64(1)
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// The select statement lets a goroutine wait on multiple communication operations.
//
// A select blocks until one of its cases can run, then it executes that case.
// It chooses one at random if multiple are ready.
func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
