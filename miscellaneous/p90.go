package miscellaneous

import "math"

func IsCorrectEightQueensProblemSolution(solution [8]int) bool {
	for i := 0; i < 8; i++ {
		if checkQueenCollision(solution, i) {
			return false
		}
	}
	return true
}

func checkQueenCollision(solution [8]int, checkedKey int) bool {
	for i := 7; i > checkedKey; i-- {
		if checkCorrespondingQueens(solution, checkedKey, i) {
			return true
		}
	}
	return false
}

func checkCorrespondingQueens(solution [8]int, keyA int, keyB int) bool {
	modValue, _ := math.Modf(float64(solution[keyA]) - float64(solution[keyB]))
	modKey, _ := math.Modf(float64(keyA) - float64(keyB))
	return modValue == modKey
}
