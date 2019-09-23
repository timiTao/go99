package arithmetic

/* P34 Calculate Euler's totient function phi(m) */

func EulerTotiem(m int) int {
	result := 0
	for i := 1; i < m; i++ {
		if IsCoprime(i, m) {
			result = result + 1
		}
	}
	return result
}
