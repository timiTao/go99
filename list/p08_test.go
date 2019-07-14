package list

import "testing"

func TestEliminateDuplicationOne(t *testing.T) {
	e := &Element{value: "a"}

	result := eliminateDuplication(e)
	if result.print() != "a" {
		t.Errorf("Expected %s. Current: %s", "a", result.print())
	}
}

func TestEliminateDuplicationTwo(t *testing.T) {
	e := &Element{value: "a", next: &Element{value: "b"}}

	result := eliminateDuplication(e)
	if result.print() != "a,b" {
		t.Errorf("Expected %s. Current: %s", "a,b", result.print())
	}
}

func TestEliminateDuplicationTwoDuplication(t *testing.T) {
	e := &Element{value: "a", next: &Element{value: "a"}}

	result := eliminateDuplication(e)
	if result.print() != "a" {
		t.Errorf("Expected %s. Current: %s", "a", result.print())
	}
}

func TestEliminateDuplicationList(t *testing.T) {
	e := &Element{value: "a"}
	e.append("a").append("a").append("a").append("b").append("c").append("c").append("a").append("a").append("d").append("e").append("e")

	result := eliminateDuplication(e)
	if result.print() != "a,b,c,a,d,e" {
		t.Errorf("Expected %s. Current: %s", "a,b,c,a,d,e", result.print())
	}
}
