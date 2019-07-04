package list

type Element struct {
	next  *Element
	value interface{}
}

func (e *Element) HasNext() bool {
	return e.next != nil
}

/* P01. Find the last element of a list. */
func (e *Element) Last() *Element {

	if e.HasNext() {
		return e.next.Last()
	}
	return e
}

/* P02. Find the last but one element of a list. */
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

/* P03. Find the K'th element of a list. */
func (e *Element) ElementAt(key int) *Element {
	if !e.HasNext() || key == 0 {
		return e
	}

	return e.next.ElementAt(key - 1)
}

/* P04. Find the number of elements of a list. */
func (e *Element) Length() int {
	if e.HasNext() {
		return e.next.Length() + 1
	}
	return 1
}

/* P05. Reverse a list. */
func (e *Element) Reverse() {
	if e.next.next != nil {
		e.next.Reverse()
	}
	e.next.next = e
	e.next = nil
}

func (e *Element) isMirrorPosition(pos int) bool {
	if pos > e.Length() {
		return false
	}
	return e.ElementAt(e.Length() - 1 - pos).value == e.value.(string)
}

/* P06. Find out whether a list is a palindrome. */
func (e *Element) print() string {
	if e.next != nil {
		return e.value.(string) + "," + e.next.print()
	}
	return e.value.(string)
}

func (e *Element) reversePrint() string {
	if e.next != nil {
		return e.next.reversePrint() + "," + e.value.(string)
	}
	return e.value.(string)
}

func (e *Element) IsPalindrome() bool {
	return e.print() == e.reversePrint()
}

/* P07. Flatten a nested list structure. */
func flatten(e *Element) *Element {
	switch t := e.value.(type) {
	case string:
		if e.HasNext() {
			return &Element{value: t, next: flatten(e.next)}
		} else {
			return &Element{value: t}
		}
	default:
		var head, lastElement, flattenElement *Element
		for _, x := range t.([]*Element) {
			flattenElement = flatten(x)
			if flattenElement == nil {
				continue
			}
			if head == nil {
				lastElement = flattenElement
				head = lastElement
			} else {
				lastElement.next = flattenElement
				lastElement = flattenElement.Last()
			}
		}
		if head == nil {
			return nil
		}
		if e.HasNext() {
			lastElement.next = flatten(e.next)
		}

		return head
	}
	panic("No expected return")
}
