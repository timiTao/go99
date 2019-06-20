package p03

import "testing"

func TestElementAtOnEmpty(t *testing.T) {
	element := Element{value: "test"}

	current := element.ElementAt(3)

	if current == nil {
		t.Errorf("Expected to found element")
		return
	}

	if current.value != "test" {
		t.Errorf("Expected %s, founded %s", "test", current.value)
	}
}

func TestElementAt(t *testing.T) {
	element := Element{value: "test", next: &Element{value: "test2", next: &Element{value: "test3"}}}

	current := element.ElementAt(2)

	if current == nil {
		t.Errorf("Expected to found element")
		return
	}

	if current.value != "test3" {
		t.Errorf("Expected %s, founded %s", "test3", current.value)
	}
}
