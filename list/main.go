package list

type Element struct {
	next  *Element
	value string
}

func (e *Element) HasNext() bool {
	return e.next != nil
}

func (e *Element) Last() (*Element, bool) {

	if e.HasNext() {
		return e.next.Last()
	}
	return e, true
}

func (e *Element) IsLastButOne() bool {
	if e.HasNext() == false {
		return false
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

func (e *Element) ElementAt(key int) *Element {
	if !e.HasNext() || key == 0 {
		return e
	}

	return e.next.ElementAt(key - 1)
}

func (e *Element) Length() int {
	if e.HasNext() {
		return e.next.Length() + 1
	}
	return 1
}

func (e *Element) Reverse() {
	if e.next.next != nil {
		e.next.Reverse()
	}
	e.next.next = e
	e.next = nil
}
