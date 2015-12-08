// Package calc is calucrator.
package calc

// Add is addition two ints.
func Add(a int, b int) (sum int) {
	sum = threeAdd(a, b, 0)
	return
}

func threeAdd(a int, b int, c int) (sum int) {
	sum = a + b + c
	return
}
