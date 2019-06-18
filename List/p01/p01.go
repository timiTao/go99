package p01

/* p01. Find the last element of a list. */
func lastElement(s []int) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	return s[len(s)-1], true
}
