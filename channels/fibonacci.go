package channels

func fibonacci(n int, c chan int64) {
	x, y := int64(0), int64(1)
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
