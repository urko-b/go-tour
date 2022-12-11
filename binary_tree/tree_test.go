package binary_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tour/tree"
	"testing"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)

	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("value: %d\n", v)
	}
}

func TestSame(t *testing.T) {
	assert.Equal(t, Same(tree.New(1), tree.New(1)), true)
	assert.Equal(t, Same(tree.New(1), tree.New(2)), false)
}
