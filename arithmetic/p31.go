package arithmetic

/* P31 (**) Determine whether a given integer number is prime. */

func IsPrime(value int) bool {
	counter := 0

	for i := 1; i < value; i++ {
		if value%i == 0 {
			counter++
		}
	}

	return counter == 1
}
