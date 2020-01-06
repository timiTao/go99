package arithmetic

/* P35. Determine the prime factors of a given positive integer. */

func FlatFactors(m int) []int {
	counter := 0
	list := []int{}

	value := m
	for i := 2; i < m; i++ {
		for value%i == 0 {
			value = value / i
			list = append(list, i)
			counter++
		}
	}

	return list
}
