package main

import "testing"

func TestlastElement(t *testing.T) {
	tempList := []int{2, 3, 5, 7, 11, 13}
	last := lastElement(tempList)

	if last != 13 {
		t.Errorf("Expected %d, founded %d", 13, last)
	}
}