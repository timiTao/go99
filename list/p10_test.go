package list

import (
	"testing"
)

func TestRunLengthEncoding(t *testing.T) {
	expected := [][]string{[]string{"4", "a"}, []string{"1", "b"}, []string{"2", "c"}, []string{"2", "a"}, []string{"1", "d"}, []string{"4", "e"}}
	value := []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}
	if !Equal(runLengthEncoding(value), expected) {
		t.Errorf("Expected %v, but got %v", expected, runLengthEncoding(value))
	}
}
