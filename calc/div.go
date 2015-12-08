// Package calc is calucrator.
package calc

import "errors"

// Divison is division two ints.
func Division(a int, b int) (div int, err error) {
	if b == 0 {
		return div, errors.New("Cannot use zero in second argument.")
	}

	div = a / b
	return
}
