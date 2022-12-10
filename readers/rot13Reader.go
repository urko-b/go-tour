package readers

import (
	"io"
)

const a = int('a')
const z = int('z')
const A = int('A')
const Z = int('Z')

func rot13(b int) int {

	isLower := b >= a && b <= z
	isUpper := b >= A && b <= Z

	if isLower {
		return a + ((b - a + 13) % 26)
	}

	if isUpper {
		return A + ((b - A + 13) % 26)
	}

	return b
}

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err != io.EOF && err != nil {
		return n, err
	}

	if err == io.EOF {
		return 0, err
	}

	for x := range b {
		b[x] = byte(rot13(int(b[x])))
	}

	return n, err
}
