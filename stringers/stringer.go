package stringers

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (r IPAddr) String() string {
	output := make([]string, 0)
	for _, v := range r {
		output = append(output, fmt.Sprintf("%d", v))
	}
	return strings.Join(output, ".")
}
