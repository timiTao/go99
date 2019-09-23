package arithmetic

import (
	"testing"
)

func TestIsCoprime(t *testing.T) {
	if !IsCoprime(35, 64) {
		t.Errorf("Expected true, but founded false")
	}
}

func TestIsNotCoprime(t *testing.T) {
	if IsCoprime(6, 64) {
		t.Errorf("Expected true, but founded false")
	}
}
