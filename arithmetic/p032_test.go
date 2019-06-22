package arithmetic

import "testing"

func TestGreatestCommonDivisor(t *testing.T) {
	commonDivisor := GreatestCommonDivisor(1989, 867)
	if commonDivisor != 51 {
		t.Errorf("Expected %d, but got %d", 51, commonDivisor)
	}
}
