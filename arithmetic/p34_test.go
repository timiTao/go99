package arithmetic

import (
	"testing"
)

func TestEulerTotiem(t *testing.T) {
	if EulerTotiem(10) != 4 {
		t.Errorf("Expected %d, but got %d", 4, EulerTotiem(10))
	}
}
