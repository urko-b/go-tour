package readers

import (
	"golang.org/x/tour/reader"
	"testing"
)

func TestReader(t *testing.T) {
	reader.Validate(MyReader{})
}
