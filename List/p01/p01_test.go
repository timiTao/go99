package p01

import "testing"

func TestLastElementEmpty(t *testing.T) {
	element := Element{value: "test"}
	last, correct := element.Last()

	if correct != true || last.value != "test" {
		t.Errorf("Expected %s, founded %s", "test", last.value)
	}
}

func TestLastElementList(t *testing.T) {
	element := Element{value: "test", next: &Element{value: "test2"}}
	last, correct := element.Last()

	if correct != true || last.value != "test2" {
		t.Errorf("Expected %s, founded %s", "test2", last.value)
	}
}
