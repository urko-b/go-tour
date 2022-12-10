package fibonacci

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	v, v1, v2 := 0, 0, 0

	return func() int {
		v2 = v1
		v1 = v
		v = v1 + v2
		if v <= 0 {
			v1 = 1
			return v
		}
		return v
	}
}
