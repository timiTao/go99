package list

import "testing"

func Test(t *testing.T) {
	test := &Element{value: "test3"}
	element := Element{value: "test", next: &Element{value: "test2", next: test}}
	element.Reverse()

	lastElement := test.Last()

	if lastElement.value != "test" {
		t.Errorf("Expected %s elements. Current: %s", test.value, lastElement.value)
	}
}
