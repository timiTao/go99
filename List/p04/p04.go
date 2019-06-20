package p04

/* p04. Find the number of elements of a list. */

type Element struct {
	next  *Element
	value string
}

func (e *Element) HasNext() bool {
	return e.next != nil
}

func (e *Element) Length() int {
	if e.HasNext() {
		return e.next.Length() + 1
	}
	return 1
}
