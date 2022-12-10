package handle_errors

import (
	"fmt"
	"testing"
)

func TestHandleError(t *testing.T) {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
