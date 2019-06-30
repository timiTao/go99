package list

import "testing"

func TestPrint(t *testing.T) {
	element := Element{value: "test1", next: &Element{value: "test2", next: &Element{value: "test3"}}}

	if element.print() != "test1,test2,test3" {
		t.Errorf("Expected element to be printed. Got: %s", element.print())
	}
}

func TestReversePrint(t *testing.T) {
	element := Element{value: "test1", next: &Element{value: "test2", next: &Element{value: "test3"}}}

	if element.reversePrint() != "test3,test2,test1" {
		t.Errorf("Expected element to be printed. Current: %s", element.reversePrint())
	}
}

func TestIsPalindrome(t *testing.T) {
	element := Element{value: "test", next: &Element{value: "test2", next: &Element{value: "test"}}}

	if !element.IsPalindrome() {
		t.Errorf("Expected element to be palindrome. Current: %s", element.print())
	}
}
