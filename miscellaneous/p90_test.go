package miscellaneous

import "testing"

func TestCorrectSolution(t *testing.T) {
	correct := IsCorrectEightQueensProblemSolution([8]int{4, 1, 5, 8, 2, 7, 3, 6})
	if correct == false {
		t.Errorf("Expected mark as correct combination")
	}
}

func TestIncorrectSolution(t *testing.T) {
	correct := IsCorrectEightQueensProblemSolution([8]int{4, 3, 5, 8, 2, 7, 1, 6})
	if correct == true {
		t.Errorf("Expected mark as incorrect combination")
	}
}
