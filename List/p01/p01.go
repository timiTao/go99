package p01

/* p01. Find the Last element of a list. */

type Element struct {
	next  *Element
	value string
}

func (e *Element) HasNext() bool {
	return e.next != nil;
}

func (e *Element) Last() (*Element, bool) {
	if e.HasNext() == false {
		return e, true;
	}

	for item := e; true; item = item.next {
		if item.HasNext() == false {
			return item, true
		}
	}
	return nil, false
}
