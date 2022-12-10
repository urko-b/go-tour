package word_count

import (
	"strings"
)

func WordCount(s string) map[string]int {
	mp := make(map[string]int)
	for _, str := range strings.Fields(s) {
		mp[str]++
	}
	return mp
}
