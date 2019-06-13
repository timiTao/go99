package P01

/* P01. Find the last element of a list. */
func lastElement(s []int) (int, bool)  {
	if len(s) == 0 {
		return 0, false
	} else {
		return s[len(s)-1], true
	}
}