package p02

/* P02 (*) Find the last but one element of a list. */

type Element struct {
	next  *Element
	value string
}

func (e *Element) HasNext() bool {
	return e.next != nil;
}

func (e *Element) IsLastButOne() bool {
	if e.HasNext() == false {
		return false;
	}

	return e.next.HasNext() == false
}

func (e *Element) LastButOne() (*Element, bool) {
	if !e.HasNext() {
		return nil, false
	}

	if e.HasNext() && !e.next.HasNext() {
		return e, true
	}

	return e.next.LastButOne()
}
