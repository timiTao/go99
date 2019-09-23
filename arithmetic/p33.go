package arithmetic

func IsCoprime(a int, b int) bool {
	return GreatestCommonDivisor(a, b) == 1
}
