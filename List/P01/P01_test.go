package P01

import "testing"

func TestLastElement(t *testing.T) {
	tempList := []int{2, 3, 5, 7, 11, 13}
	last, correct := lastElement(tempList)

	if last != 13 && correct == true {
		t.Errorf("Expected %d, founded %d", 13, last)
	}
}

func TestLastElementOnEmpty(t *testing.T) {
	tempList := []int{}
	last, correct := lastElement(tempList)

	if last != 0 || correct != false {
		t.Errorf("Expected found incorrect status")
	}
}