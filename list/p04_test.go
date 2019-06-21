package list

/* P04 (*) Find the number of elements of a list. */

import "testing"

func TestLength(t *testing.T) {
	element := Element{value: "test", next: &Element{value: "test2", next: &Element{value: "test3"}}}
	value := element.Length()
	if value != 3 {
		t.Errorf("Expected %d elements. Current: %d", 3, value)
	}
}
