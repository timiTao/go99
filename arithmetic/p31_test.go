package arithmetic

import (
	"testing"
)

func TestCorrectPrime(t *testing.T) {

	if !IsPrime(7) {
		t.Errorf("Expected true, but founded false", )
	}
}

func TestIncorrectPrime(t *testing.T) {

	if IsPrime(8) {
		t.Errorf("Expected false, but founded true", )
	}
}
