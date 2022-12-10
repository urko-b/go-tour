package word_count

import (
	"testing"

	"golang.org/x/tour/wc"
)

func TestWordCount(t *testing.T) {
	wc.Test(WordCount)

}
