package element

type Element struct {
	next  *Element
	value string
}

func (e *Element) HasNext() bool {
	return e.next != nil
}
