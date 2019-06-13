package P03

import "testing"

func TestElementAt(t *testing.T) {
	tempList := []int{2, 3, 5, 7, 11, 13}

	last, correct := elementAt(tempList, 3)

	if last != 5 || correct != true {
		t.Errorf("Expected %d, founded %d", 5, last)
	}
}

func TestElementAtOnEmpty(t *testing.T) {
	tempList := []int{}

	last, correct := elementAt(tempList, 3)

	if last != 0 || correct != false {
		t.Errorf("Expected found incorrect status")
	}
}
