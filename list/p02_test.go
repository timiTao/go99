package list

/* P02 (*) Find the last but one element of a list. */

import (
	"testing"
)

func TestLastButOneEmpty(t *testing.T) {
	element := Element{value: "test"}
	last, correct := element.LastButOne()

	if correct != false || last != nil {
		t.Errorf("Expected not to found element")
	}
}

func TestLastButOneElement(t *testing.T) {
	element := Element{value: "test", next: &Element{value: "test2", next: &Element{value: "test3"}}}
	lastButOne, correct := element.LastButOne()

	if correct != true || lastButOne.value != "test2" {
		t.Errorf("Expected %s, founded %s", "test2", lastButOne.value)
	}
}
