package logic

/* P46. Truth tables for logical expressions */
func and(a bool, b bool) bool {
	return a && b
}
func or(a bool, b bool) bool {
	return a || b
}
func nand(a bool, b bool) bool {
	return !and(a, b)
}
func nor(a bool, b bool) bool {
	return !or(a, b)
}
func xor(a bool, b bool) bool {
	return (a || b) && !(a && b)
}
