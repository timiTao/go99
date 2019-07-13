package logic

import "testing"

func TestAnd(t *testing.T) {
	if !and(true, true) {
		t.Error("Expected to be true")
	}
}
