package fibonacci

import (
	"fmt"
	"testing"
)

func Test_fibonacci(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
