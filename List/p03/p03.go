package p03

/* p03. Find the K'th element of a list. */

type Element struct {
	next  *Element
	value string
}

func (e *Element) HasNext() bool {
	return e.next != nil
}

func (e *Element) ElementAt(key int) *Element {
	if !e.HasNext() || key == 0 {
		return e
	}

	return e.next.ElementAt(key - 1)
}
