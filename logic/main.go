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
func not(a bool) bool {
	return !a
}
func impl(a bool, b bool) bool {
	return or(not(a), b)
}
func equ(a bool, b bool) bool {
	return a == b
}
func printLogicTable(f func(bool, bool) bool) [4][3]bool {
	return [4][3]bool {
	   {true, true, f(true, true)},
	   {true, false, f(true, false)},
	   {false, true, f(false, true)},
	   {false, false, f(false, false)},  
	}
}
