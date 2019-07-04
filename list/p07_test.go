package list

import "testing"

func TestRootFlatten(t *testing.T) {
	element := Element{
		value: "a",
		next: &Element{
			value: "b",
			next:  &Element{value: "c"},
		},
	}

	result := flatten(&element)
	if result.Length() != 3 {
		t.Errorf("Expected 3 elements. Got: %d", result.Length())
	}
	if result.print() != "a,b,c" {
		t.Errorf("Expected a,b,c elements. Got: %s", result.print())
	}
}

func TestListInValueRootFlatten(t *testing.T) {
	element := Element{value: []*Element{{value: "a"}, {value: "b"}}}

	result := flatten(&element)
	if result.Length() != 2 {
		t.Errorf("Expected 2 elements. Got: %d", result.Length())
	}
	if result.print() != "a,b" {
		t.Errorf("Expected a,b elements. Got: %s", result.print())
	}
}

func TestOneListFlatten(t *testing.T) {
	element := Element{
		value: "a",
		next: &Element{
			value: []*Element{{value: "b"}},
		},
	}

	result := flatten(&element)
	if result.Length() != 2 {
		t.Errorf("Expected 2 elements. Got: %d", result.Length())
	}

	if result.print() != "a,b" {
		t.Errorf("Expected a,b elements. Got: %s", result.print())
	}
}

func TestMixedListFlatten(t *testing.T) {
	element := Element{
		value: "a",
		next: &Element{
			value: []*Element{{value: "b"}},
			next:  &Element{value: "e"},
		},
	}

	result := flatten(&element)
	if result.Length() != 3 {
		t.Errorf("Expected 3 elements. Got: %d", result.Length())
	}

	if result.print() != "a,b,e" {
		t.Errorf("Expected a,b,e elements. Got: %s", result.print())
	}
}

func TestFlatten(t *testing.T) {
	element := Element{
		value: "a",
		next: &Element{
			value: []*Element{{value: "b"}, {value: "c", next: &Element{value: "d"}}},
			next:  &Element{value: "e"},
		},
	}

	result := flatten(&element)
	if result.Length() != 5 {
		t.Errorf("Expected 5 elements. Got: %d", result.Length())
	}

	if result.print() != "a,b,c,d,e" {
		t.Errorf("Expected a,b,c,d,e elements. Got: %s", result.print())
	}
}
