package miscellaneous

import "math"

func IsCorrectEightQueensProblemSolution(solution [8]int) bool {
	for i := 0; i < 8; i++ {
		if queensCollision(solution, solution[i]) {
			return false
		}
	}
	return true
}

func queensCollision(solution [8]int, key int) bool {
	for i := 0; i < 8; i++ {
		if i == key {
			continue
		}
		valueI := solution[i]
		valueKey := solution[key]

		modValue, _ := math.Modf(float64(valueI) - float64(valueKey))
		modKey, _ := math.Modf(float64(i) - float64(key))
		if modValue == modKey {
			return false
		}
	}
	return true
}
