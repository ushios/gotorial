package calc

import "testing"

func TestDivision(t *testing.T) {
	var a, b int
	a = 10
	b = 5

	div, err := Division(a, b)

	if err != nil {
		t.Errorf("err want nil.")
	}

	if div != 2 {
		t.Errorf("sum want (%d).", 2)
	}
}
