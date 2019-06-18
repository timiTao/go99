package p04

import "testing"

func TestLength(t *testing.T) {
	tempList := []int{2, 3, 5, 7, 11, 13}
	value := length(tempList)
	if value != len(tempList) {
		t.Errorf("Expected %d elements. Current: %d", len(tempList), value)
	}
}
