package arithmetic

/* P32 (**) Determine the greatest common divisor of two positive integer numbers. */

func GreatestCommonDivisor(a int, b int) int {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, a%b)
}
