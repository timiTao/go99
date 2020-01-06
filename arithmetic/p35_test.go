package arithmetic

import (
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestFlatFactors(t *testing.T) {
	expected := []int{3, 3, 5, 7}
	if !Equal(FlatFactors(315), expected) {
		t.Errorf("Expected %v, but got %d", expected, FlatFactors(315))
	}
}

func TestFlatFactors2(t *testing.T) {
	expected := []int{2, 2, 5, 5}
	if !Equal(FlatFactors(100), expected) {
		t.Errorf("Expected %v, but got %d", expected, FlatFactors(100))
	}
}
