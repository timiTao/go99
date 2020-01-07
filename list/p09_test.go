package list

import (
	"testing"
)

func Equal(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !Equal2(v, b[i]) {
			return false
		}
	}
	return true
}

func Equal2(a, b []string) bool {
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

func TestTransformDuplicatesIntoSublists(t *testing.T) {
	expected := [][]string{[]string{"a", "a", "a", "a"}, []string{"b"}, []string{"c", "c"}, []string{"a", "a"}, []string{"d"}, []string{"e", "e", "e", "e"}}
	value := []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}
	if !Equal(transformDuplicatesIntoSublists(value), expected) {
		t.Errorf("Expected %v, but got %v", expected, transformDuplicatesIntoSublists(value))
	}
}
