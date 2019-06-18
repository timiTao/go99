package p03

/* p03. Find the K'th element of a list. */
func elementAt(myList []int, key int) (int, bool) {
	if key <= 0 {
		return 0, false
	}
	if (key >= len(myList)) {
		return 0, false
	}

	val := myList[key-1]

	return val, true
}